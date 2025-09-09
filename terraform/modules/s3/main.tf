resource "aws_s3_bucket" "frontend_bucket" {
  bucket = "${var.project_name}-frontend-${var.environment}"

  tags = {
    Environment = var.environment
  }
}

resource "aws_s3_bucket_website_configuration" "frontend_website" {
  bucket = aws_s3_bucket.frontend_bucket.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "index.html"
  }
}

resource "aws_s3_bucket_policy" "frontend_bucket_policy" {
  bucket = aws_s3_bucket.frontend_bucket.id
  policy = data.aws_iam_policy_document.frontend_bucket_policy_document.json
}

data "aws_iam_policy_document" "frontend_bucket_policy_document" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.frontend_bucket.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.s3_oai.iam_arn]
    }
  }
}

resource "aws_cloudfront_origin_access_identity" "s3_oai" {
  comment = "OAI for S3 bucket access from CloudFront"
}

resource "aws_s3_bucket" "media_bucket" {
  bucket = "${var.project_name}-media-${var.environment}"

  tags = {
    Environment = var.environment
  }
}

resource "aws_s3_bucket_policy" "media_bucket_policy" {
  bucket = aws_s3_bucket.media_bucket.id
  policy = data.aws_iam_policy_document.media_bucket_policy_document.json
}

data "aws_iam_policy_document" "media_bucket_policy_document" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.media_bucket.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.s3_oai.iam_arn]
    }
  }

  statement {
    actions   = ["s3:PutObject"]
    resources = ["${aws_s3_bucket.media_bucket.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [var.ecs_task_execution_role_arn]
    }
  }

  dynamic "statement" {
    for_each = var.ecs_task_role_arn == "" ? [] : [var.ecs_task_role_arn]
    content {
      actions   = ["s3:PutObject"]
      resources = ["${aws_s3_bucket.media_bucket.arn}/*"]
      principals {
        type        = "AWS"
        identifiers = [statement.value]
      }
    }
  }
}

# OpenNext/ISR 用のキャッシュバケット（プライベート）
resource "aws_s3_bucket" "open_next_cache_bucket" {
  bucket = "${var.project_name}-open-next-cache-${var.environment}"

  tags = {
    Environment = var.environment
  }
}

resource "aws_s3_bucket_public_access_block" "open_next_cache_block" {
  bucket                  = aws_s3_bucket.open_next_cache_bucket.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}
