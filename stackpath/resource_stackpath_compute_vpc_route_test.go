package stackpath

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_client/virtual_private_cloud"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
)

func TestAccComputeVPCRoute(t *testing.T) {
	t.Parallel()

	route := &ipam_models.NetworkRoute{}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeRouteCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: `
				resource "stackpath_compute_vpc_network" "net" {
					name = "Test network"
					slug = "test-tf-network-1"
					root_subnet = "10.0.0.0/8"
				}
				// Create new route from slug and name
				resource "stackpath_compute_vpc_route" "foo" {
					name = "test-tf-route-1"
					slug = "test-tf-route-1"
					network_id = stackpath_compute_vpc_network.net.slug
					destination_prefixes = ["11.0.0.0/8"]
					gateway_selectors {
						interface_selectors {
							key = "workload.platform.stackpath.net/workload-slug"
							operator = "in"
							values = ["test"]
						}
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckRouteExists("stackpath_compute_vpc_route.foo", route),
					testAccCheckRouteMatch(route, &ipam_models.NetworkRoute{
						Name:                "test-tf-route-1",
						Slug:                "test-tf-route-1",
						DestinationPrefixes: []string{"11.0.0.0/8"},
						GatewaySelectors: []*ipam_models.RouteGatewaySelector{
							{
								InterfaceSelectors: []*ipam_models.NetworkMatchExpression{
									{
										Key:      "workload.platform.stackpath.net/workload-slug",
										Operator: "in",
										Values:   []string{"test"},
									},
								},
							},
						},
						Metadata: &ipam_models.NetworkMetadata{
							Version: "1",
							Annotations: map[string]string{
								"ipam.platform.stackpath.net/network-slug": "test-tf-network-1",
							},
						},
					}),
				),
			},
			{
				Config: `
				resource "stackpath_compute_vpc_network" "net" {
					name = "Test network"
					slug = "test-tf-network-1"
					root_subnet = "10.0.0.0/8"
				}
				resource "stackpath_compute_vpc_route" "foo" {
					name = "test-tf-route-1"
					slug = "test-tf-route-1"
					network_id = stackpath_compute_vpc_network.net.slug
					destination_prefixes = ["13.0.0.0/8"] // Update prefix
					gateway_selectors {
						interface_selectors {
							key = "workload.platform.stackpath.net/workload-slug"
							operator = "in"
							values = ["test"]
						}
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckRouteExists("stackpath_compute_vpc_route.foo", route),
					testAccCheckRouteMatch(route, &ipam_models.NetworkRoute{
						Name:                "test-tf-route-1",
						Slug:                "test-tf-route-1",
						DestinationPrefixes: []string{"13.0.0.0/8"},
						GatewaySelectors: []*ipam_models.RouteGatewaySelector{
							{
								InterfaceSelectors: []*ipam_models.NetworkMatchExpression{
									{
										Key:      "workload.platform.stackpath.net/workload-slug",
										Operator: "in",
										Values:   []string{"test"},
									},
								},
							},
						},
						Metadata: &ipam_models.NetworkMetadata{
							Version: "2",
							Annotations: map[string]string{
								"ipam.platform.stackpath.net/network-slug": "test-tf-network-1",
							},
						},
					}),
				),
			},
		},
	})
}

func testAccCheckRouteMatch(got, want *ipam_models.NetworkRoute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if want == nil && got != nil {
			return errors.New("mismatch route. got=non-nil want=nil")
		}
		if want != nil && got == nil {
			return errors.New("mismatch route. got=nil want=non-nil")
		}
		if want.Name != got.Name {
			return fmt.Errorf("mismatch route.Name. got=%s want=%s", got.Name, want.Name)
		}
		if want.Slug != got.Slug {
			return fmt.Errorf("mismatch route.Slug. got=%s want=%s", got.Slug, want.Slug)
		}
		if !reflect.DeepEqual(want.DestinationPrefixes, got.DestinationPrefixes) {
			return fmt.Errorf("mismatch route.DestinationPrefixes. got=%#v want=%#v", got.DestinationPrefixes, want.DestinationPrefixes)
		}
		if want.Metadata == nil && got.Metadata != nil {
			return errors.New("mismatch route.Metadata. got=non-nil want=nil")
		}
		if want.Metadata != nil && got.Metadata == nil {
			return errors.New("mismatch route.Metadata. got=nil want=non-nil")
		}
		if want.Metadata != nil {
			if want.Metadata.Version != got.Metadata.Version {
				return fmt.Errorf("mismatch route.Metadata.Version. got=%s want=%s", got.Metadata.Version, want.Metadata.Version)
			}
			if !reflect.DeepEqual(want.Metadata.Labels, got.Metadata.Labels) {
				return fmt.Errorf("mismatch route.Metadata.Labels. got=%#v want=%#v", got.Metadata.Labels, want.Metadata.Labels)
			}
			if !reflect.DeepEqual(want.Metadata.Annotations, got.Metadata.Annotations) {
				return fmt.Errorf("mismatch route.Metadata.Annotations. got=%#v want=%#v", got.Metadata.Annotations, want.Metadata.Annotations)
			}
		}
		return nil
	}
}

func testAccComputeCheckRouteExists(name string, route *ipam_models.NetworkRoute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)
		params := virtual_private_cloud.GetRouteParams{
			StackID: config.StackID,
			Context: context.Background(),
		}
		var err error
		params.NetworkID, params.RouteID, err = parseRouteID(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("could not retrieve network: %v", err)
		}
		found, err := config.edgeComputeNetworking.VirtualPrivateCloud.GetRoute(&params, nil)
		if err != nil {
			return fmt.Errorf("could not retrieve network: %v", err)
		}

		*route = *found.Payload.Route

		return nil
	}
}

func testAccComputeRouteCheckDestroy() resource.TestCheckFunc {
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
