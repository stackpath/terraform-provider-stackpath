# This examples provides all the options available when creating
# a network policy for StackPath EdgeCompute.
#
# Network policies leverage selectors to select the instances and
# networks that a policy should apply to. This allows you to leverage
# the same network policies for multiple workloads on your stack.
#
# As an example, you could create one policy that allows you to open
# port 22 on all workloads to a specific IP or CIDR range. You can also
# apply a general policy to apply to workloads based on roles defined
# through labels. For example, you could apply a role=web-server label
# to one or more workloads and create a new network policy that allows
# port 80 on all instances that have a label match of role=web-server.
#
# You can create multiple network policies on your stack.
resource "stackpath_compute_network_policy" "my-compute-network-policy" {
  # A human friendly name
  name = "My compute network policy"
  # A DNS compatible label value that is unique to your stack. This value
  # must be RFC 1123 compliant (only contain a-z, 0-9, -, .).
  slug = "my-compute-workload"
  # A text field that can be used to further describe the purpose of this
  # network policy.
  description = <<EOT
This network policy will open up port 80 traffic to web servers and allow
a specific CIDR range to access the instance at port 9500.
EOT

  # Instance label selectors that should be applied to determines
  # which instances the network policy should apply to. You can define
  # multiple label selectors. By not defining any instance selectors,
  # a network policy will default to applying to all workload instances.
  #
  # Below is an example of a label selector that will apply the policy
  # to all of the instances created by the workload with the slug value
  # of `my-compute-workload`.
  instance_selector {
    # Each workload instance is provisioned with a label of the workload
    # slug that created the instance. You can apply the label selector
    # to a specific workload by leveraging the workload slug label.
    key = "workload.platform.stackpath.net/workload-slug"
    # the operator is the operation that should be applied to the value
    # of the label.
    operator = "in"
    # the values that the label value should be compared to using the operator
    values = ["my-compute-workload"]
  }

  # Below is an example of using a custom label defined in a workload
  # to select all instances based on their role. For this example, we
  # assume workloads have been created that have a custom label of
  # role="database". For more information on adding custom labels
  # review the terraform Compute Workload reference file or visit the
  # developer documentation.
  instance_selector {
    # The key is the label that should be used
    key = "role"
    # the operator is the operation that should be applied to the value
    # of the label.
    operator = "in"
    # the values that the label value should be compared to using the operator
    values = ["database"]
  }

  # The types of policies that should be configured. The current
  # available options are "INGRESS" and "EGRESS".
  policy_types = ["INGRESS"]

  # The priority that should be given to the network policy. The
  # lower the number, the higher the priority.
  priority = 1000

  # An ingress policy defines the rules and actions that should be taken
  # for ingress traffic (traffic being received by your instance). Policies
  # can be created to either allow or deny traffic.
  #
  # You can define multiple ingress policies for your network policy.
  ingress {
    # A human friendly description of what the ingress policy is for.
    description = "Allow port 80 traffic for all IPs"
    # The action that should be taken for this policy. You can either
    # BLOCK or ALLOW traffic based on this policy.
    action = "ALLOW"
    # The protocol block allows provides different configuration options
    # for applying the ingress policy based on the protocol of the connection.
    protocol {
      # Config
      tcp {
        # Only apply the network policy to TCP connections on port 80.
        destination_ports = [80]
        # You can also configure the TCP ingress policy based on the source
        # port of the connection.
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

      # Uncomment this option to allow ICMP requests to the workload. This
      # block does not support any configuration options.
      # icpm {}
    }

    # Define the traffic sources the network policy should apply to.
    #
    # Using this option you can control what traffic sources the policy
    # should apply to. For example, you can apply the policy to a CIDR
    # range like 10.0.8.0/24 or apply the policy for network inbound
    # from other workload instances.
    from {
      # The CIDR block option provides the option to apply the policy to a
      # CIDR range like 10.0.8.0/24. This option allow provides an "except"
      # option to exclude a CIDR range from the policy. You can define multiple
      # CIDR block options for each ingress policy.
      #
      # The following example will apply the policy to the 10.1.0.0/16 subnet
      # but exclude the 10.1.1.0/24 subnet so that the policy does not apply to
      # traffic coming from the 10.1.1.0/24 subnet.
      ip_block {
        # the CIDR range this policy should apply to
        cidr = "10.1.0.0/16"
        # exclude this subnet from applying to this policy
        # except = ["10.1.1.0/24"]
      }

      # Adding a second CIDR subnet is as easy as defining a second CIDR block.
      # This CIDR block will apply the policy to IPs coming from the 205.185.200.0/24.
      # ip_block {
      #   cidr = "205.185.200.0/24"
      # }

      # One of the more powerful features of network policies is that you can
      # leverage instance selectors to apply an ingress policy to traffic coming
      # from instances based on those instances' labels. This allows you to build
      # robust and flexible network policies that can apply to any workloads created
      # on your stack.
      #
      # For example, if you have a database workload running with a label of
      # `role=database-server` and a workload running a web application with the
      # label `role=web-app`, you could apply a network policy to the database
      # instances and configure an ingress policy using label selectors to allow
      # port 3306 from servers with a label of role=web-app
      # instance_selector {
      #   key      = "role"
      #   operator = "in"
      #   values   = ["web-app"]
      # }
    }
  }

  # An egress policy defines the rules and actions that should be taken
  # for egress traffic (traffic being sent from your instance). Policies
  # can be created to either allow or deny traffic.
  #
  # You can define multiple egress policies for your network policy.
  egress {
    # A human friendly description of what the egress policy is for.
    description = "Allow all outbound connections on both TCP and UDP"
    # The action that should be taken for this policy. You can either
    # BLOCK or ALLOW traffic based on this policy.
    action = "ALLOW"
    # The protocol block allows provides different configuration options
    # for applying the egress policy based on the protocol of the connection.
    protocol {
      # Configure the egress policy and how it applies to TCP connections.
      # tcp {
      #   # Only apply the network policy to TCP connections destined for
      #   # port 80.
      #   destination_ports = [80]
      #   # You can also configure the TCP egress policy based on the source
      #   # port of the connection.
      #   source_ports = [80]
      # }

      # Configure the egress policy for both UDP and TCP connections
      # tcp_udp {
      #   # Only apply the network policy to TCP connections on port 80.
      #   destination_ports = [80]
      #   # You can also configure the TCP egress policy based on the source
      #   # port of the connection.
      #   source_ports = [80]
      # }

      # Configure the egress policy for only UDP connections
      # udp {
      #   # Only apply the network policy to TCP connections on port 80.
      #   destination_ports = [80]
      #   # You can also configure the TCP egress policy based on the source
      #   # port of the connection.
      #   source_ports = [80]
      # }

      # Uncomment this option to allow ICMP requests to the workload. This
      # block does not support any configuration options.
      # icpm {}
    }

    # Define the traffic sources the network policy should apply to.
    #
    # Using this option you can control what traffic sources the policy
    # should apply to. For example, you can apply the policy to a CIDR
    # range like 10.0.8.0/24 or apply the policy for network inbound
    # from other workload instances.
    from {
      # The CIDR block option provides the option to apply the policy to a
      # CIDR range like 10.0.8.0/24. This option allow provides an "except"
      # option to exclude a CIDR range from the policy. You can define multiple
      # CIDR block options for each egress policy.
      #
      # The following example will apply the policy to the 10.1.0.0/16 subnet
      # but exclude the 10.1.1.0/24 subnet so that the policy does not apply to
      # traffic coming from the 10.1.1.0/24 subnet.
      ip_block {
        # the CIDR range this policy should apply to
        cidr = "0.0.0.0/0"
        # exclude this subnet from applying to this policy
        # except = ["10.1.1.0/24"]
      }
    }
  }
}
