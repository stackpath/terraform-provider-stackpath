# This examples provides all options to create allocation and allocation claims
# to claim that allocation.

# An allocation for the claim can be defined in three mutually exclusive ways:
#
# 1. Directly via reference in resource name format.
# 2. Selecting across a set of existing allocations, allowing for the definition
#    of "pools" of addresses to use for different purposes.
# 3. Via the definition of a `template`, which will create an allocation for
#    the claim if one does not already exist. The allocation will be identified
#    by a slug unique to the claim.

# This example covers all above ways by creating 2 allocations and using/referring
# those during allocation claim creation through allocation.name and allocation.selector
# resource specs respectively. and 3rd allocation claim is created using
# allocation.template which creates allocation internally if it does not exist.

# Network allocation leverage selectors to select the edge compute location and ip family
# to allocate IP of particular ip family(IPv4/IPv6).

# Allocation resource to be claimed using allocation.name in allocation claim spec.
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
    # Apply the selectos to a specific edge compute location
    key = "cityCode"
    # The operator is the operation that should be applied to the value of the
    # label. Only the "in" operator is supported.
    operator = "in"
    # The values that the label value should be compared to
    values = ["DFW"]
  }
}

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

  depends_on = [
    stackpath_compute_network_allocation.my-compute-network-allocation-name-reference
  ]
}

# Allocation resource to be claimed using allocation.selector in allocation claim spec.
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
    values = ["DFW"]
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
