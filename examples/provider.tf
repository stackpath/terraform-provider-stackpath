# Configure the StackPath Terraform provider.

# Allow external configuration of the provider by placing configuration
# directives into variables. Populate these variables either with environment
# variables or a terraform.tfvars file.
variable "stackpath_stack_id" {}
variable "stackpath_client_id" {}
variable "stackpath_client_secret" {}

provider "stackpath" {
  # The ID or the slug of stack to provision all StackPath services on
  stack_id = var.stackpath_stack_id
  # A StackPath API client ID
  client_id = var.stackpath_client_id
  # A StackPath API client secret
  client_secret = var.stackpath_client_secret
}