terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = ">= 2.4"
    }
  }
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_dir  = var.function_source_dir
  output_path = "${path.module}/edge-ssr${var.zip_name_suffix}.zip"
}

resource "aws_iam_role" "lambda_edge_role" {
  name               = var.role_name_suffix == "" ? "${var.environment}-lambda-edge-role" : "${var.environment}-lambda-edge-${var.role_name_suffix}-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = ["edgelambda.amazonaws.com", "lambda.amazonaws.com"]
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "basic_execution" {
  role       = aws_iam_role.lambda_edge_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "edge_ssr" {
  function_name    = var.function_name != "" ? var.function_name : "${var.environment}-edge-ssr-use1"
  description      = "SSR handler for CloudFront (Lambda@Edge)"
  role             = aws_iam_role.lambda_edge_role.arn
  handler          = "index.handler"
  runtime          = "nodejs18.x"
  filename         = data.archive_file.lambda_zip.output_path
  publish          = true
  memory_size      = 512
  timeout          = 15
  skip_destroy     = var.skip_destroy

}

# Optional: allow Lambda@Edge to read/write ISR cache in a specified S3 bucket
resource "aws_iam_policy" "lambda_edge_cache_rw" {
  # count = var.cache_bucket_arn == "" ? 0 : 1
  count = 0 // destroy時用
  name  = var.role_name_suffix == "" ? "${var.environment}-lambda-edge-cache-rw" : "${var.environment}-lambda-edge-${var.role_name_suffix}-cache-rw"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "s3:ListBucket"
        ],
        Resource = [var.cache_bucket_arn]
      },
      {
        Effect = "Allow",
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:DeleteObject"
        ],
        Resource = ["${var.cache_bucket_arn}/*"]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_edge_cache_rw_attach" {
  # count      = var.cache_bucket_arn == "" ? 0 : 1
  count      = 0 // destroy時用
  role       = aws_iam_role.lambda_edge_role.name
  policy_arn = aws_iam_policy.lambda_edge_cache_rw[0].arn
}
