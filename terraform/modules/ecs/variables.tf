variable "environment" {
  description = "Deployment environment"
  type        = string
}

variable "aws_region" {
  description = "AWS region"
  type        = string
}

variable "aws_account_id" {
  description = "AWS account ID"
  type        = string
}

variable "vpc_id" {
  description = "The ID of the VPC"
  type        = string
}

variable "public_subnet_ids" {
  description = "List of public subnet IDs"
  type        = list(string)
}

variable "private_subnet_ids" {
  description = "List of private subnet IDs"
  type        = list(string)
}

variable "alb_sg_id" {
  description = "The ID of the ALB security group"
  type        = string
}

variable "posts_table_name" {
  description = "The name of the DynamoDB posts table"
  type        = string
}

variable "users_table_name" {
  description = "The name of the DynamoDB users table"
  type        = string
}

variable "aws_account_id" {
  description = "AWS Account ID"
  type        = string
}
