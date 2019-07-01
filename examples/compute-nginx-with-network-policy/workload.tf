# Define the variables that are needed to execute this plan.

variable "stackpath_stack" {
  description = "The StackPath ID of the Stack you want to create these resources in."
  type        = string
}

variable "stackpath_client_id" {
  description = "The client ID to use for API authentication. For more help, visit the get starting guide: https://stackpath.dev/docs/getting-started"
  type        = string
}

variable "stackpath_client_secret" {
  description = "The client secret to use for API authentication. For more help, visit the get starting guide: https://stackpath.dev/docs/getting-started"
  type        = string
}

provider "stackpath" {
  stack         = var.stackpath_stack
  client_id     = var.stackpath_client_id
  client_secret = var.stackpath_client_secret
}

# Create a new network policy that only applies to the nginx workload
resource "stackpath_compute_network_policy" "my-terraform-nginx-workload" {
  name = "My Terraform Nginx Network Policy"
  slug = "my-terraform-nginx-policy"

  instance_selector {
    key      = "workload.platform.stackpath.net/workload-slug"
    operator = "in"
    # grab the slug value that was created for the workload
    values = [stackpath_compute_workload.my-terraform-nginx-workload.slug]
  }

  priority = 100

  policy_types = ["INGRESS", "EGRESS"]

  # Create an ingress policy that allows all inbound connections destined for port 80
  ingress {
    description = "Allow all outbound connections on both TCP and UDP"
    action      = "ALLOW"
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

  # Create a policy that allows all outbound connections to 0.0.0.0/0
  egress {
    description = "Allow all outbound connections on both TCP and UDP"
    action      = "ALLOW"
    to {
      ip_block {
        cidr = "0.0.0.0/0"
      }
    }
  }
}

# Create a new compute workload that launches an nginx
# container to New York and JFK
resource "stackpath_compute_workload" "my-terraform-nginx-workload" {
  # A human friendly name for the workload
  name = "My Terraform Nginx Workload"
  # A DNS compatible value that uniquely identifies
  # a workload.
  slug = "my-terraform-nginx-workload"

  # Define multiple labels on the workload container. These
  # labels can be used as label selectors when applying
  # network policies.
  labels = {
    "role"        = "web-server"
    "environment" = "production"
  }

  # Define the network interfaces that should
  # be provisioned for the workload instances.
  # We currently only support a "default" network
  # for EdgeCompute workloads.
  network_interface {
    network = "default"
  }

  # Define an nginx container
  container {
    # Name that should be given to the container
    name = "app"
    # Nginx image to use for the container
    image = "nginx:latest"
    resources {
      requests = {
        "cpu"    = "1"
        "memory" = "2Gi"
      }
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
}
