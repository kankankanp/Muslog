variable "environment" {
  description = "The deployment environment."
  type        = string
}

variable "s3_bucket_regional_domain_name" {
  description = "The regional domain name of the S3 bucket for the frontend."
  type        = string
}

variable "s3_origin_access_identity_path" {
  description = "The path for the S3 Origin Access Identity."
  type        = string
}

variable "alb_dns_name" {
  description = "The DNS name of the ALB."
  type        = string
}

variable "enable_custom_domain" {
  description = "Flag to enable custom domain aliases and ACM certificate."
  type        = bool
  default     = false
}

variable "domain_name" {
  description = "The custom domain name."
  type        = string
  default     = ""
}

variable "acm_certificate_arn" {
  description = "The ARN of the ACM certificate."
  type        = string
  default     = null
}