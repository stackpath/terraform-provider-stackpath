# Create a new nginx container workload with network assignments deployed to Dallas, TX, USA.

resource "stackpath_compute_workload" "my-compute-workload" {
  # A human friendly name for the workload
  name = "My Compute Workload"
  # A DNS compatible value that uniquely identifies a workload
  slug = "my-compute-workload"

  # Define multiple labels on the workload container. These labels can be used
  # as label selectors when applying network policies.
  labels = {
    "role"        = "web-server"
    "environment" = "production"
  }

  # Define the network interfaces that should be provisioned for the workload
  # instances. 
  network_interface {
    network = "default"

    ip_families = ["IPv4", "IPv6"]

    # This is an example of what our implicit PRIMARY IPv4 assignment would look like
    assignments {
      # Unique identifier for the assignment
      slug = "primary-v4"

      # Some assignments may use the same mode, such as PRIMARY or ONE_TO_ONE_NAT
      # assignments for individual IPv4 and IPv6 addresses.
      mode = "PRIMARY"

      # This will dynamically create a allocation claim and an allocation
      # for the assignment in the network that the interface is attached to.
      allocation_claim_template {
        # IP Family of the allocation being claimed.
        ip_family = "IPv4"

        # The size of claim desired.
        prefix_length = 32

        # Reclaim policy defines what to do when the binding is removed.
        # Supported reclaim policies are RETAIN and DELETE.
        reclaim_policy {
          action = "DELETE"
        }

        # An allocation for the claim can be defined in three mutually exclusive ways:
        #
        # 1. Directly via reference in resource name format.
        # 2. Selecting across a set of existing allocations, allowing for the definition
        #    of "pools" of addresses to use for different purposes.
        # 3. Via the definition of a `template`, which will create an allocation for
        #    the claim if one does not already exist. The allocation will be identified
        #    by a slug unique to the claim.
        #
        allocation {
          template {
            # The allocation class will influence how an IP is configured on the platform,
            # such as where announcements are made, or which provisioner is responsible for
            # issuing addresses to the allocation.
            #
            # unicast and anycast classes are only valid for stack based allocations.
            # private is only valid for VPC Network allocations.
            allocation_class = "stackpath-edge/private"

            # IP Family of the allocation.
            ip_family = "IPv4"

            # The size of allocation desired.
            prefix_length = 32

            # Reclaim policy defines what to do when the binding is removed.
            reclaim_policy {
              action = "DELETE"
            }
          }
        }
      }
    }

    # This is an example of what our implicit ONE_TO_ONE_NAT IPv4 assignment would look like
    assignments {
      # Unique identifier for the assignment
      slug = "one-to-one-nat-v4"

      # Some assignments may use the same mode, such as PRIMARY or ONE_TO_ONE_NAT
      # assignments for individual IPv4 and IPv6 addresses.
      mode = "ONE_TO_ONE_NAT"

      # This will dynamically create a allocation claim and an allocation
      # for the assignment in the network that the interface is attached to.
      allocation_claim_template {
        # IP Family of the allocation being claimed.
        ip_family = "IPv4"

        # The size of claim desired.
        prefix_length = 32

        # Reclaim policy defines what to do when the binding is removed.
        # Supported reclaim policies are RETAIN and DELETE.
        reclaim_policy {
          action = "DELETE"
        }

        # An allocation for the claim can be defined in three mutually exclusive ways:
        #
        # 1. Directly via reference in resource name format.
        # 2. Selecting across a set of existing allocations, allowing for the definition
        #    of "pools" of addresses to use for different purposes.
        # 3. Via the definition of a `template`, which will create an allocation for
        #    the claim if one does not already exist. The allocation will be identified
        #    by a slug unique to the claim.
        #
        allocation {
          template {
            # The allocation class will influence how an IP is configured on the platform,
            # such as where announcements are made, or which provisioner is responsible for
            # issuing addresses to the allocation.
            #
            # unicast and anycast classes are only valid for stack based allocations.
            # private is only valid for VPC Network allocations.
            allocation_class = "stackpath-edge/unicast"

            # IP Family of the allocation.
            ip_family = "IPv4"

            # The size of allocation desired.
            prefix_length = 32

            # Reclaim policy defines what to do when the binding is removed.
            reclaim_policy {
              action = "DELETE"
            }
          }
        }
      }
    }

    # This is an example of what our implicit PRIMARY IPv6 assignment would look like
    assignments {
      # Unique identifier for the assignment
      slug = "primary-v6"

      # Some assignments may use the same mode, such as PRIMARY or ONE_TO_ONE_NAT
      # assignments for individual IPv4 and IPv6 addresses.
      mode = "PRIMARY"

      # This will dynamically create a allocation claim and an allocation
      # for the assignment in the network that the interface is attached to.
      allocation_claim_template {
        # IP Family of the allocation being claimed.
        ip_family = "IPv6"

        # The size of claim desired.
        prefix_length = 128

        # Reclaim policy defines what to do when the binding is removed.
        # Supported reclaim policies are RETAIN and DELETE.
        reclaim_policy {
          action = "DELETE"
        }

        # An allocation for the claim can be defined in three mutually exclusive ways:
        #
        # 1. Directly via reference in resource name format.
        # 2. Selecting across a set of existing allocations, allowing for the definition
        #    of "pools" of addresses to use for different purposes.
        # 3. Via the definition of a `template`, which will create an allocation for
        #    the claim if one does not already exist. The allocation will be identified
        #    by a slug unique to the claim.
        #
        allocation {
          template {
            # The allocation class will influence how an IP is configured on the platform,
            # such as where announcements are made, or which provisioner is responsible for
            # issuing addresses to the allocation.
            #
            # unicast and anycast classes are only valid for stack based allocations.
            # private is only valid for VPC Network allocations.
            allocation_class = "stackpath-edge/private"

            # IP Family of the allocation.
            ip_family = "IPv6"

            # The size of allocation desired.
            prefix_length = 128

            # Reclaim policy defines what to do when the binding is removed.
            reclaim_policy {
              action = "DELETE"
            }
          }
        }
      }
    }

    # This is an example of what our implicit ONE_TO_ONE_NAT IPv6 assignment would look like
    assignments {
      # Unique identifier for the assignment
      slug = "one-to-one-nat-v6"

      # Some assignments may use the same mode, such as PRIMARY or ONE_TO_ONE_NAT
      # assignments for individual IPv4 and IPv6 addresses.
      mode = "ONE_TO_ONE_NAT"

      # This will dynamically create a allocation claim and an allocation
      # for the assignment in the network that the interface is attached to.
      allocation_claim_template {
        # IP Family of the allocation being claimed.
        ip_family = "IPv6"

        # The size of claim desired.
        prefix_length = 128

        # Reclaim policy defines what to do when the binding is removed.
        # Supported reclaim policies are RETAIN and DELETE.
        reclaim_policy {
          action = "DELETE"
        }

        # An allocation for the claim can be defined in three mutually exclusive ways:
        #
        # 1. Directly via reference in resource name format.
        # 2. Selecting across a set of existing allocations, allowing for the definition
        #    of "pools" of addresses to use for different purposes.
        # 3. Via the definition of a `template`, which will create an allocation for
        #    the claim if one does not already exist. The allocation will be identified
        #    by a slug unique to the claim.
        #
        allocation {
          template {
            # The allocation class will influence how an IP is configured on the platform,
            # such as where announcements are made, or which provisioner is responsible for
            # issuing addresses to the allocation.
            #
            # unicast and anycast classes are only valid for stack based allocations.
            # private is only valid for VPC Network allocations.
            allocation_class = "stackpath-edge/unicast"

            # IP Family of the allocation.
            ip_family = "IPv6"

            # The size of allocation desired.
            prefix_length = 128

            # Reclaim policy defines what to do when the binding is removed.
            reclaim_policy {
              action = "DELETE"
            }
          }
        }
      }
    }
  }

  # Define secrets required for pulling private containers the workload spec.
  # This block can be repeated for each image
  # pull credential needed.
  # image_pull_credentials {
  #   docker_registry {
  #     username = "private-registry-username"
  #     password = "${file("./docker-registry-password.txt")}"
  #   }
  # }

  # Define an nginx container
  container {
    # Name that should be given to the container
    name = "app"
    # Nginx image to use for the container
    image = "nginx:latest"

    # Override the command that's used to execute the container. If this option
    # is not provided then the default entrypoint and command defined by the
    # docker image is used.
    # command = ["sleep", "infinity"]

    # Hardware resources to request of the container
    resources {
      requests = {
        # The number of CPU cores to allocate
        "cpu" = "1"
        # The amount of memory the container should use
        "memory" = "2Gi"
      }
    }

    # security_context {

    #  # Determine whether a process can request elevated privileges more
    #  # than its parent. Default is false
    #  allow_privilege_escalation = false
    #  # Should this container be run as non-root user
    #  run_as_non_root = false
    #  # What non-root user should this run as
    #  run_as_user = ""
    #  # What non-root group id should this run as
    #  run_as_group = ""
    #  # The set of linux security capabilities that your container
    #  # should have set or *dropped* when run. NET_ADMIN would be the most
    #  # common
    #  capabilities {
    #    add = [
    #      "NET_ADMIN",
    #    ]
    #    drop = [
    #      "NET_BROADCAST",
    #    ]
    #  }
    #}

    # The ports that should be publicly exposed on the containers.
    #
    # Warning, exposing these ports allows all internet traffic to access the
    # port. This option should be repeated for each port that should be exposed.
    #
    # Use network policies to add ACLs for the IPs that should be allowed to
    # access a container port.
    #
    # Ports are recorded in internal DNS SRV records for DNS-based service
    # discovery.
    port {
      # A name for the container port
      name = "http"
      # The port number that should be opened on the container
      port = 80
      # The protocol that should be allowed on the exposed port. This option
      # must be "TCP" (default) or "UDP".
      protocol = "TCP"
      # Whether or not the port is available from the public internet. This
      # defaults to false.
      #
      # The stackpath_compute_network_policy resource provides more granular
      # access to a port. See
      # https://www.terraform.io/docs/providers/stackpath/r/compute_network_policy.html
      # for more information.
      # enable_implicit_network_policy = false
    }

    # Environment variables exposed to the container are defined as key/value
    # pairs. You can define multiple environment variables for each container.
    # Each environment variable defined in a container must be unique within
    # that container.
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "this is a normal variable"
    }

    # You can also define sensitive environment variables using the secret_value
    # option. This values are not exposed in the API and are only exposed to
    # your container at runtime.
    env {
      key          = "MY_SECRET_VARIABLE"
      secret_value = "this is a secret variable"
    }
  }

  # Define the target configurations that selects where workloads should be
  # deployed.
  target {
    # Provide a name to the target. This name must be a valid DNS label as
    # described by RFC 1123. It can only contain the characters "a-z", "0-9",
    # "-", and ".".
    name = "us"
    # The scope of where the compute instance should be launched. The only
    # supported option is "cityCode".
    deployment_scope = "cityCode"

    # Create a single instance in each location that matches the selectors
    # defined below.
    min_replicas = 1

    # The maximum number of instance replicas that should be created in a target
    # deployment. This option is required when using auto-scaling options.
    # max_replicas = 10

    # The scaling configuration that should be used to determine when workload
    # instances within a target deployment should be scaled up or down. These
    # options are required when using the max_replicas option.
    # scale_settings {
    #   # Define a metric that should be used to determine whether the instances
    #   # within a target deployment should be scale up or down. When multiple
    #   # metrics are defined, the instances will be scaled when the threshold
    #   # for ANY of the metrics is reached, not when all defined metrics have
    #   # reached their threshold. The only supported metric for scaling is
    #   # "cpu".
    #   metrics {
    #     # The name of the metric that should be used to scale. The only
    #     # supported option is "cpu".
    #     metric = "cpu"
    #     # The average utilization (in percentage) when instances should be
    #     # scaled. For example, when set to 50%, the instances of a target
    #     # deployment are scaled up when the average CPU utilization of all
    #     # instances within the deployment is over 50%.
    #     average_utilization = 50
    #     # The average value (in resources) when instances should be scaled.
    #     # For example, when this is set to 500m, the instances of a target
    #     # deployment are scaled up when the average consumed CPU resources of
    #     # all instances within the deployment is over 500m, or half a CPU.
    #     average_value = "500m"
    #   }
    # }

    # Define a selector used to decide where workload instances should be
    # launched.
    selector {
      # Select the location to create an instance by the location's city code.
      # "cityCode" is the only supported option.
      key = "cityCode"
      # The operator to use when comparing values. Only the "in" operator is
      # supported.
      operator = "in"
      # The city code that instances should be created in. Cities are designated
      # by their IATA airport code.
      values = [
        "DFW",
      ]
    }
  }


  # Completely optional override of settings around the global
  # environment your container(s) are running inside
  # these settings are shared among your containers
  # container_runtime_settings {

  #  # How many seconds we should wait before killing your
  #  # container when restarting. Default is #0
  #  termination_grace_period_seconds = 30
  #  # Whether containers should be able to see each other
  #  share_process_namespace = false
  #  security_context {
  #    # User id used to execute entry point. Set to empty string to reset
  #    run_as_user = "100"
  #    # Group id used to execute entry point. Set to empty string to reset
  #    run_as_group = "100"
  #    # Indicates container should run as non-root user
  #    run_as_non_root = true
  #    supplemental_groups = [
  #      "101"
  #    ]
  #
  #   # Any SYSCTL settings you want to override,
  #   # use "setting"="value" 
  #   sysctl = {
  #      "net.core.rmem_max" = "10065408"
  #      "net.core.rmem_default" = "1006540"
  #     }
  #  }

  # If desired, override DNS/nameserver values here
  #  dns {
  #
  #    # You may repeat host_aliases as many times as you need
  #    # in order to override the lookup for the provided
  #    # 'hostnames' to resolve to the given 'address' 
  #    host_aliases {
  #      address = "192.168.3.4"
  #      hostnames = [ "domain.com" ]
  #    }

  #    # override resolv.conf (or equivalent) in your container
  #    # if you want to change the list of nameservers or
  #    # search order from our defaults. Note this will prevent
  #    # DNS discovery from working if you override nameservers
  #    resolver_config {
  #      # ordered list of nameserver ips
  #      nameservers = [ "8.8.8.8" ]
  #      # suffix(es) appended to hostname lookups
  #      search = [ "domain.com" ]
  #      # lookup options, timeout is a common one to set
  #      options = {
  #        timeout = "10"
  #      }
  #    }
  #  }
  #}
}