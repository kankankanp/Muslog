output "ecs_task_execution_role_arn" {
  description = "The ARN of the ECS task execution role"
  value       = aws_iam_role.ecs_task_execution_role.arn
}

output "ecs_cluster_arn" {
  description = "The ARN of the ECS cluster"
  value       = aws_ecs_cluster.main.arn
}

output "scheduler_task_definition_arn" {
  description = "The ARN of the scheduler ECS task definition"
  value       = aws_ecs_task_definition.scheduler.arn
}

output "ecs_task_role_arn" {
  description = "The ARN of the ECS task role"
  value       = aws_iam_role.ecs_task_role.arn
}

output "ecs_tasks_sg_id" {
  description = "The ID of the security group for the ECS tasks"
  value       = aws_security_group.ecs_tasks.id
}
