---
layout: "stackpath"
page_title: "StackPath: stackpath_compute_network_allocation_claim"
sidebar_current: "docs-stackpath-resource-compute-network-allocation-claim"
description: |-
  IP allocation claim required to claim IP from an allocation.
---

# stackpath\_compute\_network\_allocation\_claim

IP allocation claim used to claim IP from an allocation to network interfaces of StackPath computing workloads.

## Example Usage

```hcl
# Allocation claim resource using allocation.name spec to refer allocation and
# claim IP from it.
resource "stackpath_compute_network_allocation_claim" "allocation-claim-allocation-name-reference" {
  # A human friendly name
  name = "My allocation claim with name reference"
  # A DNS compatible label value that is unique to your stack. This value must
  # be RFC 1123 compliant (only contain "a-z", "0-9", "-", ".").
  slug = "allocation-claim-allocation-name-reference"

  # allocation IP family, either IPv4 or IPv6
  ip_family = "IPv4"

  # allocation prefix length
  prefix_length = 32

  # allocation reclaim policy, only RETAIN action is supported from API
  reclaim_policy {
    action = "RETAIN"
  }

  allocation {
    # name of the allocation to claim IP from
    # <replace stack slug with actual stack slug where resources are being created>
    name = "<stack-slug>/my-compute-network-allocation-name-reference"
  }
}

# Allocation claim resource using allocation.selector spec to refer allocation and
# claim IP from it.
resource "stackpath_compute_network_allocation_claim" "allocation-claim-allocation-selector" {
  # A human friendly name
  name = "My allocation claim with selector"
  # A DNS compatible label value that is unique to your stack. This value must
  # be RFC 1123 compliant (only contain "a-z", "0-9", "-", ".").
  slug = "allocation-claim-allocation-selector"

  # allocation IP family, either IPv4 or IPv6
  ip_family = "IPv4"

  # allocation prefix length
  prefix_length = 32

  # allocation reclaim policy, only RETAIN action is supported from API
  reclaim_policy {
    action = "RETAIN"
  }

  allocation {
    # use selector to match to cityCode and allocation label to claim from.
    selector {
      allocation_class = "stackpath-edge/unicast"
      match_expressions {
        key = "cityCode"
        operator = "in"
        values = ["DFW"]
      }

      match_expressions {
        key = "app"
        operator = "in"
        values = ["my-compute-network-allocation-selector"]
      }
    }
  }

  depends_on = [
    stackpath_compute_network_allocation.my-compute-network-allocation-selector
  ]
}

# Allocation claim resource using allocation.template spec to specify allocation
# specification. allocation is created internally with provided template spec and then
# used to claim IP from it.
resource "stackpath_compute_network_allocation_claim" "allocation-claim-allocation-template" {
  # A human friendly name
  name = "My allocation claim with allocation template"
  # A DNS compatible label value that is unique to your stack. This value must
  # be RFC 1123 compliant (only contain "a-z", "0-9", "-", ".").
  slug = "allocation-claim-allocation-template"

  # allocation IP family, either IPv4 or IPv6
  ip_family = "IPv4"

  # allocation prefix length
  prefix_length = 32

  # allocation reclaim policy, only RETAIN action is supported from API
  reclaim_policy {
    action = "RETAIN"
  }

  allocation {
    # use selector to match to cityCode and allocation label to claim from.
    template {
      allocation_class = "stackpath-edge/unicast"
      ip_family = "IPv4"
      prefix_length = 32
      reclaim_policy {
        action = "RETAIN"
      }

      selectors {
        key = "cityCode"
        operator = "in"
        values = ["DFW"]
      }
    }
  }
}

```

## Argument Reference

* `name` - (Required) A human readable name.
* `slug` - (Required) A programmatic name for the network allocation claim.
* `labels` - (Optional) Key/value pairs of arbitrary label names and values that can be referenced as selectors.
* `annotations` - (Optional) Key/value pairs that define StackPath-specific network allocation configuration.
* `ip_family` - (Required) A IP Family of the IPs being allocated. One of the IPv4 or IPv6 can be provided.
* `prefix_length` - (Required) A Prefix length of IP allocation. Currently only 32 and 128
prefix length values are supported for IPv4 and IPv6 respectively.
* `reclaim_policy` - (Required) A reclaim policy to be used for IP allocation. only RETAIN action is supported in reclaim policy specified in allocation resources being created from API.
* `allocation` - (Required) An allocation for the claim can be defined in three mutually exclusive ways:

1. Directly via reference in resource name format.
2. Selecting across a set of existing allocations, allowing for the definition of "pools" of addresses to use for different purposes.
3. Via the definition of a `template`, which will create an allocation for the claim if one does not already exist. The allocation will be identified by a slug unique to the claim.

See [Allocation](#allocation) below for details.

### Allocation

`allocation` take the one of the following arguments:

* `name` - (Optional) The name of allocation in resource name format.
* `selector` - (Required) A selector to select existing allocation using allocation_class
and match_expressions used to match allocation. See [Selector](#selector) below for details.
* `template` - (Required) A allocation resource template, it creates an allocation using
provided template if allocation does not exist. See [Template](#template) below for details.

### Selector

`selector` takes following arguments:

* `allocation_class` - (Required) A IP allocation class to allocate IP from. Supported values are stackpath-edge/private and stackpath-edge/unicast.
* `match_expressions` - (Required) A list of match expressions to match allocation. See [Match_Expressions](#match_expressions) below for details

### Match_Expressions

`match_expressions` take the following arguments:

* `key` - (Required) The name of the data that a selector is based on.
* `operator` - (Required) A logical operator to apply to a selector. Only the "in" operator is supported.
* `values` - (Required) Data values to look for in a label selector.

### Template

`template` take the following arguments:

* `allocation_class` - (Required) A IP allocation class to allocate IP from. Supported values are stackpath-edge/private and stackpath-edge/unicast.
* `ip_family` - (Required) A IP Family of the IPs being allocated. One of the IPv4 or IPv6 can be provided.
* `prefix_length` - (Required) A Prefix length of IP allocation. Currently only 32 and 128
prefix length values are supported for IPv4 and IPv6 respectively.
* `reclaim_policy` - (Required) A reclaim policy to be used for IP allocation. only RETAIN action is supported in reclaim policy specified in allocation resources being created from API.
* `selectors` - (Required) A edge location selector that the network allocation applies to. See [Selectors](#selectors) below for details.

### Selectors

`selectors` take the following arguments:

* `key` - (Required) The name of the data that a selector is based on.
* `operator` - (Required) A logical operator to apply to a selector. Only the "in" operator is supported.
* `values` - (Required) Data values to look for in a label selector.

## Import

StackPath compute network allocation claims can be imported by their stack-slug/allocationclaim-slug formatted id. e.g.

```bash
$terraform import stackpath_compute_network_allocation.terraform bdb77768-2938-4ad8-a736-be5290add801/allocationclaim-slug
```
