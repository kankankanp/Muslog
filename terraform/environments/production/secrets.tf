resource "aws_secretsmanager_secret" "app_secrets" {
  name = "production/app_secrets"
}

resource "aws_secretsmanager_secret_version" "app_secrets_version" {
  secret_id = aws_secretsmanager_secret.app_secrets.id
  secret_string = jsonencode({
    DB_PASSWORD           = var.db_password
    SPOTIFY_CLIENT_SECRET = var.spotify_client_secret
    GOOGLE_CLIENT_SECRET  = var.google_client_secret
  })
}
