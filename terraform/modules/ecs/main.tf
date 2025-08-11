# ECR
resource "aws_ecr_repository" "backend" {
  name = "${var.environment}/backend"
}

# CloudWatch Log Group
resource "aws_cloudwatch_log_group" "backend" {
  name              = "/ecs/${var.environment}/backend"
  retention_in_days = 7

  tags = {
    Environment = var.environment
  }
}

# IAM Roles
resource "aws_iam_role" "ecs_task_execution_role" {
  name = "${var.environment}-ecs-task-execution-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_policy" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role" "ecs_task_role" {
  name = "${var.environment}-ecs-task-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_policy" "ecs_secrets_policy" {
  name        = "${var.environment}-ecs-secrets-policy"
  description = "Allow ECS tasks to read secrets from Secrets Manager"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "secretsmanager:GetSecretValue",
          "kms:Decrypt"
        ],
        Resource = "*" # In production, you should restrict this to specific secret ARNs
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_secrets_policy_attachment" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = aws_iam_policy.ecs_secrets_policy.arn
}

# ECS
resource "aws_ecs_cluster" "main" {
  name = "${var.environment}-cluster"
}

resource "aws_security_group" "ecs_tasks" {
  name        = "${var.environment}-ecs-tasks-sg"
  description = "Allow inbound traffic from ALB"
  vpc_id      = var.vpc_id

  ingress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    security_groups = [var.alb_sg_id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Backend Task Definition and Service
resource "aws_ecs_task_definition" "backend" {
  family                   = "${var.environment}-backend"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 512
  memory                   = 1024
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_role.arn

  container_definitions = jsonencode([
    {
      name      = "backend",
      image     = "${aws_ecr_repository.backend.repository_url}:latest",
      portMappings = [
        {
          containerPort = 8080,
          hostPort      = 8080
        }
      ],
      environment = [
        { name = "DB_HOST", value = var.db_host },
        { name = "DB_PORT", value = var.db_port },
        { name = "DB_USER", value = var.db_username },
        { name = "DB_NAME", value = var.db_name },
        { name = "SPOTIFY_CLIENT_ID", value = var.spotify_client_id },
        { name = "GOOGLE_REDIRECT_URL", value = var.google_redirect_url },
        { name = "FRONTEND_URL", value = var.frontend_url },
        { name = "GOOGLE_CLIENT_ID", value = var.google_client_id }
      ],
      secrets = [
        { name = "DB_PASSWORD", valueFrom = "${var.app_secrets_secret_arn}:DB_PASSWORD::" },
        { name = "SPOTIFY_CLIENT_SECRET", valueFrom = "${var.app_secrets_secret_arn}:SPOTIFY_CLIENT_SECRET::" },
        { name = "GOOGLE_CLIENT_SECRET", valueFrom = "${var.app_secrets_secret_arn}:GOOGLE_CLIENT_SECRET::" }
      ],
      logConfiguration = {
        logDriver = "awslogs",
        options = {
          "awslogs-group"         = aws_cloudwatch_log_group.backend.name,
          "awslogs-region"        = var.aws_region,
          "awslogs-stream-prefix" = "ecs"
        }
      }
    }
  ])
}

resource "aws_ecs_service" "backend" {
  name            = "${var.environment}-backend-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.backend.arn
  desired_count   = 2
  launch_type     = "FARGATE"

  network_configuration {
    subnets         = var.private_subnet_ids
    security_groups = [aws_security_group.ecs_tasks.id, var.db_security_group_id]
  }

  load_balancer {
    target_group_arn = var.backend_target_group_arn
    container_name   = "backend"
    container_port   = 8080
  }
}

# ECR for Scheduler
resource "aws_ecr_repository" "scheduler" {
  name = "${var.environment}/scheduler"
}

# CloudWatch Log Group for Scheduler
resource "aws_cloudwatch_log_group" "scheduler" {
  name              = "/ecs/${var.environment}/scheduler"
  retention_in_days = 7

  tags = {
    Environment = var.environment
  }
}

# IAM Policy for Scheduler Task
resource "aws_iam_policy" "ecs_scheduler_task_policy" {
  name        = "${var.environment}-ecs-scheduler-task-policy"
  description = "Allow ECS scheduler task to start/stop RDS and update ECS service"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "ecs:UpdateService"
        ],
        Resource = aws_ecs_service.backend.id
      },
      {
        Effect = "Allow",
        Action = [
          "rds:StopDBCluster",
          "rds:StartDBCluster"
        ],
        Resource = var.db_cluster_arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_scheduler_task_policy_attachment" {
  role       = aws_iam_role.ecs_task_role.name
  policy_arn = aws_iam_policy.ecs_scheduler_task_policy.arn
}

# Scheduler Task Definition
resource "aws_ecs_task_definition" "scheduler" {
  family                   = "${var.environment}-scheduler"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 256
  memory                   = 512
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_role.arn

  container_definitions = jsonencode([
    {
      name      = "scheduler",
      image     = "${aws_ecr_repository.scheduler.repository_url}:latest",
      essential = true,
      environment = [
        { name = "ECS_CLUSTER_NAME", value = aws_ecs_cluster.main.name },
        { name = "ECS_SERVICE_NAME", value = aws_ecs_service.backend.name },
        { name = "DB_CLUSTER_IDENTIFIER", value = var.db_cluster_identifier },
        { name = "ECS_DESIRED_COUNT", value = "2" }
      ],
      logConfiguration = {
        logDriver = "awslogs",
        options = {
          "awslogs-group"         = aws_cloudwatch_log_group.scheduler.name,
          "awslogs-region"        = var.aws_region,
          "awslogs-stream-prefix" = "ecs-scheduler"
        }
      }
    }
  ])
}