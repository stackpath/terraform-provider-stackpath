package stackpath

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	compute "github.com/stackpath/terraform-provider-stackpath/stackpath/internal/client"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/internal/models"
)

func TestComputeWorkloadContainers(t *testing.T) {
	t.Parallel()

	workload := &models.V1Workload{}
	nameSuffix := strconv.Itoa(int(time.Now().Unix()))

	// By design, the StackPath API does not return image pull secrets to the
	// user when retrieving a workload. Expect to see an empty secret in the
	//result and suppress the diff error.
	emptyImagePullSecrets := regexp.MustCompile("(.*)image_pull_credentials.0.docker_registry.0.password:(\\s*)\"\" => \"secret registry password\"(.*)")

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testComputeWorkloadConfigContainerBasic(nameSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
				),
			},
			resource.TestStep{
				Config: testComputeWorkloadConfigContainerAddPorts(nameSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "https", "TCP", 443),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "some value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
				),
			},
			resource.TestStep{
				Config: testComputeWorkloadConfigContainerRemoveEnvVar(nameSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
				),
			},
			resource.TestStep{
				Config: testComputeWorkloadConfigContainerAddProbes(nameSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
				),
			},
			resource.TestStep{
				ExpectError: emptyImagePullSecrets,
				Config: testComputeWorkloadConfigContainerImagePullCredentials(nameSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckImagePullCredentials(workload, "docker.io", "my-registry-user", "developers@stackpath.com"),
				),
			},
			resource.TestStep{
				ExpectError: emptyImagePullSecrets,
				Config: testComputeWorkloadConfigAutoScalingConfiguration(nameSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckTargetAutoScaling(workload, "us", "cpu", 2, 4, 50),
				),
			},
			// TODO: there's a ordering issue where the order of the containers is shuffled when being read in from the API
			//   Need to ensure consistent ordering of containers when reading in state.
			//
			// resource.TestStep{
			// 	Config: testComputeWorkloadConfigContainerAddContainer(),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app-2", "nginx:v1.15.0"),
			// 	),
			// },
		},
	})
}

func TestComputeWorkloadContainersAdditionalVolume(t *testing.T) {
	t.Parallel()

	workload := &models.V1Workload{}
	nameSuffix := strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testComputeWorkloadConfigContainerAddVolumeMounts(nameSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo-volume", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadAdditionalVolume(workload, "volume", "10Gi"),
					testAccComputeWorlloadContainerVolumeMount(workload, "app", "volume", "/var/log"),
				),
			},
		},
	})
}

func TestComputeWorkloadVirtualMachines(t *testing.T) {
	t.Parallel()

	workload := &models.V1Workload{}
	nameSuffix := strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testComputeWorkloadConfigVirtualMachineBasic(nameSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.bar", workload),
					testAccComputeWorkloadCheckVirtualMachineImage(workload, "app", "stackpath-edge/centos-7:v201905012051"),
					testAccComputeWorkloadCheckVirtualMachinePort(workload, "app", "http", "TCP", 80),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
				),
			},
		},
	})
}

func testAccComputeWorkloadCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "stackpath_compute_workload" {
				continue
			}

			resp, err := config.compute.GetWorkload(&compute.GetWorkloadParams{
				StackID:    config.StackID,
				WorkloadID: rs.Primary.ID,
				Context:    context.Background(),
			}, nil)
			// Since compute workloads are deleted asyncronously, we want to look at the fact that
			// the deleteRequestedAt timestamp was set on the workload. This field is used to indicate
			// that the workload is being deleted.
			if err == nil && resp.Payload.Workload.Metadata.DeleteRequestedAt == nil {
				return fmt.Errorf("Compute workload still exists: %v", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccComputeWorlloadContainerVolumeMount(workload *models.V1Workload, containerName, volumeSlug, mountPath string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		container, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		}
		var mount *models.V1InstanceVolumeMount
		for _, v := range container.VolumeMounts {
			if v.Slug == volumeSlug {
				mount = v
				break
			}
		}
		if mount == nil {
			return fmt.Errorf("container volume mount not found: %s", volumeSlug)
		} else if mount.MountPath != mountPath {
			return fmt.Errorf("mount path '%s' does not match expected '%s'", mount.MountPath, mountPath)
		}
		return nil
	}
}

func testAccComputeWorkloadAdditionalVolume(workload *models.V1Workload, volumeName, size string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		var volume *models.V1VolumeClaim
		for _, v := range workload.Spec.VolumeClaimTemplates {
			if v.Name == volumeName {
				volume = v
				break
			}
		}
		if volume == nil {
			return fmt.Errorf("could not find volume: %s", volumeName)
		} else if volume.Spec.Resources.Requests["storage"] != size {
			return fmt.Errorf("volume size '%s' does not match expected '%s'", volume.Spec.Resources.Requests["storage"], size)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckVirtualMachinePort(workload *models.V1Workload, vmName, portName, protocol string, portNum int32) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if vm, found := workload.Spec.VirtualMachines[vmName]; !found {
			return fmt.Errorf("virtual machine was not found: %s", vmName)
		} else if port, found := vm.Ports[portName]; !found {
			return fmt.Errorf("virtual machine port not found: %s", portName)
		} else if port.Protocol != protocol {
			return fmt.Errorf("virtual machine port protocol '%s' does not match expected '%s'", port.Protocol, protocol)
		} else if port.Port != portNum {
			return fmt.Errorf("virtual machine port '%d' does not match expected '%d'", port.Port, portNum)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckVirtualMachineImage(workload *models.V1Workload, name, image string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if vm, found := workload.Spec.VirtualMachines[name]; !found {
			return fmt.Errorf("virtual machine was not found: %s", name)
		} else if vm.Image != image {
			return fmt.Errorf("virtual machine image '%s' does not match expected '%s'", vm.Image, image)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckNoImagePullCredentials(workload *models.V1Workload) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if len(workload.Spec.ImagePullCredentials) != 0 {
			return fmt.Errorf("unexpected image pull credentials set on the workload")
		}
		return nil
	}
}

func testAccComputeWorkloadCheckImagePullCredentials(workload *models.V1Workload, server, username, email string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if len(workload.Spec.ImagePullCredentials) == 0 {
			return fmt.Errorf("no image pull credentials set on the workload")
		} else if creds := workload.Spec.ImagePullCredentials[0]; creds.DockerRegistry.Server != server {
			return fmt.Errorf("image pull server '%s' does not match expected value '%s'", creds.DockerRegistry.Server, server)
		} else if creds.DockerRegistry.Email != email {
			return fmt.Errorf("image pull email '%s' does not match expected value '%s'", creds.DockerRegistry.Email, email)
		} else if creds.DockerRegistry.Username != username {
			return fmt.Errorf("image pull username '%s' does not match expected value '%s'", creds.DockerRegistry.Username, username)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckTarget(workload *models.V1Workload, targetName, scope, operator string, minReplicas int32, values ...string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		sort.Strings(values)
		if target, found := workload.Targets[targetName]; !found {
			return fmt.Errorf("target not found: %s", targetName)
		} else if target.Spec.DeploymentScope != scope {
			return fmt.Errorf("target scope '%s' does not match expected scope '%s'", target.Spec.DeploymentScope, scope)
		} else if deployment := target.Spec.Deployments; deployment.MinReplicas != minReplicas {
			return fmt.Errorf("target min replicas '%d' does not match expected '%d'", target.Spec.Deployments.MinReplicas, minReplicas)
		} else if selector := target.Spec.Deployments.Selectors[0]; selector.Operator != operator {
			return fmt.Errorf("target selector operator '%s' does not match expected operator '%s'", selector.Operator, operator)
		} else if sort.Strings(selector.Values); !reflect.DeepEqual(values, selector.Values) {
			return fmt.Errorf("target selector values %v do not match expected values %v", selector.Values, values)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerEnvVar(workload *models.V1Workload, containerName, envVar, value string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if envVarSpec, found := containerSpec.Env[envVar]; !found {
			return fmt.Errorf("environment variable not found: %s", envVar)
		} else if envVarSpec.Value != value {
			return fmt.Errorf(`environment variable '%s="%s"' does not match expected value '%s'`, envVar, envVarSpec.Value, value)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerEnvVarNotExist(workload *models.V1Workload, containerName, envVar string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if _, found := containerSpec.Env[envVar]; found {
			return fmt.Errorf("unexpected environment variable found: %s", envVar)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerPort(workload *models.V1Workload, containerName, portName, protocol string, port int32) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if portSpec, found := containerSpec.Ports[portName]; !found {
			return fmt.Errorf("port not found: %s", portName)
		} else if portSpec.Port != port {
			return fmt.Errorf("port number '%d' does not match expected port '%d'", portSpec.Port, port)
		} else if portSpec.Protocol != protocol {
			return fmt.Errorf("port protocol '%s' does not match expected protocol '%s'", portSpec.Protocol, protocol)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerPortNotExist(workload *models.V1Workload, containerName, portName string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if _, found := containerSpec.Ports[portName]; found {
			return fmt.Errorf("unexpected port found: %s", portName)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerImage(workload *models.V1Workload, containerName, image string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if containerSpec, found := workload.Spec.Containers[containerName]; !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if containerSpec.Image != image {
			return fmt.Errorf("container image '%s' does not match expected '%s'", containerSpec.Image, image)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckExists(name string, workload *models.V1Workload) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)
		found, err := config.compute.GetWorkload(&compute.GetWorkloadParams{
			WorkloadID: rs.Primary.ID,
			StackID:    config.StackID,
			Context:    context.Background(),
		}, nil)
		if err != nil {
			return fmt.Errorf("Could not retrieve workload: %v", err)
		}

		*workload = *found.Payload.Workload

		return nil
	}
}

func testAccComputeWorkloadCheckTargetAutoScaling(workload *models.V1Workload, targetName, metric string, minReplicas, maxReplicas, averageUtilization int32) resource.TestCheckFunc {
	return func(*terraform.State) error {
		target, found := workload.Targets[targetName]
		if !found {
			return fmt.Errorf("target not found: %s", targetName)
		} else if target.Spec.Deployments.MinReplicas != minReplicas {
			return fmt.Errorf("expected %d min replicas, got %d", minReplicas, target.Spec.Deployments.MinReplicas)
		} else if target.Spec.Deployments.MaxReplicas != maxReplicas {
			return fmt.Errorf("expected %d max replicas, got %d", maxReplicas, target.Spec.Deployments.MaxReplicas)
		} else if target.Spec.Deployments.ScaleSettings == nil {
			return fmt.Errorf("expected scale settings to be non-nil")
		}

		for _, m := range target.Spec.Deployments.ScaleSettings.Metrics {
			if m.Metric != metric {
				continue
			} else if m.AverageUtilization != averageUtilization {
				return fmt.Errorf("expected average utilization to be %d, got %d", averageUtilization, m.AverageUtilization)
			}
		}
		return nil
	}
}

func testComputeWorkloadConfigContainerBasic(suffix string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  network_interface {
    network = "default"
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
		requests = {
			cpu    = "1"
			memory = "2Gi"
		}
	}
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "value"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix)
}

func testComputeWorkloadConfigContainerAddPorts(suffix string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  network_interface {
    network = "default"
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
		requests = {
			cpu    = "1"
			memory = "2Gi"
		}
	}
	port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    port {
        name     = "https"
        port     = 443
        protocol = "TCP"
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "some value"
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix)
}

func testComputeWorkloadConfigContainerRemoveEnvVar(suffix string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  network_interface {
    network = "default"
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
		requests = {
			cpu    = "1"
			memory = "2Gi"
		}
	}
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix)
}

func testComputeWorkloadConfigContainerAddProbes(suffix string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  network_interface {
    network = "default"
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
		requests = {
			cpu    = "1"
			memory = "2Gi"
		}
	}
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    liveness_probe {
      period_seconds = 60
      success_threshold = 1
      failure_threshold = 4
      initial_delay_seconds = 60
      http_get {
        port = 80
        path = "/"
        scheme = "HTTP"
        http_headers = {
          "content-type" = "application/json"
        }
      }
    }
    readiness_probe {
      period_seconds = 60
      success_threshold = 1
      failure_threshold = 4
      initial_delay_seconds = 60
      timeout_seconds = 75
      tcp_socket {
        port = 80
      }
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix)
}

func testComputeWorkloadConfigContainerImagePullCredentials(suffix string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  network_interface {
    network = "default"
  }

  image_pull_credentials {
    docker_registry {
      server   = "docker.io"
      username = "my-registry-user"
      password = "secret registry password"
      email    = "developers@stackpath.com"
    }
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
		requests = {
			cpu    = "1"
			memory = "2Gi"
		}
	}
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix)
}

func testComputeWorkloadConfigContainerAddContainer(suffix string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  network_interface {
    network = "default"
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
  }

  container {
    name  = "app-2"
    image = "nginx:v1.15.0"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix)
}

func testComputeWorkloadConfigVirtualMachineBasic(suffix string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "bar" {
  name = "My Terraform Compute VM Workload - %s"
  slug = "terraform-vm-workload-%s"
  network_interface {
    network = "default"
  }

  virtual_machine {
    name  = "app"
	image = "stackpath-edge/centos-7:v201905012051"
	port {
		name     = "http"
		port     = 80
		protocol = "TCP"
	}
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
	user_data = <<EOT
package_update: true
packages:
- nginx
EOT
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix)
}

func testComputeWorkloadConfigContainerAddVolumeMounts(suffix string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo-volume" {
  name = "My Compute Workload Volume - %s"
  slug = "my-compute-workload-volume-%s"

  network_interface {
    network = "default"
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    volume_mount {
      slug = "volume"
      mount_path = "/var/log"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }

  volume_claim {
    name = "volume"
    slug = "volume"
    resources {
      requests = {
        "storage" = "10Gi"
      }
    }
  }
}`, suffix, suffix)
}

func testComputeWorkloadConfigAutoScalingConfiguration(suffix string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  network_interface {
    network = "default"
  }

  image_pull_credentials {
    docker_registry {
      server   = "docker.io"
      username = "my-registry-user"
      password = "secret registry password"
      email    = "developers@stackpath.com"
    }
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    max_replicas = 4
    scale_settings {
      metrics {
        metric = "cpu"
        average_utilization = 50
      }
    }
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix)
}
