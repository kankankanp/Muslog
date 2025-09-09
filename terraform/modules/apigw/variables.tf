variable "environment" {
  description = "Deployment environment name"
  type        = string
}

variable "lambda_function_arn" {
  description = "ARN of the regional Lambda to integrate with HTTP API"
  type        = string
}

variable "lambda_function_name" {
  description = "Name of the Lambda function (for permission)"
  type        = string
}

