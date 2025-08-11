output "alb_dns_name" {
  description = "The DNS name of the ALB."
  value       = aws_lb.main.dns_name
}

output "alb_zone_id" {
  description = "The hosted zone ID of the ALB."
  value       = aws_lb.main.zone_id
}

output "backend_target_group_arn" {
  description = "The ARN of the backend ALB target group."
  value       = aws_lb_target_group.main.arn
}
