# IAM Role for EventBridge to trigger ECS tasks
resource "aws_iam_role" "events_to_ecs_role" {
  name = "${var.environment}-scheduler-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action    = "sts:AssumeRole",
        Effect    = "Allow",
        Principal = {
          Service = "events.amazonaws.com"
        }
      },
      {
        Action    = "sts:AssumeRole",
        Effect    = "Allow",
        Principal = {
          Service = "ssm.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_policy" "events_to_ecs_policy" {
  name        = "${var.environment}-events-to-ecs-policy"
  description = "Allow EventBridge to run ECS tasks"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect   = "Allow",
        Action   = "iam:PassRole",
        Resource = [
          var.ecs_task_execution_role_arn,
          var.ecs_task_role_arn
        ]
      },
      {
        Effect   = "Allow",
        Action   = "ecs:RunTask",
        Resource = var.scheduler_task_definition_arn,
        Condition = {
          StringEquals = {
            "ecs:cluster" = var.ecs_cluster_arn
          }
        }
      },
      {
        Effect   = "Allow",
        Action   = "iam:PassRole",
        Resource = var.scheduler_execution_role_arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "events_to_ecs_policy_attachment" {
  role       = aws_iam_role.events_to_ecs_role.name
  policy_arn = aws_iam_policy.events_to_ecs_policy.arn
}

# EventBridge Rule to stop resources
resource "aws_cloudwatch_event_rule" "stop_rule" {
  name                = "${var.environment}-stop-resources-rule"
  description         = "Stop ECS and RDS resources on a schedule (JST 20:00)"
  schedule_expression = "cron(0 11 ? * MON-FRI *)" # UTC 11:00 = JST 20:00
  is_enabled          = false
}

# EventBridge Rule to start resources
resource "aws_cloudwatch_event_rule" "start_rule" {
  name                = "${var.environment}-start-resources-rule"
  description         = "Start ECS and RDS resources on a schedule (JST 08:00)"
  schedule_expression = "cron(0 23 ? * SUN-THU *)" # UTC 23:00 = JST 08:00 next day
  is_enabled          = false
}

# Target for the stop rule
resource "aws_cloudwatch_event_target" "stop_target" {
  rule      = aws_cloudwatch_event_rule.stop_rule.name
  target_id = "${var.environment}-stop-ecs-task"
  arn       = var.ecs_cluster_arn
  role_arn  = aws_iam_role.events_to_ecs_role.arn

  ecs_target {
    task_count          = 1
    task_definition_arn = var.scheduler_task_definition_arn
    launch_type         = "FARGATE"

    network_configuration {
      subnets          = var.private_subnet_ids
      security_groups  = [var.ecs_tasks_sg_id]
      assign_public_ip = false
    }
  }

  input_transformer {
    input_paths = {
      "action" = "$.detail.action"
    }
    input_template = jsonencode({
      "containerOverrides" = [
        {
          "name" = "scheduler",
          "environment" = [
            {
              "name"  = "ACTION",
              "value" = "stop"
            }
          ]
        }
      ]
    })
  }
}

# Target for the start rule
resource "aws_cloudwatch_event_target" "start_target" {
  rule      = aws_cloudwatch_event_rule.start_rule.name
  target_id = "${var.environment}-start-ecs-task"
  arn       = var.ecs_cluster_arn
  role_arn  = aws_iam_role.events_to_ecs_role.arn

  ecs_target {
    task_count          = 1
    task_definition_arn = var.scheduler_task_definition_arn
    launch_type         = "FARGATE"

    network_configuration {
      subnets          = var.private_subnet_ids
      security_groups  = [var.ecs_tasks_sg_id]
      assign_public_ip = false
    }
  }

  input_transformer {
    input_paths = {
      "action" = "$.detail.action"
    }
    input_template = jsonencode({
      "containerOverrides" = [
        {
          "name" = "scheduler",
          "environment" = [
            {
              "name"  = "ACTION",
              "value" = "start"
            }
          ]
        }
      ]
    })
  }
}

# SSM Document for manual execution
resource "aws_ssm_document" "run_scheduler_task" {
  name          = "${var.environment}-run-scheduler-task"
  document_type = "Automation"
  content = jsonencode({
    schemaVersion = "0.3",
    description   = "Run ECS Scheduler Task to start or stop resources.",
    assumeRole    = var.scheduler_execution_role_arn,
    parameters = {
      Action = {
        type        = "String",
        description = "(Required) Specify 'start' or 'stop'.",
        allowedValues = [
          "start",
          "stop"
        ]
      }
    },
    mainSteps = [
      {
        name   = "runSchedulerTask",
        action = "aws:executeAwsApi",
        inputs = {
          Service        = "ecs",
          Api            = "RunTask",
          cluster        = var.ecs_cluster_arn,
          taskDefinition = var.scheduler_task_definition_arn,
          launchType     = "FARGATE",
          networkConfiguration = {
            awsvpcConfiguration = {
              subnets        = var.private_subnet_ids,
              securityGroups = [var.ecs_tasks_sg_id],
              assignPublicIp = "DISABLED"
            }
          },
          overrides = {
            containerOverrides = [
              {
                name        = "scheduler",
                environment = [
                  {
                    name  = "ACTION",
                    value = "{{ Action }}"
                  }
                ]
              }
            ]
          }
        }
      }
    ]
  })

  tags = {
    Environment = var.environment
  }
}
