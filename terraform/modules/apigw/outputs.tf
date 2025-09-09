output "api_endpoint" {
  description = "The HTTPS endpoint of the HTTP API"
  value       = aws_apigatewayv2_api.http_api.api_endpoint
}

output "api_host" {
  description = "The hostname of the HTTP API endpoint"
  value       = regex("https?://([^/]+)/?", aws_apigatewayv2_api.http_api.api_endpoint)[0]
}

