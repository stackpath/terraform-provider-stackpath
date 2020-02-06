package stackpath

import (
	"fmt"
	"net"
)

// validateSubnet can be used for terraform schema
// validation to ensure a value is a valid subnet
func validateSubnet(cidr interface{}, key string) ([]string, []error) {
	if _, _, err := net.ParseCIDR(cidr.(string)); err != nil {
		return nil, []error{fmt.Errorf("could not parse '%v' as a valid CIDR: %v", cidr, err)}
	}
	return nil, nil
}

// Validate the visibility property on an object bucket
// This needs to be done on the client side as it is not specified in
// the create API request
func validateObjectStorageBucketVisibility(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)
	if v != "PRIVATE" && v != "PUBLIC" {
		errs = append(errs, fmt.Errorf("%q must be either PRIVATE or PUBLIC got %q", key, v))
	}
	return
}
