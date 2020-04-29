# This examples provides all the options available when creating  a network
# policy for StackPath edge compute workloads.
#
# Network policies leverage selectors to select the workload instances and
# networks that a policy should apply to. This allows you to leverage the same
# network policies for multiple workloads on your stack.
#
# As an example, create a policy that allows you to open port 22 on all
# workloads to a specific IP or CIDR range. You can also apply a general policy
# to apply to workloads based on roles defined through labels. For example,
# apply a role=web-server label to one or more workloads and create a network
# policy that allows port 80 on all instances that match the role=web-server
# label.
#
# You can create multiple network policies on your stack.
resource "stackpath_compute_network_policy" "my-compute-network-policy" {
  # A human friendly name
  name = "My compute network policy"
  # A DNS compatible label value that is unique to your stack. This value must
  # be RFC 1123 compliant (only contain "a-z", "0-9", "-", ".").
  slug = "my-compute-workload"
  # A text field that can be used to further describe the purpose of this
  # network policy.
  description = <<EOT
Open port 80 traffic to web servers and allow a specific CIDR range to access
the instance at port 9500.
EOT

  # Instance label selectors that determine which workload instances the network
  # policy should apply to. You can define multiple label selectors. If a
  # network policy doesn't have instance selectors then it applies to all
  # workload instances on the stack.
  #
  # Below is an example of a label selector that applies the policy all of the
  # instances created by the workload with the slug value "my-compute-workload".
  instance_selector {
    # Every workload instance is provisioned with a label of the workload slug
    # that created the instance. Apply the label selector to a specific workload
    # by leveraging the workload slug label.
    key = "workload.platform.stackpath.net/workload-slug"
    # The operator is the operation that should be applied to the value of the
    # label. Only the "in" operator is supported.
    operator = "in"
    # The values that the label value should be compared to
    values = ["my-compute-workload"]
  }

  # Below is an example of using a custom label defined in a workloa to select
  # all instances based on their role. For this example, assume workloads were
  # created that have a custom label of role="database". See
  # https://www.terraform.io/docs/providers/stackpath/r/compute_workload.html#selectorsFor
  # for more information.
  instance_selector {
    # The key is the label that should be used
    key = "role"
    # The operator is the operation that should be applied to the value of the
    # label. Only the "in" operator is supported.
    operator = "in"
    # The values that the label value should be compared to
    values = ["database"]
  }

  # The types of policies that should be configured, either "INGRESS" and/or
  # "EGRESS".
  policy_types = ["INGRESS"]

  # The priority that should be given to the network policy. The lower the
  # number, the higher the priority.
  #
  # Use the priority value 65534 to define multiple workload-specific policies
  # to avoid priority collisions.
  priority = 1000

  # An ingress policy defines the rules and actions taken for traffic received
  # by a workload instance. Policies can either allow or deny traffic.
  #
  # You can define multiple ingress policies for a network policy.
  ingress {
    # A human-readable description of what the ingress policy is for.
    description = "Allow port 80 traffic for all IPs"
    # The action that should be taken for this policy. You can either BLOCK or
    # ALLOW traffic based on this policy.
    action = "ALLOW"
    # The protocol block allows provides different configuration options based
    # on the connection's protocol.
    protocol {
      tcp {
        # Only apply the network policy to TCP connections on port 80.
        destination_ports = [80]
        # You can also configure the TCP ingress policy based on the source port
        # of the connection.
        # source_ports = [80]
      }

      # Configure the ingress policy for both UDP and TCP connections
      # tcp_udp {
      #   # Only apply the network policy to TCP connections on port 80.
      #   destination_ports = [80]
      #   # You can also configure the TCP ingress policy based on the source
      #   # port of the connection.
      #   source_ports = [80]
      # }

      # Configure the ingress policy for only UDP connections
      # udp {
      #   # Only apply the network policy to TCP connections on port 80.
      #   destination_ports = [80]
      #   # You can also configure the TCP ingress policy based on the source
      #   # port of the connection.
      #   source_ports = [80]
      # }

      # The ICMP, AH, GRE, and ESP protocols don't support configuration options.
      # icmp {}
      # ah {}
      # gre {}
      # esp {}
    }

    # Define the traffic sources the network policy applies to.
    #
    # Using this option you can control the traffic sources the policy applies
    # to. For example, apply the policy to a CIDR range like 10.0.8.0/24 or
    # apply the policy to inbound traffic from other workload instances.
    from {
      # The ip_block option applies the policy to a CIDR range like 10.0.8.0/24.
      # This option also provides an "except" option to exclude a CIDR range
      # from the policy.
      #
      # The following example applies the policy to the 10.1.0.0/16 subnet but
      # excludes the 10.1.1.0/24 subnet.
      ip_block {
        # The CIDR range this policy applies to
        cidr = "10.1.0.0/16"
        # Exclude this subnet from applying to this policy
        # except = ["10.1.1.0/24"]
      }

      # Multiple IP blocks can be defined.
      # ip_block {
      #   cidr = "203.0.113.0/24"
      # }

      # One of the more powerful features of network policies is applying them
      # to instances based on their labels. This allows you to build robust and
      # flexible network policies that can apply to any workloads created on
      # your stack.
      #
      # For example, if you have a database workload running with a label of
      # role=database-server and a workload running a web application with the
      # label role=web-app, then you can apply a network policy to the database
      # instances and configure an ingress policy using label selectors to allow
      # port 3306 from servers with the role=web-app label.
      # instance_selector {
      #   key      = "role"
      #   operator = "in"
      #   values   = ["web-app"]
      # }
    }
  }

  # An egress policy defines the rules and actions taken for traffic sent from a
  # workload instance. Configuration for ingress and egress policies are the
  # same, and you can define multiple egress blocks in a network policy.
  egress {
    description = "Allow all outbound connections on both TCP and UDP"
    action = "ALLOW"
  }
}
