# Deploy an nginx container to the StackPath edge network and bind network
# policies to it.

# Configure the StackPath Terraform provider
variable "stackpath_stack_id" {}
variable "stackpath_client_id" {}
variable "stackpath_client_secret" {}

provider "stackpath" {
  stack_id      = var.stackpath_stack_id
  client_id     = var.stackpath_client_id
  client_secret = var.stackpath_client_secret
}

# This isn't necessary for creating the workload, but is useful for seeing the
# nginx instances' IP addresses
output "my-terraform-workload-instances" {
  value = {
    for instance in stackpath_compute_workload.my-terraform-nginx-workload.instances:
    instance.name => instance.external_ip_address
  }
}

# Create a new nginx container workload deployed to Seattle, WA, USA and New
# York City, NY, USA.
resource "stackpath_compute_workload" "my-terraform-nginx-workload" {
  name = "My Terraform Nginx Workload"
  slug = "my-terraform-nginx-workload"

  labels = {
    "role"        = "web-server"
    "environment" = "production"
  }

  network_interface {
    network = "default"
  }

  # Define an nginx container
  container {
    name  = "app"
    image = "nginx:latest"

    resources {
      requests = {
        "cpu"    = "1"
        "memory" = "2Gi"
      }
    }

    liveness_probe {
      period_seconds        = 60
      success_threshold     = 1
      failure_threshold     = 4
      initial_delay_seconds = 60
      http_get {
        port   = 80
        path   = "/"
        scheme = "HTTP"
      }
    }

    readiness_probe {
      tcp_socket {
        port = 80
      }
      period_seconds        = 60
      success_threshold     = 1
      failure_threshold     = 4
      initial_delay_seconds = 60
    }
  }

  target {
    name             = "us"
    deployment_scope = "cityCode"
    min_replicas     = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values   = [
        "SEA",
        "JFK",
      ]
    }
  }
}

# Create a new network policy that only applies to the nginx workload by tying
# it to the workload's slug.
resource "stackpath_compute_network_policy" "my-terraform-nginx-workload" {
  name = "My Terraform Nginx Network Policy"
  slug = "my-terraform-nginx-policy"

  # Apply the newtwork to the workload's instances only by their slug
  instance_selector {
    key      = "workload.platform.stackpath.net/workload-slug"
    operator = "in"
    values   = [stackpath_compute_workload.my-terraform-nginx-workload.slug]
  }

  priority     = 100
  policy_types = ["INGRESS", "EGRESS"]

  # Allow all inbound connections destined for port 80
  ingress {
    description = "Allow all outbound connections on both TCP and UDP"
    action      = "ALLOW"
    protocol {
      tcp_udp {
        destination_ports = [80]
      }
    }
    from {
      ip_block {
        cidr = "0.0.0.0/0"
      }
    }
  }

  # Allows all outbound connections to 0.0.0.0/0
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
