---
layout: "stackpath"
page_title: "StackPath: stackpath_compute_vpc_network_subnet"
sidebar_current: "docs-stackpath-resource-compute-vpc-network-subnet"
description: |-
  User managed VPC network's subnet in Stackpath's edge computing platform
---

# stackpath\_compute\_vpc\_network\_subnet

User managed VPC network's subnet in Stackpath's edge computing platform.

## Example Usage

```hcl
resource "stackpath_compute_vpc_network_subnet" "net-subnet" {
  name        = "User defined VPC network subnet"
  slug        = "vpc-net-subnet"
  network_id  = "5daaaf31-682f-4287-8861-6ddb407b1bcf"
  # IPv4/IPv6 network subnet's prefix
  prefix      = "13.0.0.0/9"

  # Key/value pairs of arbitrary label names and values that can be referenced as selectors.
  labels = {
    "new-label" = "value1"
  }

  # Key/value pairs that define StackPath-specific vpc network subnet configuration.
  annotations = {
    "new-annotation" = "value1"
  }
} 
```

## Argument Reference

* `name` - (Required) A human readable name.
* `slug` - (Required) A programmatic name for the VPC network subnet.
* `prefix` - (Required) Subnet prefix to specify subnet range.
* `network_id` - (Required) VPC network's UUID ID or slug in which subnet needs to be created.
* `labels` - (Optional) Key/value pairs of arbitrary label names and values that can be referenced as selectors.
* `annotations` - (Optional) Key/value pairs that define StackPath-specific vpc network configuration.

## Import

StackPath compute VPC network's subnet can be imported by network-id/subnet-id formatted string. both network-id and subnet-id are UUID v4 formatted id e.g.

```
$ terraform import stackpath_compute_vpc_network.terraform 5daaaf31-682f-4287-8861-6ddb407b1bcf/9701c023-00f4-4f8b-b042-a272425fea7d
```
