variable "project_name" {
  description = "The name of the project."
  type        = string
}

variable "environment" {
  description = "The deployment environment."
  type        = string
}

variable "vpc_id" {
  description = "The ID of the VPC."
  type        = string
}

variable "public_subnet_ids" {
  description = "List of public subnet IDs."
  type        = list(string)
}

variable "alb_security_group_id" {
  description = "The ID of the ALB security group."
  type        = string
}

variable "enable_https" {
  description = "Flag to enable HTTPS listener and related resources."
  type        = bool
  default     = false
}

variable "acm_certificate_arn" {
  description = "The ARN of the ACM certificate for the HTTPS listener."
  type        = string
  default     = null
}