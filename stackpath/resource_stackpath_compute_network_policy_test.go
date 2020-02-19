package stackpath

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/ipam/ipam_client/network_policies"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
)

func TestAccComputeNetworkPolicy(t *testing.T) {
	t.Parallel()

	networkPolicy := &ipam_models.V1NetworkPolicy{}

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeNetworkPolicyCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkPolicyConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckNetworkPolicyExists("stackpath_compute_network_policy.foo", networkPolicy),
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
			if err == nil && resp.Payload.NetworkPolicy.Metadata.DeleteRequestedAt == strfmt.NewDateTime() {
				return fmt.Errorf("network policy still exists: %v", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccComputeNetworkPolicyConfigBasic() string {
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
    # the source CIDR range of 0.0.0.0/0 (all traffic) to
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
      ip_block {
        cidr = "0.0.0.0/0"
      }
    }
  }

  egress {
    description = "Allow all outbound traffic"
    action      = "ALLOW"
    to {
      ip_block {
        cidr = "0.0.0.0/0"
      }
    }
  }
}`)
}
