package stackpath

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/internal/models"
)

// convert from the terraform data structure to the workload data structure we need for API calls
func convertComputeWorkload(data *schema.ResourceData) *models.V1Workload {
	return &models.V1Workload{
		Name: data.Get("name").(string),
		Slug: data.Get("slug").(string),
		Metadata: &models.V1Metadata{
			Annotations: convertToStringMap(data.Get("annotations").(map[string]interface{})),
			Labels:      convertToStringMap(data.Get("labels").(map[string]interface{})),
		},
		Spec: &models.V1WorkloadSpec{
			Containers:           convertComputeWorkloadContainers("container", data),
			VirtualMachines:      convertComputeWorkloadVirtualMachines("virtual_machine", data),
			NetworkInterfaces:    convertComputeWorkloadNetworkInterfaces(data.Get("network_interface").([]interface{})),
			ImagePullCredentials: convertComputeWorkloadImagePullCredentials("image_pull_credentials", data),
			VolumeClaimTemplates: convertComputeWorkloadVolumeClaims("volume_claim", data),
		},
		Targets: convertComputeWorkloadTargets(data.Get("target").([]interface{})),
	}
}

func convertComputeWorkloadVirtualMachines(prefix string, data *schema.ResourceData) models.V1VirtualMachineSpecMapEntry {
	vms := make(models.V1VirtualMachineSpecMapEntry, data.Get(prefix+".#").(int))
	for i, vm := range data.Get(prefix).([]interface{}) {
		vmData := vm.(map[string]interface{})
		vms[vmData["name"].(string)] = models.V1VirtualMachineSpec{
			Image:          vmData["image"].(string),
			LivenessProbe:  convertComputeWorkloadProbe(fmt.Sprintf("%s.%d.liveness_probe", prefix, i), data),
			ReadinessProbe: convertComputeWorkloadProbe(fmt.Sprintf("%s.%d.readiness_probe", prefix, i), data),
			Ports:          convertComputeWorkloadPorts(fmt.Sprintf("%s.%d.port", prefix, i), data),
			Resources:      convertComputeWorkloadResourceRequirements(fmt.Sprintf("%s.%d.resources", prefix, i), data),
			VolumeMounts:   convertComputeWorkloadVolumeMounts(fmt.Sprintf("%s.%d.volume_mount", prefix, i), data),
		}
	}
	return vms
}

func convertComputeWorkloadVolumeClaims(prefix string, data *schema.ResourceData) []*models.V1VolumeClaim {
	volumes := make([]*models.V1VolumeClaim, 0, data.Get(prefix+".#").(int))
	for i, vol := range data.Get(prefix).([]interface{}) {
		volumeData := vol.(map[string]interface{})
		volumes = append(volumes, &models.V1VolumeClaim{
			Name: volumeData["name"].(string),
			Slug: volumeData["slug"].(string),
			Spec: &models.V1VolumeClaimSpec{
				Resources: convertComputeWorkloadResourceRequirements(fmt.Sprintf("%s.%d.resources", prefix, i), data),
			},
		})
	}
	return volumes
}

func convertComputeWorkloadImagePullCredentials(prefix string, data *schema.ResourceData) []*models.V1ImagePullCredential {
	credentials := make([]*models.V1ImagePullCredential, 0, data.Get(prefix+".#").(int))
	for _, c := range data.Get(prefix).([]interface{}) {
		// only docker registry credentials are allowed for pull credentials, these credentials are guaranteed to exist
		creds := c.(map[string]interface{})["docker_registry"].([]interface{})[0].(map[string]interface{})
		credentials = append(credentials, &models.V1ImagePullCredential{
			DockerRegistry: &models.V1DockerRegistryCredentials{
				Server:   creds["server"].(string),
				Username: creds["username"].(string),
				Password: creds["password"].(string),
				Email:    creds["email"].(string),
			},
		})
	}
	return credentials
}

func convertComputeWorkloadTargets(data []interface{}) models.V1TargetMapEntry {
	targets := make(models.V1TargetMapEntry, len(data))
	for _, t := range data {
		target := t.(map[string]interface{})
		targets[target["name"].(string)] = models.V1Target{
			Spec: &models.V1TargetSpec{
				DeploymentScope: target["deployment_scope"].(string),
				Deployments: &models.V1DeploymentSpec{
					MinReplicas: int32(target["min_replicas"].(int)),
					Selectors:   convertComputeMatchExpression(target["selector"].([]interface{})),
				},
			},
		}
	}
	return targets
}

func convertComputeWorkloadNetworkInterfaces(data []interface{}) []*models.V1NetworkInterface {
	interfaces := make([]*models.V1NetworkInterface, len(data))
	for i, n := range data {
		interfaces[i] = &models.V1NetworkInterface{
			Network: n.(map[string]interface{})["network"].(string),
		}
	}
	return interfaces
}

func convertComputeWorkloadContainers(prefix string, data *schema.ResourceData) models.V1ContainerSpecMapEntry {
	// Keep track of the names of containers that exist in the new workload definition
	containerNames := make(map[interface{}]bool, data.Get(prefix+".#").(int))
	// First we need to add the containers that still exist within the workload definition
	containers := make(models.V1ContainerSpecMapEntry, data.Get(prefix+".#").(int))
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
				containers[containerData["name"].(string)] = models.V1ContainerSpec{}
			}
		}
	}

	return containers
}

func convertComputeWorkloadContainer(prefix string, data *schema.ResourceData) models.V1ContainerSpec {
	return models.V1ContainerSpec{
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

func convertComputeWorkloadVolumeMounts(prefix string, data *schema.ResourceData) []*models.V1InstanceVolumeMount {
	mounts := make([]*models.V1InstanceVolumeMount, 0, data.Get(prefix+".#").(int))
	for _, m := range data.Get(prefix).([]interface{}) {
		mount := m.(map[string]interface{})
		mounts = append(mounts, &models.V1InstanceVolumeMount{
			Slug:      mount["slug"].(string),
			MountPath: mount["mount_path"].(string),
		})
	}
	return mounts
}

func convertComputeWorkloadResourceRequirements(prefix string, data *schema.ResourceData) *models.V1ResourceRequirements {
	return &models.V1ResourceRequirements{
		Requests: convertToStringMap(data.Get(prefix).([]interface{})[0].(map[string]interface{})["requests"].(map[string]interface{})),
	}
}

func convertComputeWorkloadProbe(prefix string, data *schema.ResourceData) *models.V1Probe {
	if !data.HasChange(prefix) && data.Get(prefix+".#").(int) == 0 {
		return nil
	} else if data.HasChange(prefix) && data.Get(prefix+".#").(int) == 0 {
		log.Printf("[DEBUG] removing probe from container: %v", prefix)
		// we are removing the probe so we should set the probe to an empty value
		return &models.V1Probe{}
	}

	probe := data.Get(prefix + ".0").(map[string]interface{})
	if len(probe) == 0 {
		log.Printf("[WARNING] length of probe is 0: %v", prefix)
		return nil
	}

	log.Printf("[DEBUG] adding probe for container: %v", prefix)

	return &models.V1Probe{
		FailureThreshold:    int32(probe["failure_threshold"].(int)),
		SuccessThreshold:    int32(probe["success_threshold"].(int)),
		InitialDelaySeconds: int32(probe["initial_delay_seconds"].(int)),
		PeriodSeconds:       int32(probe["period_seconds"].(int)),
		TimeoutSeconds:      int32(probe["timeout_seconds"].(int)),
		HTTPGet:             convertComputeWorkloadProbeHTTPGet(probe["http_get"].([]interface{})),
		TCPSocket:           convertComputeWorkloadProbeTCPSocket(probe["tcp_socket"].([]interface{})),
	}
}

func convertComputeWorkloadProbeTCPSocket(data []interface{}) *models.V1TCPSocketAction {
	if len(data) == 0 {
		return nil
	}

	tcpAction := data[0].(map[string]interface{})

	return &models.V1TCPSocketAction{
		Port: int32(tcpAction["port"].(int)),
	}
}

func convertComputeWorkloadProbeHTTPGet(data []interface{}) *models.V1HTTPGetAction {
	if len(data) == 0 {
		return nil
	}

	httpAction := data[0].(map[string]interface{})

	return &models.V1HTTPGetAction{
		Path:        httpAction["path"].(string),
		Port:        int32(httpAction["port"].(int)),
		Scheme:      strings.ToUpper(httpAction["scheme"].(string)),
		HTTPHeaders: convertToStringMap(httpAction["http_headers"].(map[string]interface{})),
	}
}

func convertComputeWorkloadPorts(prefix string, data *schema.ResourceData) models.V1InstancePortMapEntry {
	portNames := make(map[interface{}]bool, data.Get(prefix+".#").(int))
	ports := make(models.V1InstancePortMapEntry, data.Get(prefix+".#").(int))
	for _, s := range data.Get(prefix).([]interface{}) {
		portData := s.(map[string]interface{})
		// track the port names so we can keep track of what needs to be removed
		portNames[portData["name"]] = true
		ports[portData["name"].(string)] = models.V1InstancePort{
			Port:     int32(portData["port"].(int)),
			Protocol: strings.ToUpper(portData["protocol"].(string)),
		}
	}

	// Remove all ports that previously existed but were removed
	if data.HasChange(prefix) {
		old, _ := data.GetChange(prefix)
		for _, s := range old.([]interface{}) {
			portData := s.(map[string]interface{})
			if !portNames[portData["name"]] {
				ports[portData["name"].(string)] = models.V1InstancePort{}
			}
		}
	}

	return ports
}

func convertComputeWorkloadEnvironmentVariables(prefix string, data *schema.ResourceData) models.V1EnvironmentVariableMapEntry {
	envVarNames := make(map[interface{}]bool, data.Get(prefix+".#").(int))
	envVars := make(models.V1EnvironmentVariableMapEntry, data.Get(prefix+".#").(int))
	for _, s := range data.Get(prefix).([]interface{}) {
		envVarData := s.(map[string]interface{})
		log.Printf("[DEBUG] setting environment variable '%s'", envVarData["key"])
		envVarNames[envVarData["key"]] = true
		envVars[envVarData["key"].(string)] = models.V1EnvironmentVariable{
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
				envVars[envVarData["key"].(string)] = models.V1EnvironmentVariable{}
			}
		}
	}
	return envVars
}

func flattenComputeWorkload(data *schema.ResourceData, workload *models.V1Workload) {
	data.Set("name", workload.Name)
	data.Set("slug", workload.Slug)
	data.Set("labels", flattenStringMap(workload.Metadata.Labels))
	data.Set("annotations", flattenStringMap(workload.Metadata.Annotations))
	data.Set("network_interface", flattenComputeWorkloadNetworkInterfaces(workload.Spec.NetworkInterfaces))
	data.Set("container", flattenComputeWorkloadContainers("container", data, workload.Spec.Containers))
	data.Set("image_pull_credentials", flattenComputeWorkloadImagePullCredentials(workload.Spec.ImagePullCredentials))
	data.Set("virtual_machine", flattenComputeWorkloadVirtualMachines("virtual_machine", data, workload.Spec.VirtualMachines))
	data.Set("volume_claim", flattenComputeWorkloadVolumeClaims(workload.Spec.VolumeClaimTemplates))
	data.Set("target", flattenComputeWorkloadTargets("target", data, workload.Targets))
}

func flattenComputeWorkloadVolumeClaims(claims []*models.V1VolumeClaim) []interface{} {
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

func flattenComputeWorkloadVirtualMachines(prefix string, data *schema.ResourceData, vms models.V1VirtualMachineSpecMapEntry) []interface{} {
	// Ensure we keep the original order so terraform doesn't mistaken things as out of sync
	ordered := make(map[string]int, data.Get(prefix+".#").(int))
	for i, k := range data.Get(prefix).([]interface{}) {
		// Set the name of the container in the map with its expected
		// index, container names are guaranteed to be unique
		ordered[k.(map[string]interface{})["name"].(string)] = i
	}
	flattened := make([]interface{}, data.Get(prefix+".#").(int))
	for name, vm := range vms {
		flattenedVM := map[string]interface{}{
			"name":            name,
			"image":           vm.Image,
			"port":            flattenComputeWorkloadPorts(vm.Ports),
			"readiness_probe": flattenComputeWorkloadProbe(vm.ReadinessProbe),
			"liveness_probe":  flattenComputeWorkloadProbe(vm.LivenessProbe),
			"resources":       flattenComputeWorkloadResourceRequirements(vm.Resources),
			"volume_mount":    flattenComputeWorkloadVolumeMounts(vm.VolumeMounts),
		}

		if index, found := ordered[name]; found {
			flattened[index] = flattenedVM
		} else {
			flattened = append(flattened, flattenedVM)
		}
	}
	return flattened
}

func flattenComputeWorkloadImagePullCredentials(credentials []*models.V1ImagePullCredential) []interface{} {
	creds := make([]interface{}, 0, len(credentials))
	for _, c := range credentials {
		creds = append(creds, map[string]interface{}{
			"docker_registry": map[string]interface{}{
				"server":   c.DockerRegistry.Server,
				"username": c.DockerRegistry.Username,
				"email":    c.DockerRegistry.Email,
			},
		})
	}
	return creds
}

func flattenComputeWorkloadTargets(prefix string, data *schema.ResourceData, targets models.V1TargetMapEntry) []interface{} {
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
			t[index] = flattenComputeWorkloadTarget(k, v)
		} else {
			t = append(t, flattenComputeWorkloadTarget(k, v))
		}
	}
	return t
}

func flattenComputeWorkloadTarget(name string, target models.V1Target) map[string]interface{} {
	return map[string]interface{}{
		"name":             name,
		"min_replicas":     target.Spec.Deployments.MinReplicas,
		"deployment_scope": target.Spec.DeploymentScope,
		"selector":         flattenComputeMatchExpressions(target.Spec.Deployments.Selectors),
	}
}

func flattenComputeWorkloadNetworkInterfaces(networkInterfaces []*models.V1NetworkInterface) []interface{} {
	flattened := make([]interface{}, len(networkInterfaces))
	for i, n := range networkInterfaces {
		flattened[i] = map[string]interface{}{
			"network": n.Network,
		}
	}
	return flattened
}

func flattenComputeWorkloadContainers(prefix string, data *schema.ResourceData, containers models.V1ContainerSpecMapEntry) []interface{} {
	// Ensure we keep the original order so terraform doesn't mistaken things as out of sync
	ordered := make(map[string]int, data.Get(prefix+".#").(int))
	for i, k := range data.Get(prefix).([]interface{}) {
		// Set the name of the container in the map with its expected
		// index, container names are guaranteed to be unique
		ordered[k.(map[string]interface{})["name"].(string)] = i
	}
	flattened := make([]interface{}, data.Get(prefix+".#").(int))
	for name, container := range containers {
		flattenedContainer := map[string]interface{}{
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

		// In the event that a container is added to a workload outside of
		// terraform, we need to support adding unknown containers to our state.
		if index, found := ordered[name]; found {
			flattened[index] = flattenedContainer
		} else {
			flattened = append(flattened, flattenedContainer)
		}
	}
	return flattened
}

func flattenComputeWorkloadVolumeMounts(mounts []*models.V1InstanceVolumeMount) []interface{} {
	flattened := make([]interface{}, len(mounts))
	for i, mount := range mounts {
		flattened[i] = map[string]interface{}{
			"slug":       mount.Slug,
			"mount_path": mount.MountPath,
		}
	}
	return flattened
}

func flattenComputeWorkloadResourceRequirements(r *models.V1ResourceRequirements) []interface{} {
	return []interface{}{map[string]interface{}{
		"requests": map[string]interface{}{
			"cpu":    r.Requests["cpu"],
			"memory": r.Requests["memory"],
		},
	}}
}

func flattenComputeWorkloadProbe(p *models.V1Probe) []interface{} {
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

func flattenComputeWorkloadTCPSocket(tcp *models.V1TCPSocketAction) []interface{} {
	if tcp == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"port": tcp.Port,
		},
	}
}

func flattenComputeWorkloadHTTPGetAction(httpGet *models.V1HTTPGetAction) []interface{} {
	if httpGet == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"port":    httpGet.Port,
			"path":    httpGet.Path,
			"scheme":  httpGet.Scheme,
			"headers": flattenStringMap(httpGet.HTTPHeaders),
		},
	}
}

func flattenComputeWorkloadEnvVars(envVars models.V1EnvironmentVariableMapEntry) []interface{} {
	e := make([]interface{}, 0, len(envVars))
	for k, v := range envVars {
		e = append(e, map[string]interface{}{
			"key":   k,
			"value": v.Value,
		})
	}
	return e
}

func flattenComputeWorkloadPorts(ports models.V1InstancePortMapEntry) []interface{} {
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

func flattenComputeNetworkPolicyIngress(data []*models.V1Ingress) []interface{} {
	d := make([]interface{}, len(data))
	for i, ingress := range data {
		d[i] = map[string]interface{}{
			"action":      string(ingress.Action),
			"description": ingress.Description,
			"protocol":    flattenComputeNetworkPolicyProtocols(ingress.Protocols),
			"from":        flattenComputeNetworkPolicyHostRule(ingress.From),
		}
	}
	return d
}

func flattenComputeNetworkPolicyEgress(data []*models.V1Egress) []interface{} {
	d := make([]interface{}, len(data))
	for i, egress := range data {
		d[i] = map[string]interface{}{
			"action":      string(egress.Action),
			"description": egress.Description,
			"protocol":    flattenComputeNetworkPolicyProtocols(egress.Protocols),
			"to":          flattenComputeNetworkPolicyHostRule(egress.To),
		}
	}
	return d
}

func isProtocolEmpty(p *models.V1Protocols) bool {
	if p == nil {
		return true
	}
	// consider it empty when all the fields are nil
	return p.Ah == nil && p.Esp == nil && p.Gre == nil && p.Icmp == nil && p.TCP == nil && p.TCPUDP == nil && p.UDP == nil
}

func flattenComputeNetworkPolicyProtocols(data *models.V1Protocols) []interface{} {
	// The API will return an object regardless if one is passed in or not. This causes
	// terraform to believe the protocol object is defined in the API response but not
	// in terraform itself leading to an inconsistency in the state. To work around this
	// we tell terraform the object doesn't exist when the object is empty
	if isProtocolEmpty(data) {
		return nil
	}

	protocol := map[string]interface{}{}
	if data.Ah != nil {
		// no configuration options to provide
		protocol["ah"] = []interface{}{}
	}
	if data.Esp != nil {
		// no configuration options to provide
		protocol["esp"] = []interface{}{}
	}
	if data.Gre != nil {
		// no configuration options to provide
		protocol["gre"] = []interface{}{}
	}
	if data.Icmp != nil {
		// no configuration options to provide
		protocol["icmp"] = []interface{}{}
	}
	if data.TCP != nil {
		protocol["tcp"] = []interface{}{
			map[string]interface{}{
				"destination_ports": flattenStringArray(data.TCP.DestinationPorts),
				"source_ports":      flattenStringArray(data.TCP.SourcePorts),
			},
		}
	}
	if data.TCPUDP != nil {
		protocol["tcp_udp"] = []interface{}{
			map[string]interface{}{
				"destination_ports": flattenStringArray(data.TCPUDP.DestinationPorts),
				"source_ports":      flattenStringArray(data.TCPUDP.SourcePorts),
			},
		}
	}
	if data.UDP != nil {
		protocol["udp"] = []interface{}{
			map[string]interface{}{
				"destination_ports": flattenStringArray(data.UDP.DestinationPorts),
				"source_ports":      flattenStringArray(data.UDP.SourcePorts),
			},
		}
	}
	return []interface{}{protocol}
}

func flattenComputeNetworkPolicyHostRule(data *models.V1HostRule) []interface{} {
	if data == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"instance_selector": flattenComputeMatchExpressions(data.InstanceSelectors),
			"network_selector":  flattenComputeMatchExpressions(data.NetworkSelectors),
			"ip_block":          flattenComputeNetworkPolicyIPBlock(data.IPBlock),
		},
	}
}

func flattenComputeNetworkPolicyIPBlock(data []*models.V1IPBlock) []interface{} {
	flattened := make([]interface{}, len(data))
	for i, block := range data {
		flattened[i] = map[string]interface{}{
			"cidr":   block.Cidr,
			"except": flattenStringArray(block.Except),
		}
	}
	return flattened
}
