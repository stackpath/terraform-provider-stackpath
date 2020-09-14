package stackpath

import (
	"encoding/base64"
	"fmt"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/workload/workload_models"

	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// convert from the terraform data structure to the workload data structure we need for API calls
func convertComputeWorkload(data *schema.ResourceData) *workload_models.V1Workload {
	return &workload_models.V1Workload{
		Name: data.Get("name").(string),
		Slug: data.Get("slug").(string),
		Metadata: &workload_models.V1Metadata{
			Annotations: convertToStringMap(data.Get("annotations").(map[string]interface{})),
			Labels:      convertToStringMap(data.Get("labels").(map[string]interface{})),
		},
		Spec: &workload_models.V1WorkloadSpec{
			Containers:           convertComputeWorkloadContainers("container", data),
			VirtualMachines:      convertComputeWorkloadVirtualMachines("virtual_machine", data),
			NetworkInterfaces:    convertComputeWorkloadNetworkInterfaces(data.Get("network_interface").([]interface{})),
			ImagePullCredentials: convertComputeWorkloadImagePullCredentials("image_pull_credentials", data),
			VolumeClaimTemplates: convertComputeWorkloadVolumeClaims("volume_claim", data),
		},
		Targets: convertComputeWorkloadTargets(data.Get("target").([]interface{})),
	}
}

func convertComputeWorkloadVirtualMachines(prefix string, data *schema.ResourceData) workload_models.V1VirtualMachineSpecMapEntry {
	vms := make(workload_models.V1VirtualMachineSpecMapEntry, data.Get(prefix+".#").(int))
	for i, vm := range data.Get(prefix).([]interface{}) {
		vmData := vm.(map[string]interface{})
		vms[vmData["name"].(string)] = workload_models.V1VirtualMachineSpec{
			Image:          vmData["image"].(string),
			LivenessProbe:  convertComputeWorkloadProbe(fmt.Sprintf("%s.%d.liveness_probe", prefix, i), data),
			ReadinessProbe: convertComputeWorkloadProbe(fmt.Sprintf("%s.%d.readiness_probe", prefix, i), data),
			Ports:          convertComputeWorkloadPorts(fmt.Sprintf("%s.%d.port", prefix, i), data),
			Resources:      convertComputeWorkloadResourceRequirements(fmt.Sprintf("%s.%d.resources", prefix, i), data),
			VolumeMounts:   convertComputeWorkloadVolumeMounts(fmt.Sprintf("%s.%d.volume_mount", prefix, i), data),
			UserData:       base64.StdEncoding.EncodeToString([]byte(vmData["user_data"].(string))),
		}
	}
	return vms
}

func convertComputeWorkloadVolumeClaims(prefix string, data *schema.ResourceData) []*workload_models.V1VolumeClaim {
	volumes := make([]*workload_models.V1VolumeClaim, 0, data.Get(prefix+".#").(int))
	for i, vol := range data.Get(prefix).([]interface{}) {
		volumeData := vol.(map[string]interface{})
		volumes = append(volumes, &workload_models.V1VolumeClaim{
			Name: volumeData["name"].(string),
			Slug: volumeData["slug"].(string),
			Spec: &workload_models.V1VolumeClaimSpec{
				Resources: convertComputeWorkloadResourceRequirements(fmt.Sprintf("%s.%d.resources", prefix, i), data),
			},
		})
	}
	return volumes
}

func convertComputeWorkloadImagePullCredentials(prefix string, data *schema.ResourceData) workload_models.V1WrappedImagePullCredentials {
	credentials := make([]*workload_models.V1ImagePullCredential, 0, data.Get(prefix+".#").(int))
	for _, c := range data.Get(prefix).([]interface{}) {
		// only docker registry credentials are allowed for pull credentials, these credentials are guaranteed to exist
		creds := c.(map[string]interface{})["docker_registry"].([]interface{})[0].(map[string]interface{})
		credentials = append(credentials, &workload_models.V1ImagePullCredential{
			DockerRegistry: &workload_models.V1DockerRegistryCredentials{
				Server:   creds["server"].(string),
				Username: creds["username"].(string),
				Password: creds["password"].(string),
				Email:    creds["email"].(string),
			},
		})
	}
	return credentials
}

func convertComputeWorkloadTargets(data []interface{}) workload_models.V1TargetMapEntry {
	targets := make(workload_models.V1TargetMapEntry, len(data))
	for _, t := range data {
		target := t.(map[string]interface{})
		targets[target["name"].(string)] = workload_models.V1Target{
			Spec: &workload_models.V1TargetSpec{
				DeploymentScope: target["deployment_scope"].(string),
				Deployments: &workload_models.V1DeploymentSpec{
					MinReplicas:   int32(target["min_replicas"].(int)),
					MaxReplicas:   int32(target["max_replicas"].(int)),
					ScaleSettings: convertComputeWorkloadTargetScaleSettings(target["scale_settings"].([]interface{})),
					Selectors:     convertComputeMatchExpression(target["selector"].([]interface{})),
				},
			},
		}
	}
	return targets
}

func convertComputeWorkloadTargetScaleSettings(data []interface{}) *workload_models.V1ScaleSettings {
	if len(data) == 0 {
		return nil
	}

	settings := data[0].(map[string]interface{})

	metrics := make([]*workload_models.V1MetricSpec, len(settings["metrics"].([]interface{})))
	for i, metric := range settings["metrics"].([]interface{}) {
		metricData := metric.(map[string]interface{})
		metrics[i] = &workload_models.V1MetricSpec{
			Metric:             metricData["metric"].(string),
			AverageValue:       metricData["average_value"].(string),
			AverageUtilization: int32(metricData["average_utilization"].(int)),
		}
	}

	return &workload_models.V1ScaleSettings{
		Metrics: metrics,
	}
}

func convertComputeWorkloadNetworkInterfaces(data []interface{}) []*workload_models.V1NetworkInterface {
	interfaces := make([]*workload_models.V1NetworkInterface, len(data))
	for i, n := range data {
		interfaces[i] = &workload_models.V1NetworkInterface{
			Network: n.(map[string]interface{})["network"].(string),
		}
	}
	return interfaces
}

func convertComputeWorkloadContainers(prefix string, data *schema.ResourceData) workload_models.V1ContainerSpecMapEntry {
	// Keep track of the names of containers that exist in the new workload definition
	containerNames := make(map[interface{}]bool, data.Get(prefix+".#").(int))
	// First we need to add the containers that still exist within the workload definition
	containers := make(workload_models.V1ContainerSpecMapEntry, data.Get(prefix+".#").(int))
	for i, s := range data.Get(prefix).([]interface{}) {
		containerData := s.(map[string]interface{})
		// Track that we saw this container name, we check
		// for duplicates in the schema validation
		containerNames[containerData["name"]] = true

		log.Printf("[DEBUG] converting workload container '%v'", containerData["name"])
		containers[containerData["name"].(string)] = convertComputeWorkloadContainer(fmt.Sprintf("%s.%d", prefix, i), data)
	}

	// Don't perform any actions when our list of containers hasn't changed at all
	if data.HasChange("container") {
		oldContainers, _ := data.GetChange(prefix)

		// Now loop through all the bad containers and create a
		// blank entry to have the API remove the container
		for _, s := range oldContainers.([]interface{}) {
			containerData := s.(map[string]interface{})
			// When the container name was not seen in the new containers
			// it means the container was removed from the definition and
			// should be removed from the API.
			if !containerNames[containerData["name"]] {
				containers[containerData["name"].(string)] = workload_models.V1ContainerSpec{}
			}
		}
	}

	return containers
}

func convertComputeWorkloadContainer(prefix string, data *schema.ResourceData) workload_models.V1ContainerSpec {
	return workload_models.V1ContainerSpec{
		Image:          data.Get(prefix).(map[string]interface{})["image"].(string),
		Command:        convertToStringArray(data.Get(prefix + ".command").([]interface{})),
		Ports:          convertComputeWorkloadPorts(prefix+".port", data),
		Env:            convertComputeWorkloadEnvironmentVariables(prefix+".env", data),
		LivenessProbe:  convertComputeWorkloadProbe(prefix+".liveness_probe", data),
		ReadinessProbe: convertComputeWorkloadProbe(prefix+".readiness_probe", data),
		Resources:      convertComputeWorkloadResourceRequirements(prefix+".resources", data),
		VolumeMounts:   convertComputeWorkloadVolumeMounts(prefix+".volume_mount", data),
	}
}

func convertComputeWorkloadVolumeMounts(prefix string, data *schema.ResourceData) []*workload_models.V1InstanceVolumeMount {
	mounts := make([]*workload_models.V1InstanceVolumeMount, 0, data.Get(prefix+".#").(int))
	for _, m := range data.Get(prefix).([]interface{}) {
		mount := m.(map[string]interface{})
		mounts = append(mounts, &workload_models.V1InstanceVolumeMount{
			Slug:      mount["slug"].(string),
			MountPath: mount["mount_path"].(string),
		})
	}
	return mounts
}

func convertComputeWorkloadResourceRequirements(prefix string, data *schema.ResourceData) *workload_models.V1ResourceRequirements {
	return &workload_models.V1ResourceRequirements{
		Requests: convertToStringMap(data.Get(prefix).([]interface{})[0].(map[string]interface{})["requests"].(map[string]interface{})),
	}
}

func convertComputeWorkloadProbe(prefix string, data *schema.ResourceData) *workload_models.V1Probe {
	if !data.HasChange(prefix) && data.Get(prefix+".#").(int) == 0 {
		return nil
	} else if data.HasChange(prefix) && data.Get(prefix+".#").(int) == 0 {
		log.Printf("[DEBUG] removing probe from container: %v", prefix)
		// we are removing the probe so we should set the probe to an empty value
		return &workload_models.V1Probe{}
	}

	probe := data.Get(prefix + ".0").(map[string]interface{})
	if len(probe) == 0 {
		log.Printf("[WARNING] length of probe is 0: %v", prefix)
		return nil
	}

	log.Printf("[DEBUG] adding probe for container: %v", prefix)

	return &workload_models.V1Probe{
		FailureThreshold:    int32(probe["failure_threshold"].(int)),
		SuccessThreshold:    int32(probe["success_threshold"].(int)),
		InitialDelaySeconds: int32(probe["initial_delay_seconds"].(int)),
		PeriodSeconds:       int32(probe["period_seconds"].(int)),
		TimeoutSeconds:      int32(probe["timeout_seconds"].(int)),
		HTTPGet:             convertComputeWorkloadProbeHTTPGet(probe["http_get"].([]interface{})),
		TCPSocket:           convertComputeWorkloadProbeTCPSocket(probe["tcp_socket"].([]interface{})),
	}
}

func convertComputeWorkloadProbeTCPSocket(data []interface{}) *workload_models.V1TCPSocketAction {
	if len(data) == 0 {
		return nil
	}

	tcpAction := data[0].(map[string]interface{})

	return &workload_models.V1TCPSocketAction{
		Port: int32(tcpAction["port"].(int)),
	}
}

func convertComputeWorkloadProbeHTTPGet(data []interface{}) *workload_models.V1HTTPGetAction {
	if len(data) == 0 {
		return nil
	}

	httpAction := data[0].(map[string]interface{})

	return &workload_models.V1HTTPGetAction{
		Path:        httpAction["path"].(string),
		Port:        int32(httpAction["port"].(int)),
		Scheme:      strings.ToUpper(httpAction["scheme"].(string)),
		HTTPHeaders: convertToStringMap(httpAction["http_headers"].(map[string]interface{})),
	}
}

func convertComputeWorkloadPorts(prefix string, data *schema.ResourceData) workload_models.V1InstancePortMapEntry {
	portNames := make(map[interface{}]bool, data.Get(prefix+".#").(int))
	ports := make(workload_models.V1InstancePortMapEntry, data.Get(prefix+".#").(int))
	for _, s := range data.Get(prefix).([]interface{}) {
		portData := s.(map[string]interface{})
		// track the port names so we can keep track of what needs to be removed
		portNames[portData["name"]] = true
		ports[portData["name"].(string)] = workload_models.V1InstancePort{
			EnableImplicitNetworkPolicy: portData["enable_implicit_network_policy"].(bool),
			Port:                        int32(portData["port"].(int)),
			Protocol:                    strings.ToUpper(portData["protocol"].(string)),
		}
	}

	// Remove all ports that previously existed but were removed
	if data.HasChange(prefix) {
		old, _ := data.GetChange(prefix)
		for _, s := range old.([]interface{}) {
			portData := s.(map[string]interface{})
			if !portNames[portData["name"]] {
				ports[portData["name"].(string)] = workload_models.V1InstancePort{}
			}
		}
	}

	return ports
}

func convertComputeWorkloadEnvironmentVariables(prefix string, data *schema.ResourceData) workload_models.V1EnvironmentVariableMapEntry {
	envVarNames := make(map[interface{}]bool, data.Get(prefix+".#").(int))
	envVars := make(workload_models.V1EnvironmentVariableMapEntry, data.Get(prefix+".#").(int))
	for _, s := range data.Get(prefix).([]interface{}) {
		envVarData := s.(map[string]interface{})
		log.Printf("[DEBUG] setting environment variable '%s'", envVarData["key"])
		envVarNames[envVarData["key"]] = true
		envVars[envVarData["key"].(string)] = workload_models.V1EnvironmentVariable{
			Value:       envVarData["value"].(string),
			SecretValue: envVarData["secret_value"].(string),
		}
	}
	if data.HasChange(prefix) {
		old, _ := data.GetChange(prefix)
		for _, s := range old.([]interface{}) {
			envVarData := s.(map[string]interface{})
			if !envVarNames[envVarData["key"]] {
				log.Printf("[DEBUG] removing env var %v", envVarData["key"])
				envVars[envVarData["key"].(string)] = workload_models.V1EnvironmentVariable{}
			}
		}
	}
	return envVars
}

func flattenComputeWorkload(data *schema.ResourceData, workload *workload_models.V1Workload) error {
	if err := data.Set("name", workload.Name); err != nil {
		return fmt.Errorf("error setting name: %v", err)
	}

	if err := data.Set("slug", workload.Slug); err != nil {
		return fmt.Errorf("error setting slug: %v", err)
	}

	if err := data.Set("labels", flattenStringMap(workload.Metadata.Labels)); err != nil {
		return fmt.Errorf("error setting labels: %v", err)
	}

	if err := data.Set("annotations", flattenStringMap(workload.Metadata.Annotations)); err != nil {
		return fmt.Errorf("error setting annotations: %v", err)
	}

	if err := data.Set("network_interface", flattenComputeWorkloadNetworkInterfaces(workload.Spec.NetworkInterfaces)); err != nil {
		return fmt.Errorf("error setting network_interface: %v", err)
	}

	if err := data.Set("container", flattenComputeWorkloadContainers("container", data, workload.Spec.Containers)); err != nil {
		return fmt.Errorf("error setting container: %v", err)
	}

	if err := data.Set("image_pull_credentials", flattenComputeWorkloadImagePullCredentials("image_pull_credentials", data, workload.Spec.ImagePullCredentials)); err != nil {
		return fmt.Errorf("error setting image_pull_credentials: %v", err)
	}

	if err := data.Set("virtual_machine", flattenComputeWorkloadVirtualMachines("virtual_machine", data, workload.Spec.VirtualMachines)); err != nil {
		return fmt.Errorf("error setting virtual_machine: %v", err)
	}

	if err := data.Set("volume_claim", flattenComputeWorkloadVolumeClaims(workload.Spec.VolumeClaimTemplates)); err != nil {
		return fmt.Errorf("error setting volume_claim: %v", err)
	}

	if err := data.Set("target", flattenComputeWorkloadTargets("target", data, workload.Targets)); err != nil {
		return fmt.Errorf("error setting target: %v", err)
	}

	return nil
}

func flattenComputeWorkloadVolumeClaims(claims []*workload_models.V1VolumeClaim) []interface{} {
	flattened := make([]interface{}, len(claims))
	for i, claim := range claims {
		flattened[i] = map[string]interface{}{
			"name": claim.Name,
			"slug": claim.Slug,
			"resources": []interface{}{
				map[string]interface{}{
					"requests": map[string]interface{}{
						"storage": claim.Spec.Resources.Requests["storage"],
					},
				},
			},
		}
	}
	return flattened
}

// flattenComputeWorkloadVirtualMachines flattens the provided virtual machines
// with respect to the order of any virtual machines defined in the provided
// ResourceData. The prefix should be the flattened key of the list of virtual
// machines in the ResourceData.
func flattenComputeWorkloadVirtualMachines(prefix string, data *schema.ResourceData, vms workload_models.V1VirtualMachineSpecMapEntry) []interface{} {
	// Ensure we keep the original order so terraform doesn't mistaken things as out of sync
	ordered := make(map[string]int, data.Get(prefix+".#").(int))
	for i, k := range data.Get(prefix).([]interface{}) {
		// Set the name of the container in the map with its expected
		// index, container names are guaranteed to be unique
		ordered[k.(map[string]interface{})["name"].(string)] = i
	}
	flattened := make([]interface{}, data.Get(prefix+".#").(int))
	for name, vm := range vms {
		if index, found := ordered[name]; found {
			flattened[index] = flattenComputeWorkloadVirtualMachineOrdered(fmt.Sprintf("%s.%d", prefix, index), name, data, vm)
		} else {
			flattened = append(flattened, flattenComputeWorkloadVirtualMachine(name, vm))
		}
	}
	return flattened
}

// flattenComputeWorkloadVirtualMachineOrdered flattens a workload virtual machine but
// respects the ordering of the previous virtual machine entry. Ordering is important
// when dealing with updates to existing resources and accurate diffs are desired.
func flattenComputeWorkloadVirtualMachineOrdered(prefix, name string, data *schema.ResourceData, vm workload_models.V1VirtualMachineSpec) map[string]interface{} {
	decodedUserData, err := base64.StdEncoding.DecodeString(vm.UserData)
	if err != nil {
		// This error should never happen as the API only allows valid
		// base64 input and therefore should only ever output valid base64
		panic(err)
	}
	return map[string]interface{}{
		"name":            name,
		"image":           vm.Image,
		"port":            flattenComputeWorkloadPortsOrdered(prefix+".port", data, vm.Ports),
		"readiness_probe": flattenComputeWorkloadProbe(vm.ReadinessProbe),
		"liveness_probe":  flattenComputeWorkloadProbe(vm.LivenessProbe),
		"resources":       flattenComputeWorkloadResourceRequirements(vm.Resources),
		"volume_mount":    flattenComputeWorkloadVolumeMounts(vm.VolumeMounts),
		"user_data":       string(decodedUserData),
	}
}

// flattenComputeWorkloadVirtualMachine flattens the provided virtual machine
// spec as given. This implementation should only be used when the ordering of
// the returned virtual machines does not matter. When ordering is important,
// use flattenComputeWorkloadVirtualMachineOrdered.
func flattenComputeWorkloadVirtualMachine(name string, vm workload_models.V1VirtualMachineSpec) map[string]interface{} {
	decodedUserData, err := base64.StdEncoding.DecodeString(vm.UserData)
	if err != nil {
		// This error should never happen as the API only allows valid
		// base64 input and therefore should only ever output valid base64
		panic(err)
	}
	return map[string]interface{}{
		"name":            name,
		"image":           vm.Image,
		"port":            flattenComputeWorkloadPorts(vm.Ports),
		"readiness_probe": flattenComputeWorkloadProbe(vm.ReadinessProbe),
		"liveness_probe":  flattenComputeWorkloadProbe(vm.LivenessProbe),
		"resources":       flattenComputeWorkloadResourceRequirements(vm.Resources),
		"volume_mount":    flattenComputeWorkloadVolumeMounts(vm.VolumeMounts),
		"user_data":       string(decodedUserData),
	}
}

// flattenComputeWorkloadImagePullCredentials flattens the provided image pull
// credentials with respect to the order of any image pull credentials defined
// in the provided ResourceData. The prefix should be the flattened key of the
// list of image pull credentials in the ResourceData.
func flattenComputeWorkloadImagePullCredentials(prefix string, data *schema.ResourceData, credentials workload_models.V1WrappedImagePullCredentials) []interface{} {
	// Ensure we keep the original order so terraform doesn't mistaken things as out of sync
	ordered := make(map[string]int, data.Get(prefix+".#").(int))
	for i, k := range data.Get(prefix).([]interface{}) {
		// Grab the docker registry data set in the image pull
		// credentials, this guaranteed to be set.
		data := k.(map[string]interface{})["docker_registry"].([]interface{})[0].(map[string]interface{})
		// Set the order of the credentials based on the registry server
		ordered[data["server"].(string)] = i
	}
	creds := make([]interface{}, data.Get(prefix+".#").(int))
	for _, c := range credentials {
		data := map[string]interface{}{
			"docker_registry": []map[string]interface{}{
				{
					"server":   c.DockerRegistry.Server,
					"username": c.DockerRegistry.Username,
					"email":    c.DockerRegistry.Email,
				},
			},
		}
		if index, exists := ordered[c.DockerRegistry.Server]; exists {
			creds[index] = data
		} else {
			creds = append(creds, data)
		}
	}
	return creds
}

// flattenComputeWorkloadTargets flattens the provided workload targets with
// respect to the order of any targets defined in the provided ResourceData.
// The prefix should be the flattened key of the list of targets in the ResourceData.
func flattenComputeWorkloadTargets(prefix string, data *schema.ResourceData, targets workload_models.V1TargetMapEntry) []interface{} {
	// Ensure we keep the original order so terraform doesn't mistaken things as out of sync
	ordered := make(map[string]int, data.Get(prefix+".#").(int))
	for i, k := range data.Get(prefix).([]interface{}) {
		// Set the name of the container in the map with its expected
		// index, container names are guaranteed to be unique
		ordered[k.(map[string]interface{})["name"].(string)] = i
	}
	t := make([]interface{}, data.Get(prefix+".#").(int))
	for k, v := range targets {
		if index, found := ordered[k]; found {
			t[index] = flattenComputeWorkloadTarget(fmt.Sprintf("%s.%d", prefix, index), k, data, v)
		} else {
			t = append(t, flattenComputeWorkloadTarget(fmt.Sprintf("%s.%d", prefix, len(targets)), k, data, v))
		}
	}
	return t
}

// flattenComputeWorkloadTarget will flatten the provided workload target with
// respect to the original order of the target in the ResourceData. The prefix
// should be the flattened key of the target in the ResourceData.
func flattenComputeWorkloadTarget(prefix, name string, data *schema.ResourceData, target workload_models.V1Target) map[string]interface{} {
	return map[string]interface{}{
		"name":             name,
		"min_replicas":     target.Spec.Deployments.MinReplicas,
		"max_replicas":     target.Spec.Deployments.MaxReplicas,
		"deployment_scope": target.Spec.DeploymentScope,
		"scale_settings":   flattenComputeWorkloadTargetScaleSettings(prefix+".scale_settings", data, target.Spec.Deployments.ScaleSettings),
		"selector":         flattenComputeMatchExpressionsOrdered(prefix+".selector", data, target.Spec.Deployments.Selectors),
	}
}

func flattenComputeWorkloadTargetScaleSettings(prefix string, data *schema.ResourceData, settings *workload_models.V1ScaleSettings) []interface{} {
	if settings == nil {
		return nil
	}

	// Ensure we keep the original order so terraform doesn't mistaken things as out of sync
	ordered := make(map[string]int, data.Get(prefix+".0.metrics.#").(int))
	for i, k := range data.Get(prefix + ".0.metrics").([]interface{}) {
		// Set the name of the container in the map with its expected
		// index, container names are guaranteed to be unique
		ordered[k.(map[string]interface{})["metric"].(string)] = i
	}
	flattenedMetrics := make([]interface{}, data.Get(prefix+".0.metrics.#").(int))
	for _, metric := range settings.Metrics {
		if index, exists := ordered[metric.Metric]; exists {
			flattenedMetrics[index] = map[string]interface{}{
				"metric":              metric.Metric,
				"average_value":       metric.AverageValue,
				"average_utilization": metric.AverageUtilization,
			}
		} else {
			flattenedMetrics = append(flattenedMetrics, map[string]interface{}{
				"metric":              metric.Metric,
				"average_value":       metric.AverageValue,
				"average_utilization": metric.AverageUtilization,
			})
		}
	}

	return []interface{}{
		map[string]interface{}{
			"metrics": flattenedMetrics,
		},
	}
}

func flattenComputeWorkloadNetworkInterfaces(networkInterfaces []*workload_models.V1NetworkInterface) []interface{} {
	flattened := make([]interface{}, len(networkInterfaces))
	for i, n := range networkInterfaces {
		flattened[i] = map[string]interface{}{
			"network": n.Network,
		}
	}
	return flattened
}

// flattenComputeWorkloadContainers flattens the provided workload containers
// with respect to the order of any containers defined in the provided ResourceData.
// The prefix should be the flattened key of the list of containers in the ResourceData.
func flattenComputeWorkloadContainers(prefix string, data *schema.ResourceData, containers workload_models.V1ContainerSpecMapEntry) []interface{} {
	// Ensure we keep the original order so terraform doesn't mistaken things as out of sync
	ordered := make(map[string]int, data.Get(prefix+".#").(int))
	for i, k := range data.Get(prefix).([]interface{}) {
		// Set the name of the container in the map with its expected
		// index, container names are guaranteed to be unique
		ordered[k.(map[string]interface{})["name"].(string)] = i
	}
	flattened := make([]interface{}, data.Get(prefix+".#").(int))
	for name, container := range containers {
		// In the event that a container is added to a workload outside of
		// terraform, we need to support adding unknown containers to our state.
		if index, found := ordered[name]; found {
			flattened[index] = flattenComputeWorkloadContainerOrdered(fmt.Sprintf("%s.%d", prefix, index), name, data, container)
		} else {
			flattened = append(flattened, flattenComputeWorkloadContainer(name, container))
		}
	}
	return flattened
}

// flattenComputeWorkloadContainerOrdered flattens a workload container but respects
// the ordering of the previous container entry. Ordering is important when dealing
// with updates to existing resources and accurate diffs are desired.
func flattenComputeWorkloadContainerOrdered(prefix, name string, data *schema.ResourceData, container workload_models.V1ContainerSpec) map[string]interface{} {
	return map[string]interface{}{
		"name":            name,
		"image":           container.Image,
		"command":         flattenStringArray(container.Command),
		"port":            flattenComputeWorkloadPortsOrdered(prefix+".port", data, container.Ports),
		"env":             flattenComputeWorkloadEnvVarsOrdered(prefix+".env", data, container.Env),
		"readiness_probe": flattenComputeWorkloadProbe(container.ReadinessProbe),
		"liveness_probe":  flattenComputeWorkloadProbe(container.LivenessProbe),
		"resources":       flattenComputeWorkloadResourceRequirements(container.Resources),
		"volume_mount":    flattenComputeWorkloadVolumeMounts(container.VolumeMounts),
	}
}

// flattenComputeWorkloadContainer flattens a workload container as given with no
// respect to ordering. The order of the returned data is not guaranteed. If ordering
// is important, flattenComputeWorkloadContainerOrdered should be used.
func flattenComputeWorkloadContainer(name string, container workload_models.V1ContainerSpec) map[string]interface{} {
	return map[string]interface{}{
		"name":            name,
		"image":           container.Image,
		"command":         flattenStringArray(container.Command),
		"port":            flattenComputeWorkloadPorts(container.Ports),
		"env":             flattenComputeWorkloadEnvVars(container.Env),
		"readiness_probe": flattenComputeWorkloadProbe(container.ReadinessProbe),
		"liveness_probe":  flattenComputeWorkloadProbe(container.LivenessProbe),
		"resources":       flattenComputeWorkloadResourceRequirements(container.Resources),
		"volume_mount":    flattenComputeWorkloadVolumeMounts(container.VolumeMounts),
	}
}

func flattenComputeWorkloadVolumeMounts(mounts []*workload_models.V1InstanceVolumeMount) []interface{} {
	flattened := make([]interface{}, len(mounts))
	for i, mount := range mounts {
		flattened[i] = map[string]interface{}{
			"slug":       mount.Slug,
			"mount_path": mount.MountPath,
		}
	}
	return flattened
}

func flattenComputeWorkloadResourceRequirements(r *workload_models.V1ResourceRequirements) []interface{} {
	return []interface{}{map[string]interface{}{
		"requests": map[string]interface{}{
			"cpu":    r.Requests["cpu"],
			"memory": r.Requests["memory"],
		},
	}}
}

func flattenComputeWorkloadProbe(p *workload_models.V1Probe) []interface{} {
	if p == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"period_seconds":        p.PeriodSeconds,
			"timeout_seconds":       p.TimeoutSeconds,
			"success_threshold":     p.SuccessThreshold,
			"failure_threshold":     p.FailureThreshold,
			"initial_delay_seconds": p.InitialDelaySeconds,
			"http_get":              flattenComputeWorkloadHTTPGetAction(p.HTTPGet),
			"tcp_socket":            flattenComputeWorkloadTCPSocket(p.TCPSocket),
		},
	}
}

func flattenComputeWorkloadTCPSocket(tcp *workload_models.V1TCPSocketAction) []interface{} {
	if tcp == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"port": tcp.Port,
		},
	}
}

func flattenComputeWorkloadHTTPGetAction(httpGet *workload_models.V1HTTPGetAction) []interface{} {
	if httpGet == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"port":         httpGet.Port,
			"path":         httpGet.Path,
			"scheme":       httpGet.Scheme,
			"http_headers": flattenStringMap(httpGet.HTTPHeaders),
		},
	}
}

// flattenComputeWorkloadEnvVarsOrdered flattens the environment variables for a workload
// while respecting the original order of the previous environment variables. Ordering is
// important when dealing with updates to existing resources and accurate diffs are desired.
func flattenComputeWorkloadEnvVarsOrdered(prefix string, data *schema.ResourceData, envVars workload_models.V1EnvironmentVariableMapEntry) []interface{} {
	ordered := make(map[interface{}]int, data.Get(prefix+".#").(int))
	for i, n := range data.Get(prefix).([]interface{}) {
		ordered[n.(map[string]interface{})["key"].(string)] = i
	}
	e := make([]interface{}, data.Get(prefix+".#").(int))
	for key, v := range envVars {
		val := map[string]interface{}{
			"key":   key,
			"value": v.Value,
		}

		if index, exists := ordered[key]; exists {
			e[index] = val
		} else {
			e = append(e, val)
		}
	}
	return e
}

func flattenComputeWorkloadEnvVars(envVars workload_models.V1EnvironmentVariableMapEntry) []interface{} {
	e := make([]interface{}, 0, len(envVars))
	for k, v := range envVars {
		e = append(e, map[string]interface{}{
			"key":   k,
			"value": v.Value,
		})
	}
	return e
}

// flattenComputeWorkloadPortsOrdered flattens the port definitions for a workload while
// respecting the original order of the previous port definitions. Ordering is important
// when dealing with updates to existing resources and accurate diffs are desired.
func flattenComputeWorkloadPortsOrdered(prefix string, data *schema.ResourceData, ports workload_models.V1InstancePortMapEntry) []interface{} {
	ordered := make(map[interface{}]int, data.Get(prefix+".#").(int))
	for i, n := range data.Get(prefix).([]interface{}) {
		ordered[n.(map[string]interface{})["name"].(string)] = i
	}

	newPorts := make([]interface{}, data.Get(prefix+".#").(int))
	for portName, v := range ports {
		port := map[string]interface{}{
			"name":                           portName,
			"port":                           v.Port,
			"protocol":                       v.Protocol,
			"enable_implicit_network_policy": v.EnableImplicitNetworkPolicy,
		}
		if index, exists := ordered[portName]; exists {
			newPorts[index] = port
		} else {
			newPorts = append(newPorts, port)
		}
	}
	return newPorts
}

func flattenComputeWorkloadPorts(ports workload_models.V1InstancePortMapEntry) []interface{} {
	p := make([]interface{}, 0, len(ports))
	for k, v := range ports {
		p = append(p, map[string]interface{}{
			"name":     k,
			"port":     v.Port,
			"protocol": v.Protocol,
		})
	}
	return p
}
