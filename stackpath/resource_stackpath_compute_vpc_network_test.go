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

func TestAccComputeVPCNetwork(t *testing.T) {
	t.Parallel()

	network := &ipam_models.NetworkNetwork{}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeNetworkCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: `
				resource "stackpath_compute_vpc_network" "foo" {
				  name = "test-tf-network-1"
				  slug = "test-tf-network-1"
				  root_subnet = "10.0.0.0/8"
				}`,
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckNetworkExists("stackpath_compute_vpc_network.foo", network),
					testAccCheckNetworkMatch(network, &ipam_models.NetworkNetwork{
						Name:                     "test-tf-network-1",
						Slug:                     "test-tf-network-1",
						RootSubnet:               "10.0.0.0/8",
						VirtualNetworkIdentifier: 9001,
						Metadata: &ipam_models.NetworkMetadata{
							Version: "1",
						},
					}),
				),
			},
			{
				Config: `
				resource "stackpath_compute_vpc_network" "foo" {
				  name = "test-tf-network-1"
				  slug = "test-tf-network-1"
				  root_subnet = "10.0.0.0/8"
				  labels = {
					  "new-label" = "value1"
				  }
				  annotations = {
					  "new-annotation" = "value1"
				  }
				}`,
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckNetworkExists("stackpath_compute_vpc_network.foo", network),
					testAccCheckNetworkMatch(network, &ipam_models.NetworkNetwork{
						Name:                     "test-tf-network-1",
						Slug:                     "test-tf-network-1",
						RootSubnet:               "10.0.0.0/8",
						VirtualNetworkIdentifier: 9001,
						Metadata: &ipam_models.NetworkMetadata{
							Version: "2",
							Labels: map[string]string{
								"new-label": "value1",
							},
							Annotations: map[string]string{
								"new-annotation": "value1",
							},
						},
					}),
				),
			},
			{
				Config: `
				resource "stackpath_compute_vpc_network" "foo" {
				  name = "test-tf-network-1"
				  slug = "test-tf-network-1"
				  root_subnet = "10.0.0.0/8"
				  labels = {
					  "new-label" = "value1"
				  }
				  annotations = {
					  "new-annotation" = "value1"
				  }
				}`,
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckNetworkExists("stackpath_compute_vpc_network.foo", network),
					testAccCheckNetworkMatch(network, &ipam_models.NetworkNetwork{
						Name:                     "test-tf-network-1",
						Slug:                     "test-tf-network-1",
						RootSubnet:               "10.0.0.0/8",
						VirtualNetworkIdentifier: 9001,
						Metadata: &ipam_models.NetworkMetadata{
							Version: "2",
							Labels: map[string]string{
								"new-label": "value1",
							},
							Annotations: map[string]string{
								"new-annotation": "value1",
							},
						},
					}),
				),
			},
		},
	})
}

func testAccComputeCheckNetworkExists(name string, network *ipam_models.NetworkNetwork) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)
		found, err := config.edgeComputeNetworking.VirtualPrivateCloud.GetNetwork(&virtual_private_cloud.GetNetworkParams{
			NetworkID: rs.Primary.ID,
			StackID:   config.StackID,
			Context:   context.Background(),
		}, nil)
		if err != nil {
			return fmt.Errorf("could not retrieve network: %v", err)
		}

		*network = *found.Payload.Network

		return nil
	}
}

func testAccCheckNetworkMatch(got, want *ipam_models.NetworkNetwork) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if want == nil && got != nil {
			return errors.New("mismatch network. got=non-nil want=nil")
		}
		if want != nil && got == nil {
			return errors.New("mismatch network. got=nil want=non-nil")
		}
		if want.Name != got.Name {
			return fmt.Errorf("mismatch network.Name. got=%s want=%s", got.Name, want.Name)
		}
		if want.Slug != got.Slug {
			return fmt.Errorf("mismatch network.Slug. got=%s want=%s", got.Slug, want.Slug)
		}
		if want.RootSubnet != got.RootSubnet {
			return fmt.Errorf("mismatch network.RootSubnet. got=%s want=%s", got.RootSubnet, want.RootSubnet)
		}
		if (want.VirtualNetworkIdentifier > 0) != (got.VirtualNetworkIdentifier > 0) {
			return fmt.Errorf("mismatch network.VirtualNetworkIdentifier. got=%d want=%d", got.VirtualNetworkIdentifier, want.VirtualNetworkIdentifier)
		}

		if want.Metadata == nil && got.Metadata != nil {
			return errors.New("mismatch network.Metadata. got=non-nil want=nil")
		}
		if want.Metadata != nil && got.Metadata == nil {
			return errors.New("mismatch network.Metadata. got=nil want=non-nil")
		}
		if want.Metadata != nil {
			if want.Metadata.Version != got.Metadata.Version {
				return fmt.Errorf("mismatch network.Metadata.Version. got=%s want=%s", got.Metadata.Version, want.Metadata.Version)
			}
			if !reflect.DeepEqual(want.Metadata.Labels, got.Metadata.Labels) {
				return fmt.Errorf("mismatch network.Metadata.Labels. got=%#v want=%#v", got.Metadata.Labels, want.Metadata.Labels)
			}
			if !reflect.DeepEqual(want.Metadata.Annotations, got.Metadata.Annotations) {
				return fmt.Errorf("mismatch network.Metadata.Annotations. got=%#v want=%#v", got.Metadata.Annotations, want.Metadata.Annotations)
			}
		}
		return nil
	}
}

func testAccComputeNetworkCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "stackpath_compute_vpc_network" {
				continue
			}

			resp, err := config.edgeComputeNetworking.VirtualPrivateCloud.GetNetwork(&virtual_private_cloud.GetNetworkParams{
				StackID:   config.StackID,
				NetworkID: rs.Primary.ID,
				Context:   context.Background(),
			}, nil)
			// Since compute workloads are deleted asynchronously, we want to look at the fact that
			// the deleteRequestedAt timestamp was set on the workload. This field is used to indicate
			// that the workload is being deleted.
			if err == nil && *resp.Payload.Network.Metadata.DeleteRequestedAt == strfmt.NewDateTime() {
				return fmt.Errorf("network still exists: %v", rs.Primary.ID)
			}
		}

		return nil
	}
}
