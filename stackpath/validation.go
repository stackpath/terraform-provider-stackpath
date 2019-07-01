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
