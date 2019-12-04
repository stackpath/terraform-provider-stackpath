package stackpath

import (
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/internal/models"
)

func flattenComputeWorkloadInstance(instance *models.Workloadv1Instance) map[string]interface{} {
	// We don't need to worry about the order of these containers or VMs
	// since workload instances are a computed resource and Terraform
	// doesn't need to create a diff.
	containers := make([]interface{}, 0, len(instance.Containers))
	for name, container := range instance.Containers {
		containers = append(containers, flattenComputeWorkloadContainer(name, container))
	}

	virtualMachines := make([]interface{}, 0, len(instance.VirtualMachines))
	for name, vm := range instance.VirtualMachines {
		virtualMachines = append(virtualMachines, flattenComputeWorkloadVirtualMachine(name, vm))
	}

	containerStatuses := make([]interface{}, 0, len(instance.ContainerStatuses))
	for _, status := range instance.ContainerStatuses {
		containerStatuses = append(containerStatuses, flattenComputeWorkloadContainerStatus(status))
	}

	return map[string]interface{}{
		"name":                instance.Name,
		"ip_address":          instance.IPAddress,
		"external_ip_address": instance.ExternalIPAddress,
		"reason":              instance.Reason,
		"message":             instance.Message,
		"phase":               instance.Phase,
		"container":           containers,
		"virtual_machine":     virtualMachines,
	}
}

func flattenComputeWorkloadContainerStatus(status *models.V1ContainerStatus) map[string]interface{} {
	var waiting []interface{}
	if status.Waiting != nil {
		waiting = []interface{}{
			map[string]interface{}{
				"reason":  status.Waiting.Reason,
				"message": status.Waiting.Message,
			},
		}
	}

	var terminated []interface{}
	if status.Terminated != nil {
		terminated = []interface{}{
			map[string]interface{}{
				"reason":      status.Terminated.Reason,
				"message":     status.Terminated.Message,
				"exit_code":   status.Terminated.ExitCode,
				"started_at":  status.Terminated.StartedAt.String(),
				"finished_at": status.Terminated.FinishedAt.String(),
			},
		}
	}

	var running []interface{}
	if status.Running != nil {
		running = []interface{}{
			map[string]interface{}{
				"started_at": status.Running.StartedAt.String(),
			},
		}
	}

	return map[string]interface{}{
		"name":          status.Name,
		"phase":         string(status.Phase),
		"started_at":    status.StartedAt.String(),
		"finished_at":   status.FinishedAt.String(),
		"ready":         status.Ready,
		"restart_count": status.RestartCount,
		"container_id":  status.ContainerID,
		"waiting":       waiting,
		"terminated":    terminated,
		"running":       running,
	}
}
