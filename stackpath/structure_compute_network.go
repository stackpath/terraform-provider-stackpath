package stackpath

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"
)

func convertComputeNetwork(data *schema.ResourceData) *ipam_models.NetworkNetwork {
	network := ipam_models.NetworkNetwork{
		Name:       data.Get("name").(string),
		Slug:       data.Get("slug").(string),
		RootSubnet: data.Get("root_subnet").(string),
		Metadata: &ipam_models.NetworkMetadata{
			Labels:      convertWorkloadToIPAMStringMapEntry(convertToStringMap(data.Get("labels").(map[string]interface{}))),
			Annotations: convertWorkloadToIPAMStringMapEntry(convertToStringMap(data.Get("annotations").(map[string]interface{}))),
		},
	}
	if data.Get("version") != nil {
		network.Metadata.Version = data.Get("version").(string)
	}

	if rawValue := data.Get("ip_families"); rawValue != nil {
		if ipFamilies, ok := rawValue.([]interface{}); ok {
			convertedIPFamilies := make([]string, len(ipFamilies))
			for i, ipFamilyRawValue := range ipFamilies {
				if ipFamily, ok := ipFamilyRawValue.(string); ok {
					convertedIPFamilies[i] = ipFamily
				}
			}
			network.IPFamilies = convertedIPFamilies
		}
	}

	if data.Get("ipv6_subnet") != nil {
		network.IPV6Subnet = data.Get("ipv6_subnet").(string)
	}

	return &network
}
