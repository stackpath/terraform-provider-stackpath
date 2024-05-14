package stackpath

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_client/operations"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// operation wait timeout in seconds
var OperationWaitTimeout = 30 * time.Second

// convert from the terraform data structure to the allocation data structure we need for API calls
func convertComputeNetworkAllocation(data *schema.ResourceData) *ipam_models.V1Allocation {
	return &ipam_models.V1Allocation{
		Name: data.Get("name").(string),
		Slug: data.Get("slug").(string),
		Metadata: &ipam_models.Metav1Metadata{
			Annotations: convertToMetaV1StringMap(data.Get("annotations").(map[string]interface{})),
			Labels:      convertToMetaV1StringMap(data.Get("labels").(map[string]interface{})),
			Version:     data.Get("version").(string),
		},
		Spec: &ipam_models.V1AllocationSpec{
			AllocationClass: data.Get("allocation_class").(string),
			PrefixLength:    data.Get("prefix_length").(int32),
			IPFamily:        convertComputeNetworkAllocationIPFamily(data.Get("ip_family")),
			ReclaimPolicy:   convertComputeNetworkAllocationReclaimPolicy(data.Get("reclaim_policy").(interface{})),
		},
	}
}

func convertComputeNetworkAllocationIPFamily(f interface{}) *ipam_models.StackpathschemanetworkIPFamily {
	ipFamily := ipam_models.StackpathschemanetworkIPFamily(f.(string))
	return &ipFamily
}

func convertComputeNetworkAllocationReclaimPolicy(p interface{}) *ipam_models.V1ReclaimPolicy {
	reclaimPolicy := &ipam_models.V1ReclaimPolicy{}
	if p.(map[string]string)["action"] != "" {
		action := ipam_models.ReclaimPolicyReclaimPolicyAction(p.(map[string]string)["action"])
		reclaimPolicy.Action = &action
	}
	if p.(map[string]string)["idle_retention_period"] != "" {
		reclaimPolicy.IdleRetentionPeriod = p.(map[string]string)["idle_retention_period"]
	}

	return reclaimPolicy
}

func flattenComputeNetworkAllocationReclaimPolicy(r *ipam_models.V1ReclaimPolicy) map[string]string {
	return map[string]string{
		"action":                string(*r.Action),
		"idle_retention_period": r.IdleRetentionPeriod,
	}
}

func flattenComputeNetworkAllocationStatus(s *ipam_models.V1AllocationStatus) map[string]interface{} {
	conditions := make([]interface{}, 0, len(s.Conditions))
	for _, condition := range s.Conditions {
		conditions = append(
			conditions,
			map[string]interface{}{
				"type":                 condition.Type,
				"status":               condition.Status,
				"observed_version":     condition.ObservedVersion,
				"last_transition_time": condition.LastTransitionTime,
				"reason":               condition.Reason,
				"message":              condition.Message,
			},
		)
	}

	return map[string]interface{}{
		"prefix":            s.Prefix,
		"parent_allocation": s.ParentAllocation,
		"conditions":        conditions,
	}
}

func waitForIPAMOperationToBeDone(ctx context.Context, name string, meta interface{}) (ipam_models.V1Operation, error) {
	config := meta.(*Config)

	// operation name is expected to be in stacks/{stack_id}/operations/{operation_name}
	// format, extract actual operation name as GET /operation expects actual operation name
	name = extractOperationName(name)
	if name == "" {
		return ipam_models.V1Operation{}, fmt.Errorf("unable to extract operation name")
	}

	timeout := time.After(OperationWaitTimeout)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			resp, err := config.edgeComputeNetworking.Operations.GetOperation(&operations.GetOperationParams{
				StackID:       config.StackID,
				OperationName: name,
				Context:       ctx,
			}, nil)

			if err != nil {
				return ipam_models.V1Operation{}, fmt.Errorf("failed to get operation: %v", NewStackPathError(err))
			} else if resp.Payload.Done {
				return *resp.Payload, nil
			}
		case <-timeout:
			return ipam_models.V1Operation{}, fmt.Errorf("timed out waiting for operation to be done")
		}
	}
}

func extractOperationName(name string) string {
	pattern := `stacks/[a-f0-9\-]+/operations/(operation-[a-f0-9\-]+)`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(name)

	if len(match) == 0 {
		// if name passed in does not match with expected format then return as is
		return name
	}

	return match[1]
}
