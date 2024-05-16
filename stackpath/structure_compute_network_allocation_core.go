package stackpath

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_client/operations"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"
)

// operation wait timeout in seconds
var OperationWaitTimeout = 30 * time.Second

func formatAllocationID(stackID, slug string) string {
	return stackID + "/" + slug
}
func parseAllocationID(id string) (stackID, slug string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		err = fmt.Errorf("found %d parts instead of 2", len(parts))
	} else {
		stackID = parts[0]
		slug = parts[1]
	}
	return
}

func convertComputeNetworkAllocationIPFamily(f interface{}) *ipam_models.V1IPFamily {
	ipFamily := ipam_models.V1IPFamily(f.(string))
	return &ipFamily
}

func convertComputeNetworkAllocationReclaimPolicy(p interface{}) *ipam_models.V1ReclaimPolicy {
	if len(p.([]interface{})) == 0 {
		return nil
	}

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

func waitForIPAMOperationToBeDone(ctx context.Context, name string, meta interface{}) (*ipam_models.V1Operation, error) {
	config := meta.(*Config)

	// operation name is expected to be in stacks/{stack_id}/operations/{operation_name}
	// format, extract actual operation name as GET /operation expects actual operation name
	name = extractOperationName(name)
	if name == "" {
		return nil, fmt.Errorf("unable to extract operation name")
	}

	// (TODO)- WaitOperation is not working currently due to missing kong paths
	// configuration for /ipam/v1/stacks/{stack_id}/operations/{operation_name}/wait api
	// endpoint. replace below ticker based implementation to wait for operation
	// completion with WaitOperation api call once it starts working.

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
