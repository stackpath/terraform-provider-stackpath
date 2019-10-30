---
layout: "stackpath"
page_title: "Provider: StackPath"
sidebar_current: "docs-stackpath-index"
description: |-
  The StackPath provider is used to interact with resources on the StackPath edge platform.
---

# StackPath Provider

The StackPath provider is used to interact with resources on the StackPath edge platform.

The provider allows you manage your resources on the StackPath edge and integrate them with other Terraform-supported providers. It needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the StackPath Provider
provider "stackpath" {
  stack_id      = "${var.stackpath_stack_id}"
  client_id     = "${var.stackpath_client_id}"
  client_secret = "${var.stackpath_client_secret}"
}

# Create a new Edge Compute workload
resource "stackpath_compute_workload" "my-compute-workload" {
  # ...
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `stack_id` - (Required) This is the ID of stack that all new services are provisioned to. Stacks are folder-like organizational units on the StackPath platform and are typically used to organize services by project or user. Stack IDs are UUID v4 formatted strings. `stack` can be defined in either the provider definition or the `STACKPATH_STACK` environment variable.
* `client_id` - (Required) This is the API client ID of the StackPath user that will interact with Terraform. All services provisioned at StackPath through Terraform belong to their creating user. `client_id` can be defined in either the provider definition or the `STACKPATH_CLIENT_ID` environment variable.
* `client_secret` - (Required) This is the API client secret of the StackPath user that will interact with Terraform. Client secrets should be stored securely and not exposed to the public. `client_secret` can be defined in either the provider definition or the `STACKPATH_CLIENT_SECRET` environment variable.
