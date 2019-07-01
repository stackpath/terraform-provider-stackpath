package stackpath

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/internal/models"
)

func convertComputeNetworkPolicy(data *schema.ResourceData) *models.V1NetworkPolicy {
	return &models.V1NetworkPolicy{
		Name:        data.Get("name").(string),
		Slug:        data.Get("slug").(string),
		Description: data.Get("description").(string),
		Metadata: &models.V1Metadata{
			Labels:      convertToStringMap(data.Get("labels").(map[string]interface{})),
			Annotations: convertToStringMap(data.Get("annotations").(map[string]interface{})),
		},
		Spec: &models.V1NetworkPolicySpec{
			Priority:          int32(data.Get("priority").(int)),
			PolicyTypes:       convertComputeNetworkPolicyTypes(data.Get("policy_types").([]interface{})),
			InstanceSelectors: convertComputeMatchExpression(data.Get("instance_selector").([]interface{})),
			NetworkSelectors:  convertComputeMatchExpression(data.Get("network_selector").([]interface{})),
			Ingress:           convertComputeNetworkPolicyIngress("ingress", data),
			Egress:            convertComputeNetworkPolicyEgress("egress", data),
		},
	}
}

func convertComputeNetworkPolicyTypes(t []interface{}) []models.NetworkPolicySpecPolicyType {
	types := make([]models.NetworkPolicySpecPolicyType, len(t))
	for i, v := range t {
		types[i] = models.NetworkPolicySpecPolicyType(v.(string))
	}
	return types
}

func convertComputeNetworkPolicyIngress(prefix string, data *schema.ResourceData) []*models.V1Ingress {
	ingress := make([]*models.V1Ingress, data.Get(prefix+".#").(int))
	for i, v := range data.Get(prefix).([]interface{}) {
		ingressData := v.(map[string]interface{})
		ingress[i] = &models.V1Ingress{
			Action:      models.V1Action(ingressData["action"].(string)),
			Description: ingressData["description"].(string),
			Protocols:   convertComputeNetworkPolicyProtocols(ingressData["protocol"].([]interface{})),
			From:        convertComputeNetworkPolicyHostRule(ingressData["from"].([]interface{})),
		}
	}
	return ingress
}

func convertComputeNetworkPolicyEgress(prefix string, data *schema.ResourceData) []*models.V1Egress {
	egress := make([]*models.V1Egress, data.Get(prefix+".#").(int))
	for i, v := range data.Get(prefix).([]interface{}) {
		egressData := v.(map[string]interface{})
		egress[i] = &models.V1Egress{
			Action:      models.V1Action(egressData["action"].(string)),
			Description: egressData["description"].(string),
			Protocols:   convertComputeNetworkPolicyProtocols(egressData["protocol"].([]interface{})),
			To:          convertComputeNetworkPolicyHostRule(egressData["to"].([]interface{})),
		}
	}
	return egress
}

func convertComputeNetworkPolicyHostRule(data []interface{}) *models.V1HostRule {
	if len(data) == 0 {
		return nil
	}

	// will only ever have one rule defined
	rule := data[0].(map[string]interface{})

	return &models.V1HostRule{
		InstanceSelectors: convertComputeMatchExpression(rule["instance_selector"].([]interface{})),
		NetworkSelectors:  convertComputeMatchExpression(rule["network_selector"].([]interface{})),
		IPBlock:           convertComputeNetworkPolicyIPBlock(rule["ip_block"].([]interface{})),
	}
}

func convertComputeNetworkPolicyIPBlock(data []interface{}) []*models.V1IPBlock {
	blocks := make([]*models.V1IPBlock, len(data))
	for i, v := range data {
		ipBlock := v.(map[string]interface{})
		blocks[i] = &models.V1IPBlock{
			Cidr:   ipBlock["cidr"].(string),
			Except: convertToStringArray(ipBlock["except"].([]interface{})),
		}
	}
	return blocks
}

func convertComputeNetworkPolicyProtocols(data []interface{}) *models.V1Protocols {
	if len(data) == 0 {
		return &models.V1Protocols{}
	}

	// will only ever have one protocol
	protocol := data[0].(map[string]interface{})

	return &models.V1Protocols{
		Ah:     convertComputeNetworkPolicyProtocolAH(protocol["ah"].([]interface{})),
		Esp:    convertComputeNetworkPolicyProtocolESP(protocol["esp"].([]interface{})),
		Gre:    convertComputeNetworkPolicyProtocolGRE(protocol["gre"].([]interface{})),
		Icmp:   convertComputeNetworkPolicyProtocolICMP(protocol["icmp"].([]interface{})),
		TCP:    convertComputeNetworkPolicyProtocolTCP(protocol["tcp"].([]interface{})),
		TCPUDP: convertComputeNetworkPolicyProtocolTCPUDP(protocol["tcp_udp"].([]interface{})),
		UDP:    convertComputeNetworkPolicyProtocolUDP(protocol["udp"].([]interface{})),
	}
}

func convertComputeNetworkPolicyProtocolAH(data []interface{}) *models.V1ProtocolAh {
	if len(data) == 0 {
		return nil
	}
	// no configuration options available
	return &models.V1ProtocolAh{}
}

func convertComputeNetworkPolicyProtocolESP(data []interface{}) *models.V1ProtocolEsp {
	if len(data) == 0 {
		return nil
	}
	// no configuration options available
	return &models.V1ProtocolEsp{}
}

func convertComputeNetworkPolicyProtocolGRE(data []interface{}) *models.V1ProtocolGre {
	if len(data) == 0 {
		return nil
	}
	// no configuration options available
	return &models.V1ProtocolGre{}
}

func convertComputeNetworkPolicyProtocolICMP(data []interface{}) *models.V1ProtocolIcmp {
	if len(data) == 0 {
		return nil
	}
	// no configuration options available
	return &models.V1ProtocolIcmp{}
}

func convertComputeNetworkPolicyProtocolUDP(data []interface{}) *models.V1ProtocolUDP {
	if len(data) == 0 {
		return nil
	}
	// will only ever have one UDP protocol defined
	protocol := data[0].(map[string]interface{})
	return &models.V1ProtocolUDP{
		DestinationPorts: convertToStringArray(protocol["destination_ports"].([]interface{})),
		SourcePorts:      convertToStringArray(protocol["source_ports"].([]interface{})),
	}
}

func convertComputeNetworkPolicyProtocolTCP(data []interface{}) *models.V1ProtocolTCP {
	if len(data) == 0 {
		return nil
	}
	// will only ever have one UDP protocol defined
	protocol := data[0].(map[string]interface{})
	return &models.V1ProtocolTCP{
		DestinationPorts: convertToStringArray(protocol["destination_ports"].([]interface{})),
		SourcePorts:      convertToStringArray(protocol["source_ports"].([]interface{})),
	}
}

func convertComputeNetworkPolicyProtocolTCPUDP(data []interface{}) *models.V1ProtocolTCPUDP {
	if len(data) == 0 {
		return nil
	}
	// will only ever have one UDP protocol defined
	protocol := data[0].(map[string]interface{})
	return &models.V1ProtocolTCPUDP{
		DestinationPorts: convertToStringArray(protocol["destination_ports"].([]interface{})),
		SourcePorts:      convertToStringArray(protocol["source_ports"].([]interface{})),
	}
}

func flattenComputeNetworkPolicyTypes(types []models.NetworkPolicySpecPolicyType) []interface{} {
	t := make([]interface{}, len(types))
	for i, v := range types {
		t[i] = string(v)
	}
	return t
}
