package stackpath

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
)

func convertComputeNetworkPolicy(data *schema.ResourceData) *ipam_models.V1NetworkPolicy {
	return &ipam_models.V1NetworkPolicy{
		Name:        data.Get("name").(string),
		Slug:        data.Get("slug").(string),
		Description: data.Get("description").(string),
		Metadata: &ipam_models.NetworkMetadata{
			Labels:      convertWorkloadToIPAMStringMapEntry(convertToStringMap(data.Get("labels").(map[string]interface{}))),
			Annotations: convertWorkloadToIPAMStringMapEntry(convertToStringMap(data.Get("annotations").(map[string]interface{}))),
			Version:     data.Get("version").(string),
		},
		Spec: &ipam_models.V1NetworkPolicySpec{
			Priority:          int32(data.Get("priority").(int)),
			PolicyTypes:       convertComputeNetworkPolicyTypes(data.Get("policy_types").([]interface{})),
			InstanceSelectors: convertWorkloadToIPAMMatchExpression(convertComputeMatchExpression(data.Get("instance_selector").([]interface{}))),
			NetworkSelectors:  convertWorkloadToIPAMMatchExpression(convertComputeMatchExpression(data.Get("network_selector").([]interface{}))),
			Ingress:           convertComputeNetworkPolicyIngress("ingress", data),
			Egress:            convertComputeNetworkPolicyEgress("egress", data),
		},
	}
}

func convertComputeNetworkPolicyTypes(t []interface{}) []*ipam_models.NetworkPolicySpecPolicyType {
	types := make([]*ipam_models.NetworkPolicySpecPolicyType, len(t))
	for i, v := range t {
		newType := ipam_models.NetworkPolicySpecPolicyType(v.(string))
		types[i] = &newType
	}
	return types
}

func convertComputeNetworkPolicyIngress(prefix string, data *schema.ResourceData) []*ipam_models.V1Ingress {
	ingress := make([]*ipam_models.V1Ingress, data.Get(prefix+".#").(int))
	for i, v := range data.Get(prefix).([]interface{}) {
		ingressData := v.(map[string]interface{})
		action := ipam_models.V1Action(ingressData["action"].(string))
		ingress[i] = &ipam_models.V1Ingress{
			Action:      &action,
			Description: ingressData["description"].(string),
			Protocols:   convertComputeNetworkPolicyProtocols(ingressData["protocol"].([]interface{})),
			From:        convertComputeNetworkPolicyHostRule(ingressData["from"].([]interface{})),
		}
	}
	return ingress
}

func convertComputeNetworkPolicyEgress(prefix string, data *schema.ResourceData) []*ipam_models.V1Egress {
	egress := make([]*ipam_models.V1Egress, data.Get(prefix+".#").(int))
	for i, v := range data.Get(prefix).([]interface{}) {
		egressData := v.(map[string]interface{})
		action := ipam_models.V1Action(egressData["action"].(string))
		egress[i] = &ipam_models.V1Egress{
			Action:      &action,
			Description: egressData["description"].(string),
			Protocols:   convertComputeNetworkPolicyProtocols(egressData["protocol"].([]interface{})),
			To:          convertComputeNetworkPolicyHostRule(egressData["to"].([]interface{})),
		}
	}
	return egress
}

func convertComputeNetworkPolicyHostRule(data []interface{}) *ipam_models.V1HostRule {
	if len(data) == 0 {
		return nil
	}

	// will only ever have one rule defined
	rule := data[0].(map[string]interface{})

	return &ipam_models.V1HostRule{
		InstanceSelectors: convertWorkloadToIPAMMatchExpression(convertComputeMatchExpression(rule["instance_selector"].([]interface{}))),
		NetworkSelectors:  convertWorkloadToIPAMMatchExpression(convertComputeMatchExpression(rule["network_selector"].([]interface{}))),
		IPBlock:           convertComputeNetworkPolicyIPBlock(rule["ip_block"].([]interface{})),
	}
}

func convertComputeNetworkPolicyIPBlock(data []interface{}) []*ipam_models.V1IPBlock {
	blocks := make([]*ipam_models.V1IPBlock, len(data))
	for i, v := range data {
		ipBlock := v.(map[string]interface{})
		blocks[i] = &ipam_models.V1IPBlock{
			Cidr:   ipBlock["cidr"].(string),
			Except: convertToStringArray(ipBlock["except"].([]interface{})),
		}
	}
	return blocks
}

func convertComputeNetworkPolicyProtocols(data []interface{}) *ipam_models.V1Protocols {
	if len(data) == 0 {
		return &ipam_models.V1Protocols{}
	}

	// will only ever have one protocol
	protocol := data[0].(map[string]interface{})

	return &ipam_models.V1Protocols{
		Ah:     convertComputeNetworkPolicyProtocolAH(protocol["ah"].([]interface{})),
		Esp:    convertComputeNetworkPolicyProtocolESP(protocol["esp"].([]interface{})),
		Gre:    convertComputeNetworkPolicyProtocolGRE(protocol["gre"].([]interface{})),
		Icmp:   convertComputeNetworkPolicyProtocolICMP(protocol["icmp"].([]interface{})),
		TCP:    convertComputeNetworkPolicyProtocolTCP(protocol["tcp"].([]interface{})),
		TCPUDP: convertComputeNetworkPolicyProtocolTCPUDP(protocol["tcp_udp"].([]interface{})),
		UDP:    convertComputeNetworkPolicyProtocolUDP(protocol["udp"].([]interface{})),
	}
}

func convertComputeNetworkPolicyProtocolAH(data []interface{}) *ipam_models.V1ProtocolAh {
	if len(data) == 0 {
		return nil
	}
	// no configuration options available
	return &ipam_models.V1ProtocolAh{}
}

func convertComputeNetworkPolicyProtocolESP(data []interface{}) *ipam_models.V1ProtocolEsp {
	if len(data) == 0 {
		return nil
	}
	// no configuration options available
	return &ipam_models.V1ProtocolEsp{}
}

func convertComputeNetworkPolicyProtocolGRE(data []interface{}) *ipam_models.V1ProtocolGre {
	if len(data) == 0 {
		return nil
	}
	// no configuration options available
	return &ipam_models.V1ProtocolGre{}
}

func convertComputeNetworkPolicyProtocolICMP(data []interface{}) *ipam_models.V1ProtocolIcmp {
	if len(data) == 0 {
		return nil
	}
	// no configuration options available
	return &ipam_models.V1ProtocolIcmp{}
}

func convertComputeNetworkPolicyProtocolUDP(data []interface{}) *ipam_models.V1ProtocolUDP {
	if len(data) == 0 {
		return nil
	}
	// will only ever have one UDP protocol defined
	protocol := data[0].(map[string]interface{})
	return &ipam_models.V1ProtocolUDP{
		DestinationPorts: convertToStringArray(protocol["destination_ports"].([]interface{})),
		SourcePorts:      convertToStringArray(protocol["source_ports"].([]interface{})),
	}
}

func convertComputeNetworkPolicyProtocolTCP(data []interface{}) *ipam_models.V1ProtocolTCP {
	if len(data) == 0 {
		return nil
	}
	// will only ever have one UDP protocol defined
	protocol := data[0].(map[string]interface{})
	return &ipam_models.V1ProtocolTCP{
		DestinationPorts: convertToStringArray(protocol["destination_ports"].([]interface{})),
		SourcePorts:      convertToStringArray(protocol["source_ports"].([]interface{})),
	}
}

func convertComputeNetworkPolicyProtocolTCPUDP(data []interface{}) *ipam_models.V1ProtocolTCPUDP {
	if len(data) == 0 {
		return nil
	}
	// will only ever have one UDP protocol defined
	protocol := data[0].(map[string]interface{})
	return &ipam_models.V1ProtocolTCPUDP{
		DestinationPorts: convertToStringArray(protocol["destination_ports"].([]interface{})),
		SourcePorts:      convertToStringArray(protocol["source_ports"].([]interface{})),
	}
}

func flattenComputeNetworkPolicyIngress(data []*ipam_models.V1Ingress) []interface{} {
	d := make([]interface{}, len(data))
	for i, ingress := range data {
		var action string
		if ingress.Action != nil {
			action = string(*ingress.Action)
		}
		d[i] = map[string]interface{}{
			"action":      action,
			"description": ingress.Description,
			"protocol":    flattenComputeNetworkPolicyProtocols(ingress.Protocols),
			"from":        flattenComputeNetworkPolicyHostRule(ingress.From),
		}
	}
	return d
}

func flattenComputeNetworkPolicyEgress(data []*ipam_models.V1Egress) []interface{} {
	d := make([]interface{}, len(data))
	for i, egress := range data {
		var action string
		if egress.Action != nil {
			action = string(*egress.Action)
		}
		d[i] = map[string]interface{}{
			"action":      action,
			"description": egress.Description,
			"protocol":    flattenComputeNetworkPolicyProtocols(egress.Protocols),
			"to":          flattenComputeNetworkPolicyHostRule(egress.To),
		}
	}
	return d
}

func isProtocolEmpty(p *ipam_models.V1Protocols) bool {
	if p == nil {
		return true
	}
	// consider it empty when all the fields are nil
	return p.Ah == nil && p.Esp == nil && p.Gre == nil && p.Icmp == nil && p.TCP == nil && p.TCPUDP == nil && p.UDP == nil
}

func flattenComputeNetworkPolicyProtocols(data *ipam_models.V1Protocols) []interface{} {
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

func flattenComputeNetworkPolicyHostRule(data *ipam_models.V1HostRule) []interface{} {
	if data == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"instance_selector": flattenComputeMatchExpressions(convertIPAMToWorkloadMatchExpression(data.InstanceSelectors)),
			"network_selector":  flattenComputeMatchExpressions(convertIPAMToWorkloadMatchExpression(data.NetworkSelectors)),
			"ip_block":          flattenComputeNetworkPolicyIPBlock(data.IPBlock),
		},
	}
}

func flattenComputeNetworkPolicyIPBlock(data []*ipam_models.V1IPBlock) []interface{} {
	flattened := make([]interface{}, len(data))
	for i, block := range data {
		flattened[i] = map[string]interface{}{
			"cidr":   block.Cidr,
			"except": flattenStringArray(block.Except),
		}
	}
	return flattened
}

func flattenComputeNetworkPolicyTypes(types []*ipam_models.NetworkPolicySpecPolicyType) []interface{} {
	length := 0
	for _, v := range types {
		if v != nil {
			length++
		}
	}

	i := 0
	t := make([]interface{}, length)
	for _, v := range types {
		if v == nil {
			continue
		}

		t[i] = string(*v)
		i++
	}
	return t
}
