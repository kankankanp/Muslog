resource "aws_db_subnet_group" "main" {
  name       = "${var.project_name}-db-subnet-group-${var.environment}"
  subnet_ids = var.private_subnet_ids

  tags = {
    Environment = var.environment
  }
}

resource "aws_rds_cluster" "main" {
  cluster_identifier      = "${var.project_name}-db-cluster-${var.environment}"
  engine                  = "aurora-postgresql"
  engine_version          = "13.12"
  database_name           = var.db_name
  master_username         = var.db_username
  master_password         = var.db_password
  db_subnet_group_name    = aws_db_subnet_group.main.name
  vpc_security_group_ids  = [var.db_security_group_id]
  skip_final_snapshot     = true
  backup_retention_period = 7
  preferred_backup_window = "07:00-09:00"

  tags = {
    Environment = var.environment
  }
}

resource "aws_rds_cluster_instance" "main" {
  count              = var.db_instance_count
  identifier         = "${var.project_name}-db-instance-${var.environment}-${count.index}"
  cluster_identifier = aws_rds_cluster.main.id
  engine             = aws_rds_cluster.main.engine
  engine_version     = aws_rds_cluster.main.engine_version
  instance_class     = var.db_instance_class

  tags = {
    Environment = var.environment
  }
}
