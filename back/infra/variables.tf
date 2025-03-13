variable "s3_bucket_terraform_state_name" {
  type        = string
  description = "The name of the S3 bucket"
  default = "poc-ses-bucket-terraform-state"
}

variable "dynamo_db_name" {
  type        = string
  description = "The name of the DynamoDB table"
  default = "poc-ses-table"
}

variable "sender_email" {
  type        = string
  description = "The email address of the sender"
  default = ""
}

variable "s3_bucket_name" {
  type        = string
  description = "The name of the S3 bucket"
  default = ""
}

variable "dynamodb_table_name" {
  type        = string
  description = "The name of the DynamoDB table"
  default = ""
}