variable "cluster_name" {
  type        = string
  description = "Name of the cluster"
  # Bad: Missing variable validation block and default value
}

resource "google_sql_database_instance" "main" {
  name             = var.cluster_name
  database_version = "POSTGRES_14"
  region           = "us-central1"

  # Bad: Deletion protection disabled by default, violating security-by-default guideline
  deletion_protection = false

  settings {
    tier = "db-f1-micro"
  }
}
