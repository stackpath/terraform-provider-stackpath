---
layout: "stackpath"
page_title: "StackPath: stackpath_compute_vpc_network"
sidebar_current: "docs-stackpath-resource-compute-vpc-network"
description: |-
  User managed VPC network in Stackpath's edge computing platform
---

# stackpath\_compute\_vpc\_network

User managed VPC network in Stackpath's edge computing platform.

## Example Usage

```hcl
resource "stackpath_compute_vpc_network" "net" {
  name        = "User defined VPC network"
  slug        = "custom-vpc-net"

  # IPv4 network subnet
  root_subnet = "10.0.0.0/9"

  # IPv6 network subnet
  ipv6_subnet = "fc00::/64"

  # List of IPFamilies supported by VPC network
  ip_families = ["IPv4", "IPv6"]

  # Key/value pairs of arbitrary label names and values that can be referenced as selectors.
  labels = {
    "new-label" = "value1"
  }

  # Key/value pairs that define StackPath-specific vpc network configuration.
  annotations = {
    "new-annotation" = "value1"
  }
}
```

## Argument Reference

* `name` - (Required) A human readable name.
* `slug` - (Required) A programmatic name for the VPC network.
* `root_subnet` - (Optional) Network's IPv4 subnet. If not specified, default IPv4 subnet would be created in VPC network.
* `ipv6_subnet` - (Optional) Network's IPv6 subnet. If not specified, default IPv6 subnet would be created in VPC network.
* `ip_families` - (Optional) List of IP Families supported by VPC network.If VPC network is created with IPFamilies [IPv4] then network will support IPv4 only addresses, If IPFamilies is [IPv4, IPv6] then network is enabled for Dual stack networking and workloads can be requested with dual stack IP's on network interfaces requsted for this particular network. [IPv6] only VPC networks are not supported in current release.
* `labels` - (Optional) Key/value pairs of arbitrary label names and values that can be referenced as selectors.
* `annotations` - (Optional) Key/value pairs that define StackPath-specific vpc network configuration.

## Import

StackPath compute VPC network can be imported by their UUID v4 formatted id. e.g.

```
$ terraform import stackpath_compute_vpc_network.terraform bdb77768-2938-4ad8-a736-be5290add801
```
