package stackpath

import (
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/workload/workload_models"
)

func flattenComputeWorkloadInstance(instance *workload_models.Workloadv1Instance) map[string]interface{} {
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

	networkInterfaces := make([]interface{}, 0, len(instance.NetworkInterfaces))
	for _, networkInterface := range instance.NetworkInterfaces {
		networkInterfaces = append(networkInterfaces, flattenComputeWorkloadNetworkInterfaceStatus(networkInterface))
	}

	return map[string]interface{}{
		"name":                  instance.Name,
		"ip_address":            instance.IPAddress,
		"external_ip_address":   instance.ExternalIPAddress,
		"ipv6_address":          instance.IPV6Address,
		"external_ipv6_address": instance.ExternalIPV6Address,
		"reason":                instance.Reason,
		"message":               instance.Message,
		"phase":                 instance.Phase,
		"container":             containers,
		"virtual_machine":       virtualMachines,
		"network_interface":     networkInterfaces,
	}
}

func flattenComputeWorkloadNetworkInterfaceStatus(interfaceStatus *workload_models.Workloadv1NetworkInterfaceStatus) map[string]interface{} {
	var ipAddressAliases []interface{}
	if len(interfaceStatus.IPAddressAliases) > 0 {
		ipAddressAliases = flattenStringArray(interfaceStatus.IPAddressAliases)
	}

	var ipv6AddressAliases []interface{}
	if len(interfaceStatus.IPV6AddressAliases) > 0 {
		ipv6AddressAliases = flattenStringArray(interfaceStatus.IPV6AddressAliases)
	}

	return map[string]interface{}{
		"name":                 interfaceStatus.Network,
		"ip_address":           interfaceStatus.IPAddress,
		"ip_address_aliases":   ipAddressAliases,
		"gateway":              interfaceStatus.Gateway,
		"ipv6_address":         interfaceStatus.IPV6Address,
		"ipv6_address_aliases": ipv6AddressAliases,
		"ipv6_gateway":         interfaceStatus.IPV6Gateway,
	}
}

func flattenComputeWorkloadContainerStatus(status *workload_models.V1ContainerStatus) map[string]interface{} {
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

	var phase string
	if status.Phase != nil {
		phase = string(*status.Phase)
	}

	return map[string]interface{}{
		"name":          status.Name,
		"phase":         phase,
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
