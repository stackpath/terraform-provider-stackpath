# Create a new nginx container workload deployed to
# Seattle, WA, USA and New York City, NY, USA.
resource "stackpath_compute_workload" "my-compute-workload" {
  # A human friendly name for the workload
  name = "My Compute Workload"
  # A DNS compatible value that uniquely identifies
  # a workload.
  slug = "my-compute-workload"

  # Define multiple labels on the workload container. These
  # labels can be used as label selectors when applying
  # network policies.
  labels = {
    role        = "web-server"
    environment = "production"
  }

  # Define the network interfaces that should
  # be provisioned for the workload instances.
  # We currently only support a "default" network
  # for EdgeCompute workloads.
  network_interface {
    network = "default"
  }

  # Define the ImagePullSecrets that are required for
  # pulling the containers defined in the workload spec.
  # This option must not be provided for virtual machine
  # workloads. This block can be repeated for each image
  # pull credential needed.
  #
  # Uncomment if you need private registry support
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
    # Override the command that's used to execute
    # the container. If this option is not provided
    # the default entrypoint and command defined
    # by the docker image will be used.
    # command = []
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    # The ports that should be publicly exposed on
    # the containers. Warning, exposing these ports
    # will allow all internet traffic to access the
    # port. This option should be repeated for each
    # port that should be exposed.
    #
    # Use network policies to add ACLs for what IPs
    # should be allowed to access a container port.
    #
    # Ports are recorded in internal DNS SRV records
    # for DNS-based service discovery.
    port {
      # name that is given to the container port
      name = "http"
      # the port number that should be opened on
      # the container.
      port = 80
      # The protocol that should be allowed the
      # port that is expose. This option must be
      # "TCP" (default) or "UDP".
      protocol = "TCP"
      # Whether or not the port is available from the public
      # internet. This defaults to false.
      # enable_implicit_network_policy = false
    }

    # Environment variables are defined as key/value pairs. You can
    # define multiple environment variables for each container. Each
    # environment variable defined in a container must be unique within
    # that container.
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "this is a normal variable"
    }

    # You can also define sensitive environment variables using the
    # secret_value option. This values will not be exposed in the API
    # and will only be exposed to your container at runtime.
    env {
      key          = "MY_SECRET_VARIABLE"
      secret_value = "this is a secret variable"
    }

    # Define a liveness probe that is used to determine the
    # heath of an instance. The workload instance will be restarted
    # when the liveness probe begins failing.
    liveness_probe {
      # execute the liveness probe every 60 seconds
      period_seconds = 60
      # mark the probe as successful after 3 successful probes
      success_threshold = 1
      # mark the probe as failing after 4 failed liveness
      # probe checks
      failure_threshold = 4
      # wait 30 seconds before starting the liveness probe
      # checks to give the application time to start up
      initial_delay_seconds = 60
      # define the HTTP GET request that should be
      # executed for the liveness probe
      http_get {
        port = 80
        # defaults to "/"
        path = "/"
        # defaults to http
        scheme = "HTTP"
        # define the HTTP headers that should
        # be added to the HTTP GET request
        http_headers = {
          "content-type" = "application/json"
        }
      }
    }

    # Define a probe that is used to determine when the instance
    # should be considered ready to start serving traffic.
    readiness_probe {
      # This will open a TCP connection to port 80. The probe will
      # only fail if it cannot open a TCP connection to the configured
      # port.
      tcp_socket {
        port = 80
      }
      # execute the liveness probe every 60 seconds
      period_seconds = 60
      # mark the probe as successful after 3 successful probes
      success_threshold = 1
      # mark the probe as failing after 4 failed liveness
      # probe checks
      failure_threshold = 4
      # wait 30 seconds before starting the liveness probe
      # checks to give the application time to start up
      initial_delay_seconds = 60
      # define the HTTP GET request that should be
      # executed for the liveness probe
    }

    # Mount the additional volume into the container. A volume
    # must not be mounted more than once to the same container.
    # A volume can be mounted to multiple containers defined
    # in the same workload.
    volume_mount {
      slug       = "logging-volume"
      mount_path = "/var/log"
    }
  }

  # Define the target configurations that select
  # where workloads should be deployed.
  target {
    # Provider a name to the target, this name must be
    # a valid DNS label as described by RFC 1123, it can
    # only contain the characters 'a-z', '0-9', '-', '.'.
    name = "us"
    # The scope of where the compute instance should
    # be launched. The only option currently supported
    # is "cityCode".
    deployment_scope = "cityCode"

    # Create a single instance in each location that
    # matches the selectors defined below.
    min_replicas = 1

    # The maximum number of replicas of an instance that
    # should be created within a target deployment. This
    # option is required when using the auto-scaling options.
    # max_replicas = 10

    # The scaling configuration that should be used to determine
    # when workload instances within a target deployment should
    # be scaled up or down. These options are required when using
    # the max_replicas option.
    # scale_settings {
    #   # Defines a metric that should be used to determine whether
    #   # the instances within a target deployment should be scaled
    #   # up or down. When multiple metrics are defined, the instances
    #   # will be scaled when the threshold for ANY of the metrics is
    #   # reached, not when all defined metrics have reached their
    #   # threshold. Currently the only supported metric for scaling
    #   # is CPU.
    #   metrics {
    #     # The name of the metric that should be used to scale. The
    #     # only option currently supported is CPU.
    #     metric = "cpu"
    #     # The average utilization (in percent) when the instances
    #     # should be scaled. For example, when this is set to 50%,
    #     # the instances of a target deployment will be scaled up
    #     # when the average CPU utilization of all instances within
    #     # the deployment is over 50%.
    #     average_utilization = 50
    #     # The average value (in resources) when the instances should
    #     # be scaled. For example, when this is set to 500m, the
    #     # instances of a target deployment will be scaled up when the
    #     # average consumed CPU resources of all instances within the
    #     # deployment is over 500m or half a CPU.
    #     average_value = "500m"
    #   }
    # }

    # Define a selector that should be used to
    # decide where workload instances should be
    # launched. You can define multiple selectors
    # for a workload.
    selector {
      # Select the location to create an instance by
      # the city code of the location. This is currently
      # the only supported option.
      key = "cityCode"
      # The operator to use when comparing values
      operator = "in"
      # The city code that instances should be created in.
      values = [
        "SEA",
        "JFK",
      ]
    }
  }

  # Provision a new additional volume that can be
  # mounted to the containers and virtual machines
  # defined in the workload. You can define multiple
  # volume claims on a workload.
  volume_claim {
    # A human friendly label that can be updated
    name = "Logging volume"
    # A DNS compatible label that must not be changed
    # on a volume claim.
    slug = "logging-volume"
    # Define the resources that should be used to configure
    # the additional volume. The only option currently
    # supported is storage requests.
    resources {
      requests = {
        storage = "100Gi"
      }
    }
  }
}
