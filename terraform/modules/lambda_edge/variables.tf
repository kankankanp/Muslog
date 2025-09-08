variable "function_source_dir" {
  description = "Local directory containing Lambda@Edge function source."
  type        = string
}

variable "environment" {
  description = "Deployment environment name."
  type        = string
}

variable "environment_variables" {
  description = "Environment variables to inject into the Lambda@Edge function."
  type        = map(string)
  default     = {}
}

variable "cache_bucket_arn" {
  description = "Optional ARN of S3 bucket used for ISR/cache by the function. Grants s3 read/write."
  type        = string
  default     = ""
}
