provider "aws" {
  region = var.aws_region
}

data "aws_caller_identity" "current" {}

resource "aws_acm_certificate" "main" {
  count = var.enable_custom_domain ? 1 : 0

  domain_name       = var.domain_name
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "cert_validation" {
  for_each = var.enable_custom_domain ? {
    for dvo in aws_acm_certificate.main[0].domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  } : {}

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = module.route53[0].route53_zone_id
}

resource "aws_acm_certificate_validation" "main" {
  count = var.enable_custom_domain ? 1 : 0

  certificate_arn         = aws_acm_certificate.main[0].arn
  validation_record_fqdns = [for record in aws_route53_record.cert_validation : record.fqdn]
}

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
  
  enable_https          = var.enable_custom_domain
  acm_certificate_arn   = var.enable_custom_domain ? aws_acm_certificate.main[0].arn : null
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
  depends_on = [module.alb]
}

module "cloudfront" {
  source                              = "../../modules/cloudfront"
  s3_bucket_regional_domain_name      = module.s3.frontend_bucket_regional_domain_name
  s3_origin_access_identity_path      = module.s3.s3_origin_access_identity_path
  alb_dns_name                        = module.alb.alb_dns_name
  
  enable_custom_domain                = var.enable_custom_domain
  domain_name                         = var.domain_name
  acm_certificate_arn                 = var.enable_custom_domain ? aws_acm_certificate.main[0].arn : null
  environment                         = var.environment
}

module "route53" {
  count = var.enable_custom_domain ? 1 : 0

  source                               = "../../modules/route53"
  domain_name                          = var.domain_name
  cloudfront_distribution_domain_name  = module.cloudfront.cloudfront_distribution_domain_name
  cloudfront_distribution_hosted_zone_id = module.cloudfront.cloudfront_distribution_hosted_zone_id
}