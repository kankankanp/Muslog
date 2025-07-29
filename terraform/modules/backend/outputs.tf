output "ec2_public_ip" {
  description = "The public IP address of the EC2 instance"
  value       = aws_instance.backend_app.public_ip
}

output "ec2_sg_id" {
  description = "The ID of the EC2 security group"
  value       = aws_security_group.ec2_sg.id
}

output "uploads_s3_bucket_name" {
  description = "The name of the S3 bucket for uploads"
  value       = aws_s3_bucket.uploads_bucket.bucket
}