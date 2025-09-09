resource "aws_cloudfront_function" "url_rewrite" {
  name    = "${var.environment}-url-rewrite"
  runtime = "cloudfront-js-1.0"
  comment = "Rewrites URLs for the single page application"
  publish = true
  code    = file(var.url_rewrite_function_path)
}

resource "aws_cloudfront_distribution" "main" {
  enabled             = true
  is_ipv6_enabled     = true
  comment             = "CloudFront distribution for ${var.environment}"
  # SSR有効化のためdefault_root_objectは未設定（"/"をLambda@Edgeで処理）

  origin {
    domain_name = var.s3_bucket_regional_domain_name
    origin_id   = "s3-origin"

    s3_origin_config {
      origin_access_identity = var.s3_origin_access_identity_path
    }
  }

  origin {
    domain_name = var.alb_dns_name
    origin_id   = "alb-origin"

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "http-only"
      origin_ssl_protocols   = ["TLSv1.2"]
    }
  }

  dynamic "origin" {
    for_each = var.apigw_domain_name == "" ? [] : [var.apigw_domain_name]
    content {
      domain_name = origin.value
      origin_id   = "apigw-origin"
      custom_origin_config {
        http_port              = 80
        https_port             = 443
        origin_protocol_policy = "https-only"
        origin_ssl_protocols   = ["TLSv1.2"]
      }
    }
  }

  default_cache_behavior {
    target_origin_id       = var.apigw_domain_name == "" ? "s3-origin" : "apigw-origin"
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD", "OPTIONS"]
    cached_methods         = ["GET", "HEAD", "OPTIONS"]

    forwarded_values {
      query_string = true
      headers      = ["Authorization", "Content-Type", "Origin", "Accept", "Accept-Language", "User-Agent", "Host"]
      cookies {
        forward = "all"
      }
    }

    # SPA用のURLリライトはSSR未使用時のみ有効化
    dynamic "function_association" {
      for_each = var.lambda_edge_origin_request_arn == "" && var.apigw_domain_name == "" ? [1] : []
      content {
        event_type   = "viewer-request"
        function_arn = aws_cloudfront_function.url_rewrite.arn
      }
    }

    # Lambda@Edge SSR（使用時のみ）
    dynamic "lambda_function_association" {
      for_each = var.apigw_domain_name == "" && var.lambda_edge_origin_request_arn != "" ? [var.lambda_edge_origin_request_arn] : []
      content {
        event_type   = "origin-request"
        lambda_arn   = lambda_function_association.value
        include_body = true
      }
    }
  }

  ordered_cache_behavior {
    path_pattern     = "/api/*"
    target_origin_id = "alb-origin"
    viewer_protocol_policy = "allow-all" # Allow HTTP for initial setup
    allowed_methods        = ["GET", "HEAD", "OPTIONS", "PUT", "POST", "PATCH", "DELETE"]
    cached_methods         = ["GET", "HEAD", "OPTIONS"]

    forwarded_values {
      query_string = true
      headers      = ["Authorization", "Content-Type", "Origin"]
      cookies {
        forward = "all"
      }
    }
  }

  # Serve Next static assets from S3 with caching
  ordered_cache_behavior {
    path_pattern           = "/_next/static/*"
    target_origin_id       = "s3-origin"
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]

    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
  }

  # 拡張子を含むリクエスト（例: *.ico, *.png, *.js, *.css）はS3から配信（SSRを通さない）
  ordered_cache_behavior {
    path_pattern           = "*.*"
    target_origin_id       = "s3-origin"
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]

    forwarded_values {
      query_string = true
      cookies {
        forward = "none"
      }
    }
  }

  # Next.js 画像最適化（APIGW経由で実行）
  ordered_cache_behavior {
    path_pattern           = "/_next/image*"
    target_origin_id       = var.apigw_domain_name == "" ? "s3-origin" : "apigw-origin"
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]

    forwarded_values {
      query_string = true
      cookies {
        forward = "none"
      }
    }
  }

  # Next.js データフェッチ（APIGWへ）
  ordered_cache_behavior {
    path_pattern           = "/_next/data/*"
    target_origin_id       = var.apigw_domain_name == "" ? "s3-origin" : "apigw-origin"
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]

    forwarded_values {
      query_string = true
      cookies {
        forward = "all"
      }
    }
  }

  # SSR demo path handled by Lambda@Edge without viewer-request rewrite（APIGW利用時は無効化）
  ordered_cache_behavior {
    path_pattern           = "/ssr/*"
    target_origin_id       = var.apigw_domain_name == "" ? "s3-origin" : "apigw-origin"
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD", "OPTIONS"]
    cached_methods         = ["GET", "HEAD", "OPTIONS"]

    forwarded_values {
      query_string = true
      headers      = ["Authorization", "Content-Type", "Origin", "Accept", "Accept-Language", "User-Agent", "Host"]
      cookies {
        forward = "all"
      }
    }

    dynamic "lambda_function_association" {
      for_each = var.apigw_domain_name == "" && var.lambda_edge_origin_request_arn != "" ? [var.lambda_edge_origin_request_arn] : []
      content {
        event_type   = "origin-request"
        lambda_arn   = lambda_function_association.value
        include_body = true
      }
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }

  custom_error_response {
    error_code         = 403
    response_code      = 200
    response_page_path = "/index.html"
  }

  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/index.html"
  }

  tags = {
    Environment = var.environment
  }
}
