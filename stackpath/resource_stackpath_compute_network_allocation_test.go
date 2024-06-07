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
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_client/allocations"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"
)

func TestAccComputeNetworkAllocation(t *testing.T) {
	t.Parallel()

	allocation := &ipam_models.V1Allocation{}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeNetworkAllocationCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: `
				resource "stackpath_compute_network_allocation" "foo" {
				  name = "test-tf-network-allocation-1"
				  slug = "test-tf-network-allocation-1"
					labels = {
						"app" = "my-compute-network-allocation-selector"
					}
					allocation_class = "stackpath-edge/unicast"
					ip_family = "IPv4"
					prefix_length = 32
					reclaim_policy {
						action = "RETAIN"
					}
					selectors {
						key = "cityCode"
						operator = "in"
						values = ["EC4LAB01"]
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckNetworkAllocationExists("stackpath_compute_network_allocation.foo", allocation),
					testAccCheckNetworkAllocationMatch(allocation, &ipam_models.V1Allocation{
						Name: "test-tf-network-allocation-1",
						Slug: "test-tf-network-allocation-1",
						Metadata: &ipam_models.Metav1Metadata{
							Version: "1",
							Labels: ipam_models.Metav1StringMapEntry{
								"app": "my-compute-network-allocation-selector",
							},
						},
						Spec: &ipam_models.V1AllocationSpec{
							AllocationClass: "stackpath-edge/unicast",
							IPFamily:        ipam_models.NewV1IPFamily("IPv4"),
							PrefixLength:    int32(32),
							ReclaimPolicy: &ipam_models.V1ReclaimPolicy{
								Action: ipam_models.NewReclaimPolicyReclaimPolicyAction("RETAIN"),
							},
							Selectors: []*ipam_models.Metav1MatchExpression{
								{
									Key:      "cityCode",
									Operator: "in",
									Values:   []string{"EC4LAB01"},
								},
							},
						},
					}),
				),
			},
		},
	})
}

func testAccComputeCheckNetworkAllocationExists(name string, allocation *ipam_models.V1Allocation) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)

		_, slug, err := parseAllocationID(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("failed to parse allocation ID (%s): %v", rs.Primary.ID, err)
		}

		found, err := config.edgeComputeNetworking.Allocations.GetAllocation(&allocations.GetAllocationParams{
			AllocationSlug: slug,
			StackID:        config.StackID,
			Context:        context.Background(),
		}, nil)
		if err != nil {
			return fmt.Errorf("could not retrieve network allocation: %v", err)
		}

		*allocation = *found.Payload.Allocation

		return nil
	}
}

func testAccCheckNetworkAllocationMatch(got, want *ipam_models.V1Allocation) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if want == nil && got != nil {
			return errors.New("mismatch network allocation. got=non-nil want=nil")
		}
		if want != nil && got == nil {
			return errors.New("mismatch network allocation. got=nil want=non-nil")
		}
		if want.Name != got.Name {
			return fmt.Errorf("mismatch allocation.Name. got=%s want=%s", got.Name, want.Name)
		}
		if want.Slug != got.Slug {
			return fmt.Errorf("mismatch allocation.Slug. got=%s want=%s", got.Slug, want.Slug)
		}
		if want.Metadata == nil && got.Metadata != nil {
			return errors.New("mismatch allocation.Metadata. got=non-nil want=nil")
		}
		if want.Metadata != nil && got.Metadata == nil {
			return errors.New("mismatch allocation.Metadata. got=nil want=non-nil")
		}
		if want.Metadata != nil {
			if want.Metadata.Version != got.Metadata.Version {
				return fmt.Errorf("mismatch allocation.Metadata.Version. got=%s want=%s", got.Metadata.Version, want.Metadata.Version)
			}
			if !reflect.DeepEqual(want.Metadata.Labels, got.Metadata.Labels) {
				return fmt.Errorf("mismatch allocation.Metadata.Labels. got=%#v want=%#v", got.Metadata.Labels, want.Metadata.Labels)
			}
			if !reflect.DeepEqual(want.Metadata.Annotations, got.Metadata.Annotations) {
				return fmt.Errorf("mismatch allocation.Metadata.Annotations. got=%#v want=%#v", got.Metadata.Annotations, want.Metadata.Annotations)
			}
		}

		if want.Spec == nil && got.Spec != nil {
			return errors.New("mismatch allocation.Spec. got=non-nil want=nil")
		}
		if want.Spec != nil && got.Spec == nil {
			return errors.New("mismatch allocation.Spec. got=nil want=non-nil")
		}
		if want.Spec != nil {
			if !reflect.DeepEqual(want.Spec, got.Spec) {
				return fmt.Errorf("mismatch allocation.Spec. got=%#v want=%#v", got.Spec, want.Spec)
			}
		}
		return nil
	}
}

func testAccComputeNetworkAllocationCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "stackpath_compute_network_allocation" {
				continue
			}

			_, slug, err := parseAllocationID(rs.Primary.ID)
			if err != nil {
				return fmt.Errorf("failed to parse allocation ID (%s): %v", rs.Primary.ID, err)
			}

			resp, err := config.edgeComputeNetworking.Allocations.GetAllocation(&allocations.GetAllocationParams{
				StackID:        config.StackID,
				AllocationSlug: slug,
				Context:        context.Background(),
			}, nil)
			// Since compute workloads are deleted asynchronously, we want to look at the fact that
			// the deleteRequestedAt timestamp was set on the workload. This field is used to indicate
			// that the workload is being deleted.
			if err == nil && *resp.Payload.Allocation.Metadata.DeleteRequestedAt == strfmt.NewDateTime() {
				return fmt.Errorf("network allocation still exists: %v", rs.Primary.ID)
			}
		}

		return nil
	}
}
