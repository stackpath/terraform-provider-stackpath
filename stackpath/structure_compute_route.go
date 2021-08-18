package stackpath

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
)

func formatRouteID(networkID, routeID string) string {
	return networkID + "/" + routeID
}
func parseRouteID(id string) (networkID, routeID string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		err = fmt.Errorf("found %d parts instead of 2", len(parts))
	} else {
		networkID = parts[0]
		routeID = parts[1]
	}
	return
}

func convertComputeRoute(data *schema.ResourceData) *ipam_models.NetworkRoute {
	network := ipam_models.NetworkRoute{
		Name:                data.Get("name").(string),
		Slug:                data.Get("slug").(string),
		NetworkID:           data.Get("network_id").(string),
		DestinationPrefixes: convertToStringArray(data.Get("destination_prefixes").([]interface{})),
		GatewaySelectors:    convertResourceDataToGatewaySelectors(data.Get("gateway_selectors").([]interface{})),
		Metadata: &ipam_models.NetworkMetadata{
			Labels:      convertWorkloadToIPAMStringMapEntry(convertToStringMap(data.Get("labels").(map[string]interface{}))),
			Annotations: convertWorkloadToIPAMStringMapEntry(convertToStringMap(data.Get("annotations").(map[string]interface{}))),
		},
	}
	if data.Get("version") != nil {
		network.Metadata.Version = data.Get("version").(string)
	}
	return &network
}

func flattenGatewaySelectors(selectors []*ipam_models.RouteGatewaySelector) []interface{} {
	output := make([]interface{}, 0, len(selectors))
	for _, selector := range selectors {
		newSelector := flattenGatewaySelector(selector)
		if newSelector != nil {
			output = append(output, newSelector)
		}
	}
	return output
}

func flattenGatewaySelector(selector *ipam_models.RouteGatewaySelector) *map[string]interface{} {
	return &map[string]interface{}{
		"interface_selectors": flattenInterfaceSelectors(selector.InterfaceSelectors),
	}
}

func flattenInterfaceSelectors(selectors []*ipam_models.NetworkMatchExpression) []interface{} {
	output := make([]interface{}, 0, len(selectors))
	for _, selector := range selectors {
		newSelector := flattenInterfaceSelector(selector)
		if newSelector != nil {
			output = append(output, newSelector)
		}
	}
	return output
}

func flattenInterfaceSelector(selector *ipam_models.NetworkMatchExpression) *map[string]interface{} {
	return &map[string]interface{}{
		"key":      selector.Key,
		"operator": selector.Operator,
		"values":   flattenStringArray(selector.Values),
	}
}

func convertResourceDataToGatewaySelectors(selectors []interface{}) []*ipam_models.RouteGatewaySelector {
	output := make([]*ipam_models.RouteGatewaySelector, 0, len(selectors))
	for _, selector := range selectors {
		newSelector := convertResourceDataToGatewaySelector(selector.(map[string]interface{}))
		if newSelector != nil {
			output = append(output, newSelector)
		}
	}
	return output
}

func convertResourceDataToGatewaySelector(selector map[string]interface{}) *ipam_models.RouteGatewaySelector {
	return &ipam_models.RouteGatewaySelector{
		InterfaceSelectors: convertResourceDataToInterfaceSelectors(selector["interface_selectors"].([]interface{})),
	}
}

func convertResourceDataToInterfaceSelectors(selectors []interface{}) []*ipam_models.NetworkMatchExpression {
	output := make([]*ipam_models.NetworkMatchExpression, 0, len(selectors))
	for _, selector := range selectors {
		newSelector := convertResourceDataToInterfaceSelector(selector.(map[string]interface{}))
		if newSelector != nil {
			output = append(output, newSelector)
		}
	}
	return output
}

func convertResourceDataToInterfaceSelector(selector map[string]interface{}) *ipam_models.NetworkMatchExpression {
	return &ipam_models.NetworkMatchExpression{
		Key:      selector["key"].(string),
		Operator: selector["operator"].(string),
		Values:   convertToStringArray(selector["values"].([]interface{})),
	}
}
