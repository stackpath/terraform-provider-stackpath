package stackpath

import (
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
	// We allow updates to:
	//  - metadata.annotations
	//  - metadata.labels
	//  - name
	//  - spec.reclaimPolicy.action
	//  - spec.reclaimPolicy.idleRetentionPeriod

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
