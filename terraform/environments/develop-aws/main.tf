provider "aws" {
  region = var.aws_region
}

module "network" {
  source      = "../../modules/network"
  aws_region  = var.aws_region
  environment = var.environment
}

module "dynamodb" {
  source      = "../../modules/dynamodb"
  environment = var.environment
}

module "ecs" {
  source              = "../../modules/ecs"
  environment         = var.environment
  vpc_id              = module.network.vpc_id
  public_subnet_ids   = module.network.public_subnet_ids
  private_subnet_ids  = module.network.private_subnet_ids
  alb_sg_id           = module.network.alb_sg_id
  posts_table_name    = module.dynamodb.posts_table_name
  users_table_name    = module.dynamodb.users_table_name
}