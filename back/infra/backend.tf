terraform {
  backend "s3" {
    bucket         = var.s3_bucket_terraform_state_name
    key            = "infra/terraform.tfstate"
    region         = "sa-east-1"
    encrypt        = true
    dynamodb_table = var.dynamodb_table_name
  }
}
