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

func TestAccComputeNetworkAllocationClaim(t *testing.T) {
	t.Parallel()

	allocationclaim := &ipam_models.V1AllocationClaim{}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeNetworkAllocationClaimCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: `
				resource "stackpath_compute_network_allocation_claim" "foo" {
				  name = "test-tf-network-allocationclaim-1"
				  slug = "test-tf-network-allocationclaim-1"
					labels = {
						"app" = "my-compute-network-allocationclaim-selector"
					}
					ip_family = "IPv4"
					prefix_length = 32
					reclaim_policy {
						action = "RETAIN"
					}
					allocation {
						template {
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
						}
					}				
				}`,
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckNetworkAllocationClaimExists("stackpath_compute_network_allocation_claim.foo", allocationclaim),
					testAccCheckNetworkAllocationClaimMatch(allocationclaim, &ipam_models.V1AllocationClaim{
						Name: "test-tf-network-allocationclaim-1",
						Slug: "test-tf-network-allocationclaim-1",
						Metadata: &ipam_models.Metav1Metadata{
							Version: "2",
							Labels: ipam_models.Metav1StringMapEntry{
								"app": "my-compute-network-allocationclaim-selector",
							},
						},
						Spec: &ipam_models.V1AllocationClaimSpec{
							IPFamily:     ipam_models.NewV1IPFamily("IPv4"),
							PrefixLength: int32(32),
							ReclaimPolicy: &ipam_models.V1ReclaimPolicy{
								Action: ipam_models.NewReclaimPolicyReclaimPolicyAction("RETAIN"),
							},
							Allocation: &ipam_models.AllocationClaimSpecAllocationClaimSpecAllocation{
								Template: &ipam_models.V1Allocation{
									Name: "test-tf-network-allocationclaim-1",
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
								},
							},
						},
					}),
				),
			},
		},
	})
}

func testAccComputeCheckNetworkAllocationClaimExists(name string, allocationclaim *ipam_models.V1AllocationClaim) resource.TestCheckFunc {
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
			return fmt.Errorf("failed to parse allocationclaim ID (%s): %v", rs.Primary.ID, err)
		}

		found, err := config.edgeComputeNetworking.Allocations.GetAllocationClaim(&allocations.GetAllocationClaimParams{
			AllocationClaimSlug: slug,
			StackID:             config.StackID,
			Context:             context.Background(),
		}, nil)
		if err != nil {
			return fmt.Errorf("could not retrieve network allocationclaim: %v", err)
		}

		*allocationclaim = *found.Payload.AllocationClaim

		return nil
	}
}

func testAccCheckNetworkAllocationClaimMatch(got, want *ipam_models.V1AllocationClaim) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if want == nil && got != nil {
			return errors.New("mismatch network allocationclaim. got=non-nil want=nil")
		}
		if want != nil && got == nil {
			return errors.New("mismatch network allocationclaim. got=nil want=non-nil")
		}
		if want.Name != got.Name {
			return fmt.Errorf("mismatch allocationclaim.Name. got=%s want=%s", got.Name, want.Name)
		}
		if want.Slug != got.Slug {
			return fmt.Errorf("mismatch allocationclaim.Slug. got=%s want=%s", got.Slug, want.Slug)
		}
		if want.Metadata == nil && got.Metadata != nil {
			return errors.New("mismatch allocationclaim.Metadata. got=non-nil want=nil")
		}
		if want.Metadata != nil && got.Metadata == nil {
			return errors.New("mismatch allocationclaim.Metadata. got=nil want=non-nil")
		}
		if want.Metadata != nil {
			if want.Metadata.Version != got.Metadata.Version {
				return fmt.Errorf("mismatch allocationclaim.Metadata.Version. got=%s want=%s", got.Metadata.Version, want.Metadata.Version)
			}
			if !reflect.DeepEqual(want.Metadata.Labels, got.Metadata.Labels) {
				return fmt.Errorf("mismatch allocationclaim.Metadata.Labels. got=%#v want=%#v", got.Metadata.Labels, want.Metadata.Labels)
			}
			if !reflect.DeepEqual(want.Metadata.Annotations, got.Metadata.Annotations) {
				return fmt.Errorf("mismatch allocationclaim.Metadata.Annotations. got=%#v want=%#v", got.Metadata.Annotations, want.Metadata.Annotations)
			}
		}

		if want.Spec == nil && got.Spec != nil {
			return errors.New("mismatch allocationclaim.Spec. got=non-nil want=nil")
		}
		if want.Spec != nil && got.Spec == nil {
			return errors.New("mismatch allocationclaim.Spec. got=nil want=non-nil")
		}
		if want.Spec != nil {
			if !reflect.DeepEqual(want.Spec.IPFamily, got.Spec.IPFamily) {
				return fmt.Errorf("mismatch allocationclaim.Spec.IPFamily got=%#v want=%#v", got.Spec.IPFamily, want.Spec.IPFamily)
			}
			if !reflect.DeepEqual(want.Spec.PrefixLength, got.Spec.PrefixLength) {
				return fmt.Errorf("mismatch allocationclaim.Spec.PrefixLength got=%#v want=%#v", got.Spec.PrefixLength, want.Spec.PrefixLength)
			}
			if !reflect.DeepEqual(want.Spec.ReclaimPolicy, got.Spec.ReclaimPolicy) {
				return fmt.Errorf("mismatch allocationclaim.Spec.ReclaimPolicy got=%#v want=%#v", got.Spec.ReclaimPolicy, want.Spec.ReclaimPolicy)
			}
			if !reflect.DeepEqual(want.Spec.Allocation, got.Spec.Allocation) {
				return fmt.Errorf("mismatch allocationclaim.Spec.Allocation got=%#v want=%#v", got.Spec.Allocation, want.Spec.Allocation)
			}
		}
		return nil
	}
}

func testAccComputeNetworkAllocationClaimCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "stackpath_compute_network_allocation_claim" {
				continue
			}

			_, slug, err := parseAllocationID(rs.Primary.ID)
			if err != nil {
				return fmt.Errorf("failed to parse allocationclaim ID (%s): %v", rs.Primary.ID, err)
			}

			resp, err := config.edgeComputeNetworking.Allocations.GetAllocationClaim(&allocations.GetAllocationClaimParams{
				StackID:             config.StackID,
				AllocationClaimSlug: slug,
				Context:             context.Background(),
			}, nil)
			// Since compute allocation claims are deleted asynchronously, we want to look at the fact that
			// the deleteRequestedAt timestamp was set on the allocation claim. This field is used to
			// indicate that the allocation claim is being deleted.
			if err == nil && *resp.Payload.AllocationClaim.Metadata.DeleteRequestedAt == strfmt.NewDateTime() {
				return fmt.Errorf("network allocation claim still exists: %v", rs.Primary.ID)
			}
		}

		return nil
	}
}
