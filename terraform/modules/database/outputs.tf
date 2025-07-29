output "rds_endpoint" {
  description = "The endpoint of the RDS instance"
  value       = aws_db_instance.main.address
}

output "rds_sg_id" {
  description = "The ID of the RDS security group"
  value       = aws_security_group.rds_sg.id
}
