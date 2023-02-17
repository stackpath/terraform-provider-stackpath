---
layout: "stackpath"
page_title: "StackPath: stackpath_compute_vpc_route"
sidebar_current: "docs-stackpath-resource-compute-vpc-route"
description: |-
  VPC Network's routing control for StackPath computing workloads.
---

# stackpath\_compute\_network\_policy

VPC Network's routing control for StackPath computing workloads.

## Example Usage

```hcl
resource "stackpath_compute_vpc_route" "route-wl" {
  name        = "route traffic from instance gateway"
  slug        = "route-traffic"
  network_id  = "5daaaf31-682f-4287-8861-6ddb407b1bcf"
  destination_prefixes = ["11.0.0.0/8"]
  gateway_selectors {
    interface_selectors {
      key = "workload.platform.stackpath.net/workload-slug"
      operator = "in"
      values = ["test"]
    }
  }
}
```

## Argument Reference

* `name` - (Required) A human readable name.
* `slug` - (Required) A programmatic name for the network policy.
* `network_id` - (Required) VPC network's UUID ID or slug in which subnet needs to be created.
* `destination_prefixes` - (Required) A list of destination network's prefixes to apply routing to.
* `gateway_selectors` - (Required) A routing gateway through which traffic is routed. currently workload instance interface is supported as routing gateway hence only interface_selector is supported under gateway_selectors. See [Gateway Selectors](#gateway-selectors) below for details.
* `labels` - (Optional) Key/value pairs of arbitrary label names and values that can be referenced as selectors.
* `annotations` - (Optional) Key/value pairs that define StackPath-specific network route configuration.

### Gateway Selectors

Gateway selector support interface selector to select interface as routing gateway.

`interface_selector` take the following arguments:

* `key` - (Required) The name of the data that a selector is based on.
* `operator` - (Required) A logical operator to apply to a selector. Only the "in" operator is supported.
* `values` - (Required) Data values to look for in a label selector.

## Import

StackPath compute VPC network's route can be imported by network-id/route-id formatted string. both network-id and route-id are UUID v4 formatted id e.g.

```
$ terraform import stackpath_compute_vpc_route.terraform bdb77768-2938-4ad8-a736-be5290add801/bdb77768-2938-4ad8-a736-be5290add802
```
