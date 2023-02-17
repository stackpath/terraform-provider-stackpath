---
layout: "stackpath"
page_title: "StackPath: stackpath_compute_workload"
sidebar_current: "docs-stackpath-resource-compute-workload"
description: |-
  A computing application deployed to StackPath's edge network.
---

# stackpath\_compute\_workload

A computing application deployed to StackPath's edge network.

## Example Usage
### Containers
```hcl
resource "stackpath_compute_workload" "my-compute-workload" {
  name = "my-compute-workload"
  slug = "my-compute-workload"

  annotations = {
    # request an anycast IP
    "anycast.platform.stackpath.net" = "true"
  }

  network_interface {
    network = "default"
  }

  container {
    # Name that should be given to the container
    name = "app"
    
    # Nginx image to use for the container
    image = "nginx:latest"
    
    # Override the command that's used to execute the container. If this option 
    # is not provided the default entrypoint and command defined by the docker 
    # image will be used.
    # command = []
    resources {
      requests = {
        "cpu"    = "1"
        "memory" = "2Gi"
      }
    }

    env {
      key   = "VARIABLE_NAME"
      value = "VALUE"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    max_replicas = 2
    scale_settings {
      metrics {
        metric = "cpu"
        # Scale up when CPU averages 50%.
        average_utilization = 50
      }
    }
    # Deploy these 1 to 2 instances in Dallas, TX, USA and Amsterdam, NL.
    deployment_scope = "cityCode"
    selector {
      key      = "cityCode"
      operator = "in"
      values   = [
        "DFW", "AMS"
      ]
    }
  }
}
```

### Virtual Machines
```hcl
resource "stackpath_compute_workload" "my-compute-workload" {
  name = "my-compute-workload"
  slug = "my-compute-workload"

  annotations = {
    # request an anycast IP
    "anycast.platform.stackpath.net" = "true"
  }

  network_interface {
    network = "default"
  }

  virtual_machine {
    # Name that should be given to the VM
    name = "app"
    
    # StackPath image to use for the VM
    image = "stackpath-edge/ubuntu-1804-bionic:v201909061930"

    # Cloud-init user data. 
    #
    # Provide at least a public key so you can SSH into VM instances after
    # they're started. See https://cloudinit.readthedocs.io/en/latest/topics/examples.html
    # for more information.
    user_data = <<EOT
#cloud-config
ssh_authorized_keys:
 - ssh-rsa <your public key>
EOT

    resources {
      requests = {
        "cpu"    = "1"
        "memory" = "2Gi"
      }
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    max_replicas = 2
    scale_settings {
      metrics {
        metric = "cpu"
        # Scale up when CPU averages 50%.
        average_utilization = 50
      }
    }
    # Deploy these 1 to 2 instances in Dallas, TX, USA and Amsterdam, NL.
    deployment_scope = "cityCode"
    selector {
      key      = "cityCode"
      operator = "in"
      values   = [
        "DFW", "AMS"
      ]
    }
  }
}
```

## Argument Reference

* `name` - (Required) A human readable name.
* `slug` - (Required) A programmatic name for the workload. Workload slugs are used to build the workload's instance names and cannot be changed after creation.
* `labels` - (Optional) Key/value pairs of arbitrary label names and values that can be referenced as [selectors](#selectors) by [network policies](/docs/providers/stackpath/r/compute_network_policy.html). 
* `annotations` - (Optional) Key/value pairs that define StackPath-specific workload configuration.
* `network_interface` - (Required) Networks to place the compute instance on. See [Network Interfaces](#network-interfaces) below for details.
* `image_pull_credentials` - (Optional) Credentials to pull container images with. See [Image Pull Credentials](#image-pull-credentials) below for details.
* `virtual_machine` - (Optional) Virtual machine configuration. StackPath supports a single virtual machine specification in a workload. At least one of `virtual_machine` or `container` must be provided. See [Virtual Machines](#virtual-machines) below for details.
* `container` - (Optional) Container configuration. At least one of `virtual_machine` or `container` must be provided. See [Containers](#containers) below for details.
* `volume_claim` - (Optional) Storage that is mounted to a compute workload's instances. See [Volume Claims](#volume-claims) below for details.
* `target` - (Required) How the compute workload should be deployed across the StackPath edge platform. See [Deployment Targets](#deployment-targets) below for details.

### Network Interfaces

`network_interfaces` supports the following arguments:

* `network` - (Required) A name that can be referenced by a [selector](#selectors) by [network policies](/docs/providers/stackpath/r/compute_network_policy.html). Both default and user created VPC networks are supported.
* `enable_one_to_one_nat` - (Optional) Boolean Flag to specify enabling of one to one
nat to interface VPC IP address. Default is true.
* `subnet` - (Optional) A name of IPv4 subnet to be used for IPv4 IP allocation for interface. subnet should belong to network specified in `network` field. if not specified then default IPv4 subnet will be used.
* `ipv6_subnet` - (Optional) A name of IPv6 subnet to be used for IPv6 IP allocation for interface. subnet should belong to network specified in `network` field. if not specified then default IPv6 subnet will be used.
* `ip_families` - (Optional) List of IP Families from which IP allocation should happen to network interface. Default is [IPv4]. Currently this supports IPv4-only [IPv4] and Dual stack- [IPv4, IPv6]. It does not suport IPv6 only- [IPv6] requests.
If Dual stack [IPv4, IPv6] is requested then VPC network specified in `network` field
is expected to enabled for Dual stack networking. Bu default all Default networks are enabled for Dual stack, any user created VPC network needs to be created with [IPv4, IPv6] IPFamilies to enable it for Dual stack networking.

### Image Pull Credentials

`image_pull_credentials` supports the following arguments:

* `docker_registry` - (Required) Authentication configuration needed to pull images from a Docker registry. See [Docker Registry Credentials](#docker-registry-credentials) below for details.

#### Docker Registry Credentials

`docker_registry` supports the following arguments:

* `server` - (Optional) The address of a Docker registry server. Defaults to "hub.docker.com".
* `username` - (Required) A username to connect the Docker registry.
* `password` - (Required) A password to connect to the Docker registry.
* `email` - (Optional) An email address to use with the Docker registry account.

### Virtual Machines

`virtual_machine` supports the following arguments:

* `name` - (Required) A virtual machine's name.
* `image` - (Required) The disk image to run as a virtual machine.
* `port` - (Optional) Network ports to expose from the virtual machine. Ports can also be used for internal DNS-based service discovery. See [Network Ports](#network-ports) below for details.
* `liveness_probe` - (Optional) Criteria to determine if the compute workload is online. See [Probes](#probes) below for details.
* `readiness_probe` - (Optional) Criteria to determine if the compute workload is ready to serve requests. See [Probes](#probes) below for details.
* `resources` - (Required) Hardware resources required by the virtual machine. See [Resources](#resources) below for details.
* `volume_mount` - (Optional) Storage volumes to mount in the virtual machine. See [Volume Mounts](#volume-mounts) below for details.
* `user_data` - (Optional) [Cloud-init](https://cloud-init.io/) user data.

### Containers

`container` supports the following arguments:

* `name` - (Required) A container's name.
* `image` - (Required) A container's image location.
* `command` - (Optional) The command to execute a container.
* `env` - (Optional) Environment variables to set in the container instance. See [Environment Variables](#environment-variables) below for details. 
* `port` - (Optional) Networking ports to expose from the container. Ports can also be used for internal DNS-based service discovery. See [Network Ports](#network-ports) below for details.
* `liveness_probe` - (Optional) Criteria to determine if the compute workload is online. See [Probes](#probes) below for details.
* `readiness_probe` - (Optional) Criteria to determine if the compute workload is ready to serve requests. See [Probes](#probes) below for details.
* `resources` - (Required) Hardware resources required by the container. See [Resources](#resources) below for details.
* `volume_mount` - (Optional) Storage volumes to mount in the container. See [Volume Mounts](#volume-mounts) below for details.

#### Environment Variables

`env` supports the following arguments:

* `key` - (Required) The environment variable name.
* `value` - (Optional) The environment variable value. One of `value` or `secret_value` must be provided.
* `secret_value` - (Optional) A sensitive environment variable value. This value cannot be read after it is set. One of `value` or `secret_value` must be provided.

### Network Ports

`port` supports the following arguments:

* `name` - (Required) The network port's name.
* `port` - (Required) The network port's number.
* `protocol` - (Optional) The network port's protocol, either "tcp" or "udp". Defaults to "tcp".
* `enable_implicit_network_policy` - (Optional) Whether or not the network port is accessible from the public Internet. Defaults to `false`. 

### Volume Claims

`volume_claim` supports the following arguments:

* `name` - (Required) A human readable name.
* `slug` - (Optional) A programmatic slug. Reference this slug when [mounting](#volume-mounts) the claim into a workload's instances.
* `resources` - (Required) Hardware resources to allocate to the volume claim. See [Resources](#resources) below for details.

### Probes

`liveness_probe` and `readiness_probe` take the following arguments:

* `http_get` - (Optional) HTTP request information. One of `http_get` or `tcp_socket` must be provided. See [HTTP probes](#http-probes) below for details
* `tcp_socket` - (Optional) TCP socket information. One of `http_get` or `tcp_socket` must be provided. See [TCP probes](#tcp-probes) below for details
* `initial_delay_seconds` - (Optional) The initial delay before the probe starts. Defaults to 0.
* `timeout_seconds` - (Optional) The number of seconds before the probe times out and is considered a failure. Defaults to 10.
* `period` - (Optional) The frequency of the probe. Defaults to 60.
* `success_threshold` - (Required) The minimum consecutive successes required before a probe is considered successful after a failure. This must be 1 for liveness probes.
* `failure_threshold` - (Required) The amount of failures seen before the probe is considered a failure.

#### HTTP Probes

`http_get` takes the following arguments:

* `path` - (Optional) The URL path to request from the application. Defaults to "/".
* `port` - (Required) The TCP port the HTTP service listens on.
* `scheme` - (Optional) The URL scheme to query the application with. Defaults to "http".
* `http_headers` - (Optional) HTTP header names and values to send to the HTTP service. 

#### TCP probes

`tcp_socket` takes the following arguments:

* `port` - (Required) The TCP port number to connect to. 

### Resources

`resources` takes the following arguments:

* `requests` - (Required) Key/value pairs of hardware resource types and values.

### Volume Mounts

`volume-mount` takes the following arguments:

* `slug` - (Required) The slug of the [volume claim](#volume-claim) to mount into the workload's instances.
* `mount_path` - (Required) The path the volume is mounted to in a workload's instances.

### Deployment Targets

`target` takes the following arguments: 

* `name` - (Required) A human readable name.
* `min_replicas` - (Required) The minimum number of instances that should be deployed to a target.
* `max_replicas` - (Optional) The maximum number of instances that should be deployed to a target.
* `scale_settings` - (Optional) How to auto-scale the number of instances in the deployment target. See [Scaling Settings](#scaling-settings) below for details.
* `deployment_scope` - (Optional) Criteria that defines a deployment target. Defaults to "cityCode".
* `selector` - (Required) The value of the deployment scope to deploy to. See [Selectors](#selectors) below for details.

#### Scaling Settings

`scale_settings` takes the following arguments:

* `metrics` - (Required) Scaling metrics. See [Scaling Metrics](#scaling-metrics) below for details.

##### Scaling Metrics 

`metrics` takes the following arguments:

* `metric` - (Required) A hardware metric to use as a scaling basis. Currently, only the "cpu" metric is supported.
* `average_utilization` - (Optional) The `metric`'s average utilization that should trigger scaling. One of `average_utilization` or `average_value` must be provided.
* `average_value` - (Optional) The `metric`'s average value that should trigger scaling. One of `average_utilization` or `average_value` must be provided.

#### Selectors

`selector` takes the following arguments:

* `key` - (Required) The name of the data that a selector is based on.
* `operator` - (Required) A logical operator to apply to a selector. Only the "in" operator is supported.
* `values` - (Required) Data values to look for in a label selector.

## Instances

A workload instance is a collection of containers or a virtual machine created based on the template provided in a workload. Instances are accessed via a `stackpath_compute_workload`'s computed `instances` field.

### Example Usage

```hcl
# Output a StackPath compute workload's instances' name, internal IP addresses, 
# and status
output "my-compute-workload-instances" {
  value = {
    for instance in stackpath_compute_workload.my-compute-workload.instances:
    instance.name => {
      ip_address = instance.external_ip_address
      phase      = instance.phase
    }
  }
}
```

### Instance Fields

* `name` - (Required) An instance's name. Names are generated from their corresponding workload's slug, followed by a unique hash.
* `metadata` - (Optional) Metadata associated with a running instance, including the workload's `labels` and [annotations](#annotations), both supplied by the user and generated by StackPath. 
* `location` - (Optional) The instance's physical location. See [Locations](#locations) below for details.
* `external_ip_address` - (Optional) An IPv4 address bound to the instance.
* `ip_address` - (Optional) An instance's internal IPv4 address.
* `external_ipv6_address` - (Optional) An IPv6 address bound to the instance.
* `ipv6_address` - (Optional) An instance's internal IPv6 address.
* `network_interface` - (Optional) A network interface bound to an instance. See [Instance Network Interfaces](#instance-network-interfaces) below for details.
* `virtual_machine` - (Optional) An instance's [virtual machine](#virtual-machines) specification. An instance has either a `virtual_machine` or `container`.
* `container` - (Optional) An instance's [container](#containers) specification. An instance has either a `virtual_machine` or `container`.
* `phase` - (Optional) An instance's current status, such as "STARTING", "RUNNING", "FAILED", or "STOPPED".
* `reason` - (Optional) A short reason why an instance is in its current phase.
* `message` - (Optional) A longer message with details why an instance is in its current phase.

### Locations

`locaton` has the following fields:

* `name` - (Optional) A location's name.
* `city`- (Optional) A location's city.
* `city_code` - (Optional) A city's [IATA code](https://en.wikipedia.org/wiki/IATA_airport_code).
* `subdivision` - (Optional) A location's subdivision.
* `subdivision_code` - (Optional) A subdivision's [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) code.
* `country` - (Optional) A location's country.
* `country_code` - (Optional) A country's [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.
* `region` - (Optional) A location's region.
* `region_code` - (Optional) A region's GeoIP code.
* `continent` - (Optional) A location's continent.
* `continent_code` - (Optional) A continent's GeoIP code.
* `latitude` - (Optional) A location's latitude coordinate.
* `longitude` - (Optional) A location's longitude coordinate.

### Instance Network Interfaces

`network_interface` has the following fields:

* `network` - (Required) The name of the [workload network interface](#network-interfaces).
* `ip_address` - (Required) A network interface's primary IPv4 address.
* `ip_address_aliases` - (Optional) Additional IPv4 addresses bound to a network interface.
* `gateway` - (Required) A network interface IPv4 subnet's gateway IP address.
* `ipv6_address` - (Optional) A network interface's primary IPv6 address.
* `ipv6_address_aliases` - (Optional) Additional IPv6 addresses bound to a network interface.
* `ipv6_gateway` - (Optional) A network interface IPv6 subnet's gateway IP address.

## Import

StackPath compute workloads can be imported by their UUID v4 formatted id. e.g.

```
$ terraform import stackpath_compute_workload.terraform bdb77768-2938-4ad8-a736-be5290add801
```
