# StackPath provider variables
variable "stackpath_stack_id" {}
variable "stackpath_client_id" {}
variable "stackpath_client_secret" {}

# AWS provider variables
variable "aws_region" {
  default = "us-east-1"
}
variable "aws_access_key" {}
variable "aws_secret_key" {}

# Google provider variables
variable "gcp_credentials_file" {}
variable "gcp_project" {}
variable "gcp_region" {}
variable "gcp_zone" {}

provider "stackpath" {
  stack_id      = var.stackpath_stack_id
  client_id     = var.stackpath_client_id
  client_secret = var.stackpath_client_secret
}

provider "aws" {
  region     = var.aws_region
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

provider "google" {
  credentials = file(var.gcp_credentials_file)
  project     = var.gcp_project
  region      = var.gcp_region
  zone        = var.gcp_zone
}
