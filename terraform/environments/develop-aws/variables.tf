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
