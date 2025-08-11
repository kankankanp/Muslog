output "scheduler_execution_role_arn" {
  description = "The ARN of the scheduler execution role"
  value       = aws_iam_role.events_to_ecs_role.arn
}
