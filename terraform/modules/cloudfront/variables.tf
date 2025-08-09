variable "s3_bucket_regional_domain_name" {
  description = "The regional domain name of the S3 bucket."
  type        = string
}

variable "s3_origin_access_identity_arn" {
  description = "The ARN of the S3 Origin Access Identity."
  type        = string
}

variable "alb_dns_name" {
  description = "The DNS name of the ALB."
  type        = string
}

variable "domain_name" {
  description = "The main domain name."
  type        = string
}

variable "environment" {
  description = "The deployment environment (e.g., develop, production)."
  type        = string
}
