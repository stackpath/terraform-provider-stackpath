package stackpath

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
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
	return &network
}
