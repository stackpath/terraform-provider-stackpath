# This examples provides all the options available when creating a network
# allocation for StackPath edge compute workloads.
#
# Network allocation leverage selectors to select the edge compute location.

resource "stackpath_compute_network_allocation" "my-compute-network-allocation" {
  # A human friendly name
  name = "My compute network allocation"
  # A DNS compatible label value that is unique to your stack. This value must
  # be RFC 1123 compliant (only contain "a-z", "0-9", "-", ".").
  slug = "my-compute-network-allocation"

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
