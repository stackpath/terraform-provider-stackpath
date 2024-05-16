# This examples provides all the options available when creating a network
# allocation for StackPath edge compute workloads.
#
# Network allocation leverage selectors to select the edge compute location.

resource "stackpath_compute_network_allocation" "my-compute-network-allocation-name-reference" {
  # A human friendly name
  name = "My compute network allocation name reference"
  # A DNS compatible label value that is unique to your stack. This value must
  # be RFC 1123 compliant (only contain "a-z", "0-9", "-", ".").
  slug = "my-compute-network-allocation-name-reference"

  # allocation class name
  allocation_class = "stackpath-edge/unicast"

  # allocation IP family, either IPv4 or IPv6
  ip_family = "IPv4"

  # allocation prefix length
  prefix_length = 32

  # allocation reclaim policy, only RETAIN action is supported from API
  reclaim_policy {
    action = "RETAIN"
  }

  selectors {
    # Apply the selectos to a specific edge compute location
    key = "cityCode"
    # The operator is the operation that should be applied to the value of the
    # label. Only the "in" operator is supported.
    operator = "in"
    # The values that the label value should be compared to
    values = ["EC4LAB01"]
  }
}

# This examples provides all the options available when creating a network
# allocation Claim with allocation name reference.
#
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
    name = "oc-testing-stack-3b4060/my-compute-network-allocation-name-reference"
  }

  depends_on = [
    stackpath_compute_network_allocation.my-compute-network-allocation-name-reference
  ]
}

# Allocation resource to be claimed using allocation selector in allocation claim spec.
resource "stackpath_compute_network_allocation" "my-compute-network-allocation-selector" {
  # A human friendly name
  name = "My compute network allocation selector"
  # A DNS compatible label value that is unique to your stack. This value must
  # be RFC 1123 compliant (only contain "a-z", "0-9", "-", ".").
  slug = "my-compute-network-allocation-selector"

  # add label to allocation which can be used later for allocation claim with
  # selector having match expression to match label
  labels = {
    "app" = "my-compute-network-allocation-selector"
  }

  # allocation class name
  allocation_class = "stackpath-edge/unicast"

  # allocation IP family, either IPv4 or IPv6
  ip_family = "IPv4"

  # allocation prefix length
  prefix_length = 32

  # allocation reclaim policy, only RETAIN action is supported from API
  reclaim_policy {
    action = "RETAIN"
  }

  selectors {
    # Apply the selectos to a specific edge compute location
    key = "cityCode"
    # The operator is the operation that should be applied to the value of the
    # label. Only the "in" operator is supported.
    operator = "in"
    # The values that the label value should be compared to
    values = ["EC4LAB01"]
  }
}


# This examples provides all the options available when creating a network
# allocation Claim with allocation selector.
#
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
        values = ["EC4LAB01"]
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

# This examples provides all the options available when creating a network
# allocation Claim with allocation template.
#
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
        values = ["EC4LAB01"]
      }
    }
  }
}