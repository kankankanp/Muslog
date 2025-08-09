output "db_cluster_endpoint" {
  description = "The endpoint of the RDS cluster."
  value       = aws_rds_cluster.main.endpoint
}

output "db_cluster_port" {
  description = "The port of the RDS cluster."
  value       = aws_rds_cluster.main.port
}
