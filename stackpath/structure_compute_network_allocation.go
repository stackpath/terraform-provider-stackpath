package stackpath

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_client/operations"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// operation wait timeout in seconds
var OperationWaitTimeout = 30 * time.Second

func formatAllocationID(stackID, allocationSlug string) string {
	return stackID + "/" + allocationSlug
}
func parseAllocationID(id string) (stackID, allocationSlug string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		err = fmt.Errorf("found %d parts instead of 2", len(parts))
	} else {
		stackID = parts[0]
		allocationSlug = parts[1]
	}
	return
}

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
			PrefixLength:    int32(data.Get("prefix_length").(int)),
			IPFamily:        convertComputeNetworkAllocationIPFamily(data.Get("ip_family")),
			ReclaimPolicy:   convertComputeNetworkAllocationReclaimPolicy(data.Get("reclaim_policy")),
			Selectors:       convertComputeMetaV1MatchExpression(data.Get("selectors").([]interface{})),
		},
	}
}

// convert from the terraform data structure to the allocation data structure we need for update API
func convertComputeNetworkAllocationUpdate(data *schema.ResourceData) *ipam_models.V1Allocation {
	// prepare data structure with only fields which are allowed to update, passing in
	// any additional fields which are not allowed to update causes update api to throw
	// validation error. hence to allow succesful updates for data changes we are preparing
	// request body only with fields which are allowed to pass in.
	return &ipam_models.V1Allocation{
		Name: data.Get("name").(string),
		Metadata: &ipam_models.Metav1Metadata{
			Annotations: convertToMetaV1StringMap(data.Get("annotations").(map[string]interface{})),
			Labels:      convertToMetaV1StringMap(data.Get("labels").(map[string]interface{})),
		},
		Spec: &ipam_models.V1AllocationSpec{
			ReclaimPolicy: convertComputeNetworkAllocationReclaimPolicy(data.Get("reclaim_policy")),
		},
	}
}

func convertComputeNetworkAllocationIPFamily(f interface{}) *ipam_models.V1IPFamily {
	ipFamily := ipam_models.V1IPFamily(f.(string))
	return &ipFamily
}

func convertComputeNetworkAllocationReclaimPolicy(p interface{}) *ipam_models.V1ReclaimPolicy {
	reclaimPolicyData := p.([]interface{})[0].(map[string]interface{})
	reclaimPolicy := &ipam_models.V1ReclaimPolicy{}
	if reclaimPolicyData["action"].(string) != "" {
		action := ipam_models.ReclaimPolicyReclaimPolicyAction(reclaimPolicyData["action"].(string))
		reclaimPolicy.Action = &action
	}
	if reclaimPolicyData["idle_retention_period"].(string) != "" {
		reclaimPolicy.IdleRetentionPeriod = reclaimPolicyData["idle_retention_period"].(string)
	}

	return reclaimPolicy
}

func flattenComputeNetworkAllocationReclaimPolicy(r *ipam_models.V1ReclaimPolicy) []interface{} {
	var action string
	if r.Action != nil {
		action = string(*r.Action)
	}

	return []interface{}{
		map[string]interface{}{
			"action":                action,
			"idle_retention_period": r.IdleRetentionPeriod,
		},
	}
}

func flattenComputeNetworkAllocationStatus(s *ipam_models.V1AllocationStatus) []interface{} {
	conditions := make([]interface{}, 0, len(s.Conditions))
	for _, condition := range s.Conditions {
		status := ""
		if condition.Status != nil {
			status = string(*condition.Status)
		}

		conditions = append(
			conditions,
			map[string]interface{}{
				"type":                 condition.Type,
				"status":               status,
				"observed_version":     condition.ObservedVersion,
				"last_transition_time": condition.LastTransitionTime.String(),
				"reason":               condition.Reason,
				"message":              condition.Message,
			},
		)
	}

	return []interface{}{
		map[string]interface{}{
			"prefix":            s.Prefix,
			"parent_allocation": s.ParentAllocation,
			"conditions":        conditions,
		},
	}
}

func waitForIPAMOperationToBeDone(ctx context.Context, name string, meta interface{}) (*ipam_models.V1Operation, error) {
	config := meta.(*Config)

	// operation name is expected to be in stacks/{stack_id}/operations/{operation_name}
	// format, extract actual operation name as GET /operation expects actual operation name
	name = extractOperationName(name)
	if name == "" {
		return nil, fmt.Errorf("unable to extract operation name")
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
				return nil, fmt.Errorf("failed to get operation: %v", NewStackPathError(err))
			}

			if resp.Payload.Done {
				return resp.Payload, nil
			}
		case <-timeout:
			return nil, fmt.Errorf("timed out waiting for operation to be done")
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
