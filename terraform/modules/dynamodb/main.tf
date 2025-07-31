resource "aws_dynamodb_table" "posts" {
  name         = "${var.environment}-posts"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    Name = "${var.environment}-posts-table"
  }
}

resource "aws_dynamodb_table" "users" {
  name         = "${var.environment}-users"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    Name = "${var.environment}-users-table"
  }
}
