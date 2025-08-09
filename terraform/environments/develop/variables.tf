variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "ap-northeast-1"
}

variable "environment" {
  description = "Deployment environment (e.g., develop, production)"
  type        = string
  default     = "develop"
}

variable "project_name" {
  description = "The name of the project."
  type        = string
  default     = "simple-blog"
}



variable "db_name" {
  description = "The name of the database."
  type        = string
  default     = "simpleblogdb"
}

variable "db_username" {
  description = "The master username for the database."
  type        = string
  default     = "admin"
}

variable "db_password" {
  description = "The master password for the database."
  type        = string
  sensitive   = true
}