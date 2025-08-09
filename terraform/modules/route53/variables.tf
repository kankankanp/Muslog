variable "domain_name" {
  description = "The domain name for Route 53 zone."
  type        = string
}

variable "cloudfront_distribution_domain_name" {
  description = "The domain name of the CloudFront distribution."
  type        = string
}

variable "cloudfront_distribution_hosted_zone_id" {
  description = "The CloudFront distribution's hosted zone ID."
  type        = string
}