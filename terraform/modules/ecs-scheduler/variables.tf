variable "environment" {
  description = "Deployment environment"
  type        = string
}

variable "ecs_cluster_arn" {
  description = "The ARN of the ECS cluster"
  type        = string
}

variable "scheduler_task_definition_arn" {
  description = "The ARN of the scheduler ECS task definition"
  type        = string
}

variable "ecs_task_execution_role_arn" {
  description = "The ARN of the ECS task execution role"
  type        = string
}

variable "ecs_task_role_arn" {
  description = "The ARN of the ECS task role"
  type        = string
}

variable "private_subnet_ids" {
  description = "List of private subnet IDs for the ECS task"
  type        = list(string)
}

variable "ecs_tasks_sg_id" {
  description = "The ID of the security group for the ECS tasks"
  type        = string
}

variable "scheduler_execution_role_arn" {
  description = "The ARN of the role for the scheduler to assume"
  type        = string
}
