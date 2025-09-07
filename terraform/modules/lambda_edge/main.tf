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
  output_path = "${path.module}/edge-ssr.zip"
}

resource "aws_iam_role" "lambda_edge_role" {
  name               = "${var.environment}-lambda-edge-role"
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
  function_name    = "${var.environment}-edge-ssr"
  description      = "SSR handler for CloudFront (Lambda@Edge)"
  role             = aws_iam_role.lambda_edge_role.arn
  handler          = "index.handler"
  runtime          = "nodejs18.x"
  filename         = data.archive_file.lambda_zip.output_path
  publish          = true
  memory_size      = 512
  timeout          = 5
}
