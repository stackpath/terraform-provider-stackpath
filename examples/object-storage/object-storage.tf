# StackPath provider variables
variable "stackpath_stack_id" {}
variable "stackpath_client_id" {}
variable "stackpath_client_secret" {}
variable "bucket_label" {}

# Create provider
provider "stackpath" {
  stack_id      = var.stackpath_stack_id
  client_id     = var.stackpath_client_id
  client_secret = var.stackpath_client_secret
}

# Create bucket
resource "stackpath_object_storage_bucket" "bucket" {
  label = var.bucket_label
  region = "us-east-2"
  visibility = "PRIVATE"
}
