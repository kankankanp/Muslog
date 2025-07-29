variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "ap-northeast-1"
}

variable "environment" {
  description = "Deployment environment (e.g., develop, production)"
  type        = string
  default     = "develop"
}

variable "ec2_ami" {
  description = "AMI for the EC2 instance"
  type        = string
  default     = "ami-0c55b159cbfafe1f0" # Example: Amazon Linux 2 in ap-northeast-1
}

variable "ec2_instance_type" {
  description = "Instance type for the EC2 instance"
  type        = string
  default     = "t2.micro"
}

variable "ec2_key_pair_name" {
  description = "EC2 Key Pair name"
  type        = string
  default     = "your-key-pair-name" # IMPORTANT: Change this to your actual key pair name
}

variable "db_instance_class" {
  description = "Instance class for the RDS instance"
  type        = string
  default     = "db.t2.micro"
}

variable "db_username" {
  description = "Username for the RDS database"
  type        = string
  default     = "admin"
}

variable "db_password" {
  description = "Password for the RDS database"
  type        = string
  default     = "password" # IMPORTANT: Change this to a strong password and consider using AWS Secrets Manager
}