resource "aws_security_group" "ec2_sg" {
  name        = "${var.environment}-ec2-sg"
  description = "Allow HTTP/HTTPS/SSH inbound traffic"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.environment}-ec2-sg"
  }
}

# S3 Bucket for uploaded files
resource "aws_s3_bucket" "uploads_bucket" {
  bucket = "${var.environment}-simple-blog-uploads"

  tags = {
    Name = "${var.environment}-simple-blog-uploads"
  }
}

# IAM Role for EC2 to access S3
resource "aws_iam_role" "ec2_s3_role" {
  name = "${var.environment}-ec2-s3-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })

  tags = {
    Name = "${var.environment}-ec2-s3-role"
  }
}

resource "aws_iam_policy" "ec2_s3_policy" {
  name        = "${var.environment}-ec2-s3-policy"
  description = "IAM policy for EC2 to access S3 bucket for uploads"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:PutObject",
          "s3:GetObject",
          "s3:DeleteObject",
          "s3:ListBucket",
        ]
        Effect   = "Allow"
        Resource = [
          aws_s3_bucket.uploads_bucket.arn,
          "${aws_s3_bucket.uploads_bucket.arn}/*",
        ]
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ec2_s3_attach" {
  role       = aws_iam_role.ec2_s3_role.name
  policy_arn = aws_iam_policy.ec2_s3_policy.arn
}

resource "aws_iam_instance_profile" "ec2_s3_profile" {
  name = "${var.environment}-ec2-s3-profile"
  role = aws_iam_role.ec2_s3_role.name
}

resource "aws_instance" "backend_app" {
  ami           = var.ec2_ami
  instance_type = var.ec2_instance_type
  subnet_id     = var.public_subnet_ids[0]
  vpc_security_group_ids = [aws_security_group.ec2_sg.id]
  key_name      = var.ec2_key_pair_name
  iam_instance_profile = aws_iam_instance_profile.ec2_s3_profile.name

  tags = {
    Name = "${var.environment}-backend-app"
  }
}