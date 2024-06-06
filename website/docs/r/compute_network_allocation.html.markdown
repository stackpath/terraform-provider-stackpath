---
layout: "stackpath"
page_title: "StackPath: stackpath_compute_network_allocation"
sidebar_current: "docs-stackpath-resource-compute-network-allocation"
description: |-
  IP allocations are used to allocate an IP address for immediate or future use with an allocation claim.
---

# stackpath\_compute\_network\_allocation

IP allocations are used to allocate an IP address for immediate or future use with an allocation claim.

## Example Usage

```hcl
# Allocation resource to be used to allocation IPv4 IP addresses.
resource "stackpath_compute_network_allocation" "my-compute-network-allocation-name-reference" {
  # A human friendly name
  name = "My compute network allocation name reference"
  # A DNS compatible label value that is unique to your stack. This value must
  # be RFC 1123 compliant (only contain "a-z", "0-9", "-", ".").
  slug = "my-compute-network-allocation-name-reference"

  # allocation class name
  # only stackpath-edge/private and stackpath-edge/unicast allocation classes are supported for now
  allocation_class = "stackpath-edge/unicast"

  # allocation IP family, either IPv4 or IPv6
  ip_family = "IPv4"

  # allocation prefix length
  # 32 and 128 are the only values supported for IPv4 and IPv6 respectively for now
  prefix_length = 32

  # allocation reclaim policy, only RETAIN action is supported from API
  reclaim_policy {
    action = "RETAIN"
  }

  selectors {
    # Apply the selectors to a specific edge compute location
    key = "cityCode"
    # The operator is the operation that should be applied to the value of the
    # label. Only the "in" operator is supported.
    operator = "in"
    # The values that the label value should be compared to
    values = ["DFW"]
  }
}
```

## Argument Reference

* `name` - (Required) A human readable name.
* `slug` - (Required) A programmatic name for the network allocation.
* `labels` - (Optional) Key/value pairs of arbitrary label names and values that can be referenced as selectors.
* `annotations` - (Optional) Key/value pairs that define StackPath-specific network allocation configuration.
* `allocation_class` - (Required) A IP allocation class to allocate IP from. Supported values are stackpath-edge/private and stackpath-edge/unicast.
* `ip_family` - (Required) An IP Family of the IPs being allocated. One of the IPv4 or IPv6 can be provided.
* `prefix_length` - (Required) A Prefix length of IP allocation. Currently only 32 and 128
prefix length values are supported for IPv4 and IPv6 respectively.
* `reclaim_policy` - (Required) A reclaim policy to be used for IP allocation. Only the RETAIN action is supported in reclaim policy specified in allocation resources being created from API.
* `selectors` - (Required) An edge location selector that the network allocation applies to. See [Selectors](#selectors) below for details.

### Selectors

`selectors` take the following arguments:

* `key` - (Required) The name of the data that a selector is based on.
* `operator` - (Required) A logical operator to apply to a selector. Only the "in" operator is supported.
* `values` - (Required) Data values to look for in a label selector.

## Import

StackPath compute network allocations can be imported by their stack-slug/allocation-slug formatted id. e.g.

```bash
$terraform import stackpath_compute_network_allocation.terraform bdb77768-2938-4ad8-a736-be5290add801/allocation-slug
```
