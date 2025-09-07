output "lambda_function_qualified_arn" {
  description = "Qualified ARN (versioned) of the Lambda@Edge function"
  value       = aws_lambda_function.edge_ssr.qualified_arn
}

