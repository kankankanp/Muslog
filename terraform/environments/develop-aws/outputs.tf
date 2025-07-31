output "vpc_id" {
  description = "The ID of the VPC"
  value       = module.network.vpc_id
}

output "public_subnet_ids" {
  description = "The IDs of the public subnets"
  value       = module.network.public_subnet_ids
}

output "private_subnet_ids" {
  description = "The IDs of the private subnets"
  value       = module.network.private_subnet_ids
}

output "ec2_public_ip" {
  description = "The public IP address of the EC2 instance"
  value       = module.backend.ec2_public_ip
}

output "rds_endpoint" {
  description = "The endpoint of the RDS instance"
  value       = module.database.rds_endpoint
}

output "frontend_s3_bucket_name" {
  description = "The name of the S3 bucket for the frontend"
  value       = module.frontend.frontend_s3_bucket_name
}

output "cloudfront_domain_name" {
  description = "The domain name of the CloudFront distribution"
  value       = module.frontend.cloudfront_domain_name
}

output "uploads_s3_bucket_name" {
  description = "The name of the S3 bucket for uploads"
  value       = module.backend.uploads_s3_bucket_name
}
