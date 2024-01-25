package stackpath

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_client/network_policies"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"
)

func TestAccComputeNetworkPolicy(t *testing.T) {
	t.Parallel()

	networkPolicy := &ipam_models.V1NetworkPolicy{}

	cidrList := []string{"0.0.0.0/0"}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeNetworkPolicyCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkPolicyConfigBasic(cidrList),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckNetworkPolicyExists("stackpath_compute_network_policy.foo", networkPolicy),
					testAccComputeCheckNetworkPolicyIpBlocks("ingress", cidrList, networkPolicy),
					testAccComputeCheckNetworkPolicyIpBlocks("egress", cidrList, networkPolicy),
				),
			},
		},
	})
}

func TestAccComputeNetworkPolicyDualStack(t *testing.T) {
	t.Parallel()

	networkPolicy := &ipam_models.V1NetworkPolicy{}

	cidrList := []string{"0.0.0.0/0", "::/0"}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeNetworkPolicyCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkPolicyConfigBasic(cidrList),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckNetworkPolicyExists("stackpath_compute_network_policy.foo", networkPolicy),
					testAccComputeCheckNetworkPolicyIpBlocks("ingress", cidrList, networkPolicy),
					testAccComputeCheckNetworkPolicyIpBlocks("egress", cidrList, networkPolicy),
				),
			},
		},
	})
}

func testAccComputeCheckNetworkPolicyExists(name string, policy *ipam_models.V1NetworkPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)
		found, err := config.edgeComputeNetworking.NetworkPolicies.GetNetworkPolicy(&network_policies.GetNetworkPolicyParams{
			NetworkPolicyID: rs.Primary.ID,
			StackID:         config.StackID,
			Context:         context.Background(),
		}, nil)
		if err != nil {
			return fmt.Errorf("could not retrieve network policy: %v", err)
		}

		*policy = *found.Payload.NetworkPolicy

		return nil
	}
}

func testAccComputeCheckNetworkPolicyIpBlocks(
	policyType string,
	expectedCIDRs []string,
	policy *ipam_models.V1NetworkPolicy,
) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		var want []*ipam_models.V1IPBlock
		for _, cidr := range expectedCIDRs {
			want = append(want, &ipam_models.V1IPBlock{
				Cidr: cidr,
			})
		}

		var got []*ipam_models.V1IPBlock
		if policyType == "ingress" {
			got = policy.Spec.Ingress[0].From.IPBlock
		} else if policyType == "egress" {
			got = policy.Spec.Egress[0].To.IPBlock
		}

		if got == nil {
			return errors.New("got nil host rule ip blocks")
		}

		if len(got) != len(want) {
			return fmt.Errorf("mismatch in length of ipblocks list, got length=%d want length=%v", len(got), len(want))
		}

		for i, _ := range want {
			if got[i].Cidr != want[i].Cidr {
				return fmt.Errorf("mismatch host rule ip block. got=%v want=%v", got[i], want[i])
			}
		}
		return nil
	}
}

func testAccComputeNetworkPolicyCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "stackpath_compute_workload" {
				continue
			}

			resp, err := config.edgeComputeNetworking.NetworkPolicies.GetNetworkPolicy(&network_policies.GetNetworkPolicyParams{
				StackID:         config.StackID,
				NetworkPolicyID: rs.Primary.ID,
				Context:         context.Background(),
			}, nil)
			// Since compute workloads are deleted asynchronously, we want to look at the fact that
			// the deleteRequestedAt timestamp was set on the workload. This field is used to indicate
			// that the workload is being deleted.
			if err == nil && *resp.Payload.NetworkPolicy.Metadata.DeleteRequestedAt == strfmt.NewDateTime() {
				return fmt.Errorf("network policy still exists: %v", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccComputeNetworkPolicyConfigBasic(cidrList []string) string {
	ipBlocks := ""
	for _, cidr := range cidrList {
		ipBlocks = ipBlocks + fmt.Sprintf(`
		ip_block {
			cidr = "%s"
		}
		`, cidr)
	}

	return fmt.Sprintf(`
resource "stackpath_compute_network_policy" "foo" {
  name = "test-terraform-workload-allow-port-80"
  slug = "test-terraform-workload-allow-port-80"

  instance_selector {
    key      = "workload.platform.stackpath.net/workload-slug"
    operator = "in"
    values = [
      "my-terraform-workload"
    ]
  }

  policy_types = ["INGRESS", "EGRESS"]

  priority = 1000

  ingress {
    # Configure the network policy to allow traffic from
    # the source CIDR range of 0.0.0.0/0 or ::/0 (all traffic) to
    # hit port 80.
    description = "Allow all port 80 traffic"
    action      = "ALLOW"
    protocol {
      tcp {
        # configure the destination ports that should be allowed
        destination_ports = [80]
      }
    }
    from {
      %[1]s
    }
  }

  egress {
    description = "Allow all outbound traffic"
    action      = "ALLOW"
    to {
      %[1]s
    }
  }
}`, ipBlocks)
}
