package stackpath

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/internal/client"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/internal/models"
)

func TestAccComputeNetworkPolicy(t *testing.T) {
	t.Parallel()

	networkPolicy := &models.V1NetworkPolicy{}

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccComputeNetworkPolicyCheckDestroy(),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeNetworkPolicyConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeCheckNetworkPolicyExists("stackpath_compute_network_policy.foo", networkPolicy),
				),
			},
		},
	})
}

func testAccComputeCheckNetworkPolicyExists(name string, policy *models.V1NetworkPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)
		found, err := config.ipam.GetNetworkPolicy(&client.GetNetworkPolicyParams{
			NetworkPolicyID: rs.Primary.ID,
			StackID:         config.Stack,
			Context:         context.Background(),
		}, nil)
		if err != nil {
			return fmt.Errorf("Could not retrieve network policy: %v", err)
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

			resp, err := config.ipam.GetNetworkPolicy(&client.GetNetworkPolicyParams{
				StackID:         config.Stack,
				NetworkPolicyID: rs.Primary.ID,
				Context:         context.Background(),
			}, nil)
			// Since compute workloads are deleted asyncronously, we want to look at the fact that
			// the deleteRequestedAt timestamp was set on the workload. This field is used to indicate
			// that the workload is being deleted.
			if err == nil && resp.Payload.NetworkPolicy.Metadata.DeleteRequestedAt == nil {
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
