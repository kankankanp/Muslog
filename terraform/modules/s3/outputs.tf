output "frontend_bucket_name" {
  description = "The name of the frontend S3 bucket."
  value       = aws_s3_bucket.frontend_bucket.id
}

output "frontend_bucket_regional_domain_name" {
  description = "The regional domain name of the frontend S3 bucket."
  value       = aws_s3_bucket.frontend_bucket.bucket_regional_domain_name
}

output "s3_origin_access_identity_path" {
  description = "The path for the S3 Origin Access Identity."
  value       = aws_cloudfront_origin_access_identity.s3_oai.cloudfront_access_identity_path
}

output "media_bucket_name" {
  description = "The name of the media S3 bucket."
  value       = aws_s3_bucket.media_bucket.id
}
