variable "environment" {
  description = "Deployment environment (e.g., develop, production)"
  type        = string
}

variable "vpc_id" {
  description = "The ID of the VPC"
  type        = string
}

variable "public_subnet_ids" {
  description = "The IDs of the public subnets"
  type        = list(string)
}

variable "ec2_ami" {
  description = "AMI for the EC2 instance"
  type        = string
}

variable "ec2_instance_type" {
  description = "Instance type for the EC2 instance"
  type        = string
}

variable "ec2_key_pair_name" {
  description = "EC2 Key Pair name"
  type        = string
}
