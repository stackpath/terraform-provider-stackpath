package stackpath

import (
	"fmt"
	"net"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// validateSubnet can be used for terraform schema
// validation to ensure a value is a valid subnet
func validateSubnet(cidr interface{}, path cty.Path) diag.Diagnostics {
	if _, _, err := net.ParseCIDR(cidr.(string)); err != nil {
		return diag.FromErr(fmt.Errorf("could not parse '%v' as a valid CIDR: %v", cidr, err))
	}

	return diag.Diagnostics{}
}

// Validate the visibility property on an object bucket
// This needs to be done on the client side as it is not specified in
// the create API request
func validateObjectStorageBucketVisibility(val interface{}, path cty.Path) diag.Diagnostics {
	v := val.(string)
	if v != "PRIVATE" && v != "PUBLIC" {
		return diag.FromErr(fmt.Errorf("%q must be either PRIVATE or PUBLIC got %q", path, v))
	}
	return diag.Diagnostics{}
}
