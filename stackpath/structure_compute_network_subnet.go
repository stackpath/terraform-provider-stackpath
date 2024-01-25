package stackpath

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"
)

func formatNetworkSubnetID(networkID, subnetID string) string {
	return networkID + "/" + subnetID
}
func parseNetworkSubnetID(id string) (networkID, subnetID string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		err = fmt.Errorf("found %d parts instead of 2", len(parts))
	} else {
		networkID = parts[0]
		subnetID = parts[1]
	}
	return
}

func convertComputeNetworkSubnet(data *schema.ResourceData) *ipam_models.NetworkNetworkSubnet {
	networkSubnet := ipam_models.NetworkNetworkSubnet{
		Name:      data.Get("name").(string),
		Slug:      data.Get("slug").(string),
		NetworkID: data.Get("network_id").(string),
		Prefix:    data.Get("prefix").(string),
		Metadata: &ipam_models.NetworkMetadata{
			Labels:      convertWorkloadToIPAMStringMapEntry(convertToStringMap(data.Get("labels").(map[string]interface{}))),
			Annotations: convertWorkloadToIPAMStringMapEntry(convertToStringMap(data.Get("annotations").(map[string]interface{}))),
		},
	}
	if data.Get("version") != nil {
		networkSubnet.Metadata.Version = data.Get("version").(string)
	}

	return &networkSubnet
}
