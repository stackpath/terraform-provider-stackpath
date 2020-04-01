---
layout: "stackpath"
page_title: "StackPath: stackpath_compute_network_policy"
sidebar_current: "docs-stackpath-resource-compute-network-policy"
description: |-
  Network ingress and egress control of StackPath computing workloads.
---

# stackpath\_compute\_network\_policy

Network ingress and egress control of StackPath computing workloads.

## Example Usage

```hcl
resource "stackpath_compute_network_policy" "web-server" {
  name        = "Allow HTTP traffic to web servers"
  slug        = "web-servers-allow-http"
  description = "A network policy that allows HTTP access to instances with the web server role"
  priority    = 20000

  # Apply this network policy to every workload instance on the stack with the 
  # "web-server" role.
  instance_selector {
    key      = "role"
    operator = "in"
    values   = ["web-server"]
  }

  # Apply this network policy to specific workload instances. Use the key 
  # "workload.platform.stackpath.net/workload-slug" to target instances by slug 
  # or use the key "workload.platform.stackpath.net/workload-id" to target 
  # instances by ID.
  # 
  # Use the priority value 65534 to define multiple workload-specific policies 
  # to avoid priority collisions.
  instance_selector {
    key      = "workload.platform.stackpath.net/workload-slug"
    operator = "in"
    values   = ["my-workload-slug"]
  }

  policy_types = ["INGRESS"]

  ingress {
    action      = "ALLOW"
    description = "Allow port 80 traffic from all IPs"
    protocol {
      tcp {
        destination_ports = [80]
      }
    }
    from {
      ip_block {
        cidr = "0.0.0.0/0"
      }
    }
  }
}
```

## Argument Reference

* `name` - (Required) A human readable name.
* `slug` - (Required) A programmatic name for the network policy.
* `description` - (Optional) A brief description.
* `labels` - (Optional) Key/value pairs of arbitrary label names and values that can be referenced as selectors. 
* `annotations` - (Optional) Key/value pairs that define StackPath-specific network policy configuration.
* `instance_selector` - (Optional) A compute workload instance that the network policy applies to. A network policy with no selectors applies to all networks and all instances in the stack. See [Selectors](#selectors) below for details.
* `network_selector` - (Optional) A network that the network policy applies to. A network policy with no selectors applies to all networks and all instances in the stack. See [Selectors](#selectors) below for details.
* `policy_types` - (Required) A list of network policy types, either "INGRESS" and/or "EGRESS". 
* `priority` - (Required) A priority value between 1 and 65000. Higher priority network policies override lower priority policies, and priorities must be unique across the stack.
* `egress` - (Optional) Outbound networking information. See [Egress](#egress) below for details.
* `ingress` - (Optional) Inbound networking information. See [Ingress](#ingress) below for details.

### Egress

`egress` takes the following arguments:

* `action` - (Required) How a network policy treats outbound traffic, either "ALLOW" or "BLOCK".
* `description` - (Optional) A brief description.
* `protocol` - (Optional) Network protocol specific information. See [Network Protocols](#network-protocols) below for details.
* `to` - (Optional) Allow or block outbound traffic to the specified targets. See [Network Selectors](#network-selectors) below for details.

### Ingress

`ingress` takes the following arguments:

* `action` - (Required) How a network policy treats outbound traffic, either "ALLOW" or "BLOCK".
* `description` - (Optional) A brief description.
* `protocol` - (Optional) Network protocol specific information. See [Network Protocols](#network-protocols) below for details.
* `from` - (Optional) Allow or block inbound traffic from the specified targets. See [Network Selectors](#network-selectors) below for details.

### Network Protocols

`protocol` takes the following arguments:

* `ah` - (Optional) Allow or block the IPSec Authentication Header protocol. This argument block has no configuration.
* `esp` - (Optional) Allow or block the IPSec Encapsulating Security Payload protocol. This argument block has no configuration.
* `gre` - (Optional) Allow or block the Generic Routing Encapsulation protocol. This argument block has no configuration.
* `icmp` - (Optional) Allow or block the Internet Control Message Protocol. This argument block has no configuration.
* `tcp` - (Optional) Allow or block Transmission Control Protocol connections. See [Network Ports](#network-ports) below for details.
* `udp` - (Optional) Allow or block User Datagram Protocol connections. See [Network Ports](#network-ports) below for details.
* `tcp_udp` - (Optional) Allow or block both TCP and UDP connections. See [Network Ports](#network-ports) below for details.

### Network Ports

`tcp`, `udp`, and `tcp_udp` take the following arguments:

* `source_ports` - (Optional) A list of destination ports.
* `destination_ports` - (Optional) A list of destination ports.

### Network Selectors

`to` and `from` take the following arguments:

* `instance_selector` - (Optional) Target the given compute workload instances. See [Selectors](#selectors) below for details.
* `network_selector` - (Optional) Target the given networks. See [Selectors](#selectors) below for details.
* `ip_block` - (Optional) Target the given IP address blocks. See [IP Blocks](#ip-blocks) below for details. 

#### IP Blocks

`ip_block` takes the following arguments:

* `cidr` - (Required) A CIDR formatted subnet.
* `except` - (Optional) A list of CIDR formatted subnets to exclude from the `cidr` subnet.

### Selectors

`instance_selector` and `network_selector` take the following arguments:

* `key` - (Required) The name of the data that a selector is based on.
* `operator` - (Required) A logical operator to apply to a selector. Only the "in" operator is supported.
* `values` - (Required) Data values to look for in a label selector.

## Import

StackPath compute network policies can be imported by their UUID v4 formatted id. e.g.

```
$ terraform import stackpath_compute_network_policy.terraform bdb77768-2938-4ad8-a736-be5290add801
```
