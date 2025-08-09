variable "project_name" {
  description = "The name of the project."
  type        = string
}

variable "environment" {
  description = "The deployment environment (e.g., develop, production)."
  type        = string
}

variable "private_subnet_ids" {
  description = "A list of private subnet IDs for the DB subnet group."
  type        = list(string)
}

variable "db_security_group_id" {
  description = "The ID of the security group for the RDS cluster."
  type        = string
}

variable "db_name" {
  description = "The name of the database."
  type        = string
}

variable "db_username" {
  description = "The master username for the database."
  type        = string
}

variable "db_password" {
  description = "The master password for the database."
  type        = string
  sensitive   = true
}

variable "db_instance_class" {
  description = "The instance class for the RDS instances."
  type        = string
  default     = "db.t3.medium"
}

variable "db_instance_count" {
  description = "The number of RDS instances to create."
  type        = number
  default     = 1
}
