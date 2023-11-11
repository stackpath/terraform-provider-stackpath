package stackpath

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_client/virtual_private_cloud"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
)

func TestAccComputeVPCNetworkSubnetIPv4(t *testing.T) {
	t.Parallel()

	subnet := &ipam_models.NetworkNetworkSubnet{}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeSubnetCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: `
				resource "stackpath_compute_vpc_network" "net" {
					name = "test-tf-network-3"
					slug = "test-tf-network-3"
					root_subnet = "10.0.0.0/8"
				}
				// Create new network subnet from slug and name
				resource "stackpath_compute_vpc_network_subnet" "foo" {
					name = "test-tf-subnet-1"
					slug = "test-tf-subnet-1"
					network_id = stackpath_compute_vpc_network.net.slug
					prefix = "11.0.0.0/9"
				}`,
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckSubnetExists("stackpath_compute_vpc_network_subnet.foo", subnet),
					testAccCheckSubnetMatch(subnet, &ipam_models.NetworkNetworkSubnet{
						Name:   "test-tf-subnet-1",
						Slug:   "test-tf-subnet-1",
						Prefix: "11.0.0.0/9",
						Metadata: &ipam_models.NetworkMetadata{
							Version: "1",
						},
					}),
				),
			},
		},
	})
}

func TestAccComputeVPCNetworkSubnetIPv6(t *testing.T) {
	t.Parallel()

	subnet := &ipam_models.NetworkNetworkSubnet{}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeSubnetCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: `
				resource "stackpath_compute_vpc_network" "net" {
					name = "test-tf-network-4"
					slug = "test-tf-network-4"
					root_subnet = "10.0.0.0/8"
					ip_families = ["IPv4", "IPv6"]
					ipv6_subnet = "fc11::/64"
				}
				// Create new network subnet from slug and name
				resource "stackpath_compute_vpc_network_subnet" "foo" {
					name = "test-tf-subnet-ipv6"
					slug = "test-tf-subnet-ipv6"
					network_id = stackpath_compute_vpc_network.net.slug
					prefix = "fc01::/64"
				}`,
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckSubnetExists("stackpath_compute_vpc_network_subnet.foo", subnet),
					testAccCheckSubnetMatch(subnet, &ipam_models.NetworkNetworkSubnet{
						Name:   "test-tf-subnet-ipv6",
						Slug:   "test-tf-subnet-ipv6",
						Prefix: "fc01::/64",
						Metadata: &ipam_models.NetworkMetadata{
							Version: "1",
						},
					}),
				),
			},
		},
	})
}

func testAccCheckSubnetMatch(got, want *ipam_models.NetworkNetworkSubnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if want == nil && got != nil {
			return errors.New("mismatch subnet. got=non-nil want=nil")
		}
		if want != nil && got == nil {
			return errors.New("mismatch subnet. got=nil want=non-nil")
		}
		if want.Name != got.Name {
			return fmt.Errorf("mismatch subnet.Name. got=%s want=%s", got.Name, want.Name)
		}
		if want.Slug != got.Slug {
			return fmt.Errorf("mismatch subnet.Slug. got=%s want=%s", got.Slug, want.Slug)
		}
		if want.Prefix != got.Prefix {
			return fmt.Errorf("mismatch subnet.Prefix. got=%s want=%s", got.Slug, want.Slug)
		}
		if want.Metadata == nil && got.Metadata != nil {
			return errors.New("mismatch subnet.Metadata. got=non-nil want=nil")
		}
		if want.Metadata != nil && got.Metadata == nil {
			return errors.New("mismatch subnet.Metadata. got=nil want=non-nil")
		}
		if want.Metadata != nil {
			if want.Metadata.Version != got.Metadata.Version {
				return fmt.Errorf("mismatch subnet.Metadata.Version. got=%s want=%s", got.Metadata.Version, want.Metadata.Version)
			}
			if !reflect.DeepEqual(want.Metadata.Labels, got.Metadata.Labels) {
				return fmt.Errorf("mismatch subnet.Metadata.Labels. got=%#v want=%#v", got.Metadata.Labels, want.Metadata.Labels)
			}
		}
		return nil
	}
}

func testAccComputeCheckSubnetExists(name string, subnet *ipam_models.NetworkNetworkSubnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)
		params := virtual_private_cloud.GetNetworkSubnetParams{
			StackID: config.StackID,
			Context: context.Background(),
		}
		var err error
		params.NetworkID, params.SubnetID, err = parseNetworkSubnetID(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("could not retrieve subnet: %v", err)
		}
		found, err := config.edgeComputeNetworking.VirtualPrivateCloud.GetNetworkSubnet(&params, nil)
		if err != nil {
			return fmt.Errorf("could not retrieve subnet: %v", err)
		}

		*subnet = *found.Payload.Subnet

		return nil
	}
}

func testAccComputeSubnetCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "stackpath_compute_vpc_route" {
				continue
			}
			params := virtual_private_cloud.GetRouteParams{
				StackID: config.StackID,
				Context: context.Background(),
			}
			fmt.Sscanf(rs.Primary.ID, "%s/%s", params.NetworkID, params.RouteID)

			resp, err := config.edgeComputeNetworking.VirtualPrivateCloud.GetRoute(&params, nil)
			// Since compute workloads are deleted asynchronously, we want to look at the fact that
			// the deleteRequestedAt timestamp was set on the workload. This field is used to indicate
			// that the workload is being deleted.
			if err == nil && *resp.Payload.Route.Metadata.DeleteRequestedAt == strfmt.NewDateTime() {
				return fmt.Errorf("route still exists: %v", rs.Primary.ID)
			}
		}

		return nil
	}
}
