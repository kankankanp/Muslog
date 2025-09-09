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

variable "function_name" {
  description = "Override Lambda function name. Leave empty to use default."
  type        = string
  default     = ""
}

variable "role_name_suffix" {
  description = "Optional suffix to make IAM role name unique per function (e.g., 'image')."
  type        = string
  default     = ""
}

variable "skip_destroy" {
  description = "When true, do not delete the Lambda function on destroy (useful for Lambda@Edge replicas)."
  type        = bool
  default     = false
}

variable "zip_name_suffix" {
  description = "Optional suffix to make archive filename unique per instance (e.g., '-image')."
  type        = string
  default     = ""
}
