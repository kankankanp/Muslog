variable "environment" {
  description = "Deployment environment"
  type        = string
}

variable "aws_region" {
  description = "AWS region"
  type        = string
}

variable "aws_account_id" {
  description = "AWS account ID"
  type        = string
}

variable "vpc_id" {
  description = "The ID of the VPC"
  type        = string
}

variable "public_subnet_ids" {
  description = "List of public subnet IDs"
  type        = list(string)
}

variable "private_subnet_ids" {
  description = "List of private subnet IDs"
  type        = list(string)
}

variable "alb_sg_id" {
  description = "The ID of the ALB security group"
  type        = string
}

variable "db_host" {
  description = "The hostname of the RDS database."
  type        = string
}

variable "db_port" {
  description = "The port of the RDS database."
  type        = string
}

variable "db_username" {
  description = "The username for the RDS database."
  type        = string
}



variable "db_name" {
  description = "The name of the RDS database."
  type        = string
}

variable "db_security_group_id" {
  description = "The ID of the security group for the RDS database."
  type        = string
}

variable "backend_target_group_arn" {
  description = "The ARN of the backend ALB target group."
  type        = string
}

variable "app_secrets_secret_arn" {
  description = "The ARN of the secret containing application secrets."
  type        = string
  default     = ""
}

variable "google_redirect_url" {
  description = "The redirect URL for Google OAuth"
  type        = string
}

variable "frontend_url" {
  description = "The base URL of the frontend"
  type        = string
}

variable "db_cluster_arn" {
  description = "The ARN of the RDS cluster"
  type        = string
}

variable "db_cluster_identifier" {
  description = "The identifier of the RDS cluster"
  type        = string
}

variable "spotify_client_id" {
  description = "Spotify client ID"
  type        = string
  default     = ""
}

variable "google_client_id" {
  description = "Google client ID"
  type        = string
  default     = ""
}



