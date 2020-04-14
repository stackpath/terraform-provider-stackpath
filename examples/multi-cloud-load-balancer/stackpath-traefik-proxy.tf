resource "stackpath_compute_workload" "traefik-lb" {
  name = "traefik-lb"
  slug = "traefik-lb"

  annotations = {
    # request an anycast IP for a workload
    "anycast.platform.stackpath.net" = "true"
  }

  network_interface {
    network = "default"
  }

  container {
    # Name that should be given to the container
    name = "app"
    # Nginx image to use for the container
    image = "scotwells/multi-cloud-traefik:latest"
    # Override the command that's used to execute
    # the container. If this option is not provided
    # the default entrypoint and command defined
    # by the docker image will be used.
    # command = []
    resources {
      requests = {
        "cpu"    = "1"
        "memory" = "2Gi"
      }
    }

    env {
      key   = "BACKEND_1"
      value = "http://${aws_instance.web_server_01.public_ip}/"
    }

    env {
      key   = "BACKEND_2"
      value = "http://${google_compute_instance.default.network_interface.0.access_config.0.nat_ip}/"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    max_replicas = 2
    scale_settings {
      metrics {
        metric = "cpu"
        # scale up when CPU averages 50%
        average_utilization = 50
      }
    }
    deployment_scope = "cityCode"
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "IAD", "JFK", "ORD", "ATL", "MIA",
        "DFW", "DEN", "SEA", "LAX", "SJC",
        "YYZ", "AMS", "LHR", "FRA", "WAW",
        "SIN", "GRU", "MEL", "NRT", "MAD",
        "ARN", "HKG",
      ]
    }
  }
}

output "traefik-anycast-ip" {
  value = replace(lookup(stackpath_compute_workload.traefik-lb.annotations, "anycast.platform.stackpath.net/subnets", ""), "/32", "")
}

output "traefik-workload-instances" {
  value = {
    for instance in stackpath_compute_workload.traefik-lb.instances:
    instance.name => {
      ip_address = instance.external_ip_address
      phase      = instance.phase
    }
  }
}
