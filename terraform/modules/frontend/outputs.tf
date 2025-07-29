output "frontend_s3_bucket_name" {
  description = "The name of the S3 bucket for the frontend"
  value       = aws_s3_bucket.frontend_bucket.bucket
}

output "cloudfront_domain_name" {
  description = "The domain name of the CloudFront distribution"
  value       = aws_cloudfront_distribution.frontend_distribution.domain_name
}
