output "lambda_function_qualified_arn" {
  description = "Qualified ARN (versioned) of the Lambda@Edge function"
  value       = aws_lambda_function.edge_ssr.qualified_arn
}

output "lambda_role_name" {
  description = "Name of the Lambda@Edge IAM role"
  value       = aws_iam_role.lambda_edge_role.name
}

output "lambda_role_arn" {
  description = "ARN of the Lambda@Edge IAM role"
  value       = aws_iam_role.lambda_edge_role.arn
}
