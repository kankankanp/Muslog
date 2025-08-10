provider "aws" {
  region = var.aws_region
}

data "aws_caller_identity" "current" {}

module "network" {
  source      = "../../modules/network"
  aws_region  = var.aws_region
  environment = var.environment
}

module "s3" {
  source                          = "../../modules/s3"
  project_name                    = var.project_name
  environment                     = var.environment
  ecs_task_execution_role_arn     = module.ecs.ecs_task_execution_role_arn
}

module "alb" {
  source                = "../../modules/alb"
  project_name          = var.project_name
  environment           = var.environment
  vpc_id                = module.network.vpc_id
  public_subnet_ids     = module.network.public_subnet_ids
  alb_security_group_id = module.network.alb_sg_id
}

module "rds" {
  source               = "../../modules/rds"
  project_name         = var.project_name
  environment          = var.environment
  private_subnet_ids   = module.network.private_subnet_ids
  db_security_group_id = module.network.db_sg_id
  db_name              = var.db_name
  db_username          = var.db_username
  db_password          = var.db_password
}

module "ecs" {
  source                  = "../../modules/ecs"
  environment             = var.environment
  aws_region              = var.aws_region
  aws_account_id          = data.aws_caller_identity.current.account_id
  vpc_id                  = module.network.vpc_id
  public_subnet_ids       = module.network.public_subnet_ids
  private_subnet_ids      = module.network.private_subnet_ids
  alb_sg_id               = module.network.alb_sg_id
  db_host                 = module.rds.db_cluster_endpoint
  db_port                 = module.rds.db_cluster_port
  db_username             = var.db_username
  db_password             = var.db_password
  db_name                 = var.db_name
  db_security_group_id    = module.network.db_sg_id
  backend_target_group_arn  = module.alb.alb_target_group_arn
  app_secrets_secret_arn    = aws_secretsmanager_secret.app_secrets.arn
  depends_on = [module.alb]
}

module "cloudfront" {
  source                              = "../../modules/cloudfront"
  s3_bucket_regional_domain_name      = module.s3.frontend_bucket_regional_domain_name
  s3_origin_access_identity_path      = module.s3.s3_origin_access_identity_path
  alb_dns_name                        = module.alb.alb_dns_name
  environment                         = var.environment
}

resource "aws_secretsmanager_secret" "db_password" {
  name = "production/db_password"
}

resource "aws_secretsmanager_secret_version" "db_password_version" {
  secret_id     = aws_secretsmanager_secret.db_password.id
  secret_string = var.db_password
}
