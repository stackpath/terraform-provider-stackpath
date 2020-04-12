# Create a new nginx container workload deployed to Seattle, WA, USA and New
# York City, NY, USA.
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
  # instances. StackPath only supports a "default" network for edge compute
  # workloads.
  network_interface {
    network = "default"
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
    # command = []

    # Hardware resources to request of the contsainer
    resources {
      requests = {
        # The number of CPU cores to allocate
        "cpu" = "1"
        # The amount of memory the container should use
        "memory" = "2Gib"
      }
    }

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

    # Mount an additional volume into the container. A volume can not be mounted
    # more than once to the same container. A volume can be mounted to multiple
    # containers in the same workload.
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