output "posts_table_name" {
  description = "The name of the DynamoDB posts table"
  value       = aws_dynamodb_table.posts.name
}

output "users_table_name" {
  description = "The name of the DynamoDB users table"
  value       = aws_dynamodb_table.users.name
}
