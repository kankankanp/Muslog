provider "aws" {
  region = var.aws_region
}

module "network" {
  source      = "../../modules/network"
  aws_region  = var.aws_region
  environment = var.environment
}

module "backend" {
  source            = "../../modules/backend"
  environment       = var.environment
  vpc_id            = module.network.vpc_id
  public_subnet_ids = module.network.public_subnet_ids
  ec2_ami           = var.ec2_ami
  ec2_instance_type = var.ec2_instance_type
  ec2_key_pair_name = var.ec2_key_pair_name
}

module "database" {
  source             = "../../modules/database"
  environment        = var.environment
  vpc_id             = module.network.vpc_id
  private_subnet_ids = module.network.private_subnet_ids
  ec2_sg_id          = module.backend.ec2_sg_id
  db_instance_class  = var.db_instance_class
  db_username        = var.db_username
  db_password        = var.db_password
}

module "frontend" {
  source      = "../../modules/frontend"
  environment = var.environment
}