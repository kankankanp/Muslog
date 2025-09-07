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

variable "lambda_edge_origin_request_arn" {
  description = "(Optional) Qualified ARN of Lambda@Edge function for origin-request."
  type        = string
  default     = ""
}

variable "url_rewrite_function_path" {
  description = "The path to the URL rewrite function code."
  type        = string
}

