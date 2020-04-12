# Create a new Ubuntu virtual machine workload deployed to Seattle, WA, USA and
# New York City, NY, USA.
resource "stackpath_compute_workload" "my-compute-workload" {
  # A human friendly name for the workload
  name = "My Compute Workload"
  # A DNS compatible value that uniquely identifies a workload
  slug = "my-compute-workload"

  # Define multiple labels on the workload VM. These labels can be used as label
  # selectors when applying network policies.
  labels = {
    "role"        = "web-server"
    "environment" = "production"
  }

  # Define the network interfaces that should be provisioned for the workload
  # instances. StackPath only supports a "default" network for edge compute
  # workloads.
  network_interface {
    network = "default"
  }

  # Define an Ubuntu virtual machine
  virtual_machine {
    # Name that should be given to the VM
    name = "app"

    # StackPath image to use for the VM
    image = "stackpath-edge/ubuntu-1804-bionic:v201909061930"

    # Hardware resources dedicated to the VM
    resources {
      requests = {
        # The number of CPU cores to allocate
        "cpu" = "1"
        # The amount of memory the VM should have
        "memory" = "2Gib"
      }
    }

    # The ports that should be publicly exposed on the VM.
    #
    # Warning, exposing these ports allows all internet traffic to access the
    # port. This option should be repeated for each port that should be exposed.
    #
    # Use network policies to add ACLs for the IPs that should be allowed to
    # access a virtual machine port.
    #
    # Ports are recorded in internal DNS SRV records for DNS-based service
    # discovery.
    port {
      # A name for the VM port
      name = "http"
      # The port number that should be opened on the VM
      port = 80
      # The protocol that should be allowed on the exposed port. This option
      # must be "TCP" (default) or "UDP".
      protocol = "TCP"
      # Whether or not the port is available from the public internet. This
      # defaults to false.
      # enable_implicit_network_policy = false
    }

    # Cloud-init user data. Provide at least a public key
    # so you can SSH into VM instances after they're started.
    # See https://cloudinit.readthedocs.io/en/latest/topics/examples.html
    # for more information.
    user_data = <<EOT
#cloud-config
ssh_authorized_keys:
 - ssh-rsa <your public key>
EOT

    # Define a liveness probe that is used to determine the heath of a workload
    # instance. The workload instance is restarted when the liveness probe
    # begins failing.
    liveness_probe {
      # Execute the probe every 60 seconds
      period_seconds = 60
      # Mark the probe as successful after 1 successful probe
      success_threshold = 1
      # Mark the probe as failing after 4 failed checks
      failure_threshold = 4
      # Wait 60 seconds before starting probe checks to give the application
      # time to start up
      initial_delay_seconds = 60
      # Define the HTTP GET request that should be executed for the liveness
      # probe
      http_get {
        port = 80
        # Defaults to "/"
        path = "/"
        # Defaults to "HTTP"
        scheme = "HTTP"
        # Define HTTP headers that should be added to the HTTP GET request
        http_headers = {
          "content-type" = "application/json"
        }
      }
    }

    # Define a probe to determine when the instance is ready to serve traffic.
    readiness_probe {
      # This opens a TCP connection to port 80. The probe will only fail if it
      # cannot open a TCP connection to the port.
      tcp_socket {
        port = 80
      }
      # Execute the probe every 60 seconds
      period_seconds = 60
      # Mark the probe as successful after 1 successful probe
      success_threshold = 1
      # Mark the probe as failing after 4 failed checks
      failure_threshold = 4
      # Wait 60 seconds before starting probe checks to give the application
      # time to start up
      initial_delay_seconds = 60
    }

    # Mount an additional volume into the virtual machine. A volume can not be
    # mounted more than once to the same VM. A volume can be mounted to multiple
    # VMs in the same workload.
    volume_mount {
      slug       = "logging-volume"
      mount_path = "/var/log"
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
        "SEA",
        "JFK",
      ]
    }
  }

  # Provision a new additional volume that can be mounted to the containers and
  # virtual machines defined in the workload.
  volume_claim {
    # A human friendly label that can be updated
    name = "Logging volume"
    # A DNS compatible label that must not be changed
    slug = "logging-volume"
    # The resources used to configure the additional volume
    resources {
      requests = {
        storage = "100Gi"
      }
    }
  }
}
