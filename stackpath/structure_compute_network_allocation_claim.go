package stackpath

import (
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// convert from the terraform data structure to the allocation data structure we need for API calls
func convertComputeNetworkAllocationClaim(data *schema.ResourceData) *ipam_models.V1AllocationClaim {
	return &ipam_models.V1AllocationClaim{
		Name: data.Get("name").(string),
		Slug: data.Get("slug").(string),
		Metadata: &ipam_models.Metav1Metadata{
			Annotations: convertToMetaV1StringMap(data.Get("annotations").(map[string]interface{})),
			Labels:      convertToMetaV1StringMap(data.Get("labels").(map[string]interface{})),
			Version:     data.Get("version").(string),
		},
		Spec: &ipam_models.V1AllocationClaimSpec{
			PrefixLength:    int32(data.Get("prefix_length").(int)),
			IPFamily:        convertComputeNetworkAllocationIPFamily(data.Get("ip_family")),
			ReclaimPolicy:   convertComputeNetworkAllocationReclaimPolicy(data.Get("reclaim_policy")),
			ResourceBinding: convertComputeNetworkAllocationResourceBinding(data.Get("resource_binding")),
			Allocation:      convertComputeNetworkAllocationClaimSpecAllocation(data.Get("allocation")),
		},
	}
}

// convert from the terraform data structure to the allocation data structure we need for update API
func convertComputeNetworkAllocationClaimUpdate(data *schema.ResourceData) *ipam_models.V1AllocationClaim {
	// prepare data structure with only fields which are allowed to update, passing in
	// any additional fields which are not allowed to update causes update api to throw
	// validation error. hence to allow succesful updates for data changes we are preparing
	// request body only with fields which are allowed to pass in.
	return &ipam_models.V1AllocationClaim{}
}

func convertComputeNetworkAllocationResourceBinding(p interface{}) *ipam_models.V1TypedResourceReference {
	if len(p.([]interface{})) == 0 {
		return nil
	}

	resourceBinding := &ipam_models.V1TypedResourceReference{}
	resourceBindingData := p.([]interface{})[0].(map[string]interface{})
	if resourceBindingData["type"].(string) != "" {
		resourceBinding.Type = resourceBindingData["type"].(string)
	}

	if resourceBindingData["name"].(string) != "" {
		resourceBinding.Name = resourceBindingData["name"].(string)
	}

	return resourceBinding
}

func convertComputeNetworkAllocationClaimSpecAllocation(p interface{}) *ipam_models.AllocationClaimSpecAllocationClaimSpecAllocation {
	if len(p.([]interface{})) == 0 {
		return nil
	}

	allocationClaimSpecAllocation := &ipam_models.AllocationClaimSpecAllocationClaimSpecAllocation{}
	allocationClaimSpecAllocationData := p.([]interface{})[0].(map[string]interface{})
	if name, ok := allocationClaimSpecAllocationData["name"]; ok {
		allocationClaimSpecAllocation.Name = name.(string)
	}

	if selector, ok := allocationClaimSpecAllocationData["selector"]; ok {
		allocationClaimSpecAllocation.Selector = convertAllocationClaimSpecAllocationSelector(selector)
	}

	if template, ok := allocationClaimSpecAllocationData["template"]; ok {
		allocationClaimSpecAllocation.Template = convertAllocationClaimSpecAllocationTemplate(template)
	}

	return allocationClaimSpecAllocation
}

func convertAllocationClaimSpecAllocationSelector(p interface{}) *ipam_models.AllocationClaimSpecAllocationClaimSpecAllocationSelector {
	if len(p.([]interface{})) == 0 {
		return nil
	}

	selector := &ipam_models.AllocationClaimSpecAllocationClaimSpecAllocationSelector{}
	selectorData := p.([]interface{})[0].(map[string]interface{})

	if allocationClass, ok := selectorData["allocation_class"]; ok {
		selector.AllocationClass = allocationClass.(string)
	}

	if matchExpressions, ok := selectorData["match_expressions"]; ok {
		selector.MatchExpressions = convertComputeMetaV1MatchExpression(matchExpressions.([]interface{}))
	}

	return selector
}

func convertAllocationClaimSpecAllocationTemplate(p interface{}) *ipam_models.V1Allocation {
	if len(p.([]interface{})) == 0 {
		return nil
	}

	allocation := &ipam_models.V1Allocation{}
	allocationData := p.([]interface{})[0].(map[string]interface{})

	allocation.Spec = &ipam_models.V1AllocationSpec{
		AllocationClass: allocationData["allocation_class"].(string),
		PrefixLength:    int32(allocationData["prefix_length"].(int)),
		IPFamily:        convertComputeNetworkAllocationIPFamily(allocationData["ip_family"]),
		ReclaimPolicy:   convertComputeNetworkAllocationReclaimPolicy(allocationData["reclaim_policy"]),
		Selectors:       convertComputeMetaV1MatchExpression(allocationData["selectors"].([]interface{})),
	}

	return allocation
}

func flattenComputeNetworkAllocationClaimStatus(s *ipam_models.V1AllocationClaimStatus) []interface{} {
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

	claimStatus := map[string]interface{}{
		"prefix":     s.Prefix,
		"allocation": s.Allocation,
	}

	if s.ResourceBinding != nil {
		claimStatus["resource_binding"] = []interface{}{
			map[string]interface{}{
				"type": s.ResourceBinding.Type,
				"name": s.ResourceBinding.Name,
			},
		}
	}
	if len(conditions) > 0 {
		claimStatus["conditions"] = conditions
	}

	return []interface{}{claimStatus}
}

func flattenComputeNetworkAllocationResourceBinding(r *ipam_models.V1TypedResourceReference) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"type": r.Type,
			"name": r.Name,
		},
	}
}

func flattenComputeNetworkAllocationClaimSpecAllocation(r *ipam_models.AllocationClaimSpecAllocationClaimSpecAllocation) []interface{} {
	allocation := map[string]interface{}{}

	if r != nil {
		if r.Name != "" {
			allocation["name"] = r.Name
		}

		if r.Selector != nil {
			allocation["selector"] = flattenAllocationClaimSpecAllocationSelector(r.Selector)
		}
		if r.Template != nil {
			allocation["template"] = flattenAllocationClaimSpecAllocationTemplate(r.Template)
		}
	}

	return []interface{}{allocation}
}

func flattenAllocationClaimSpecAllocationSelector(r *ipam_models.AllocationClaimSpecAllocationClaimSpecAllocationSelector) []interface{} {
	selector := map[string]interface{}{
		"allocation_class": r.AllocationClass,
	}

	if len(r.MatchExpressions) > 0 {
		selector["match_expressions"] = flattenComputeMetaV1MatchExpressionsOrdered(r.MatchExpressions)
	}

	return []interface{}{selector}
}

func flattenAllocationClaimSpecAllocationTemplate(r *ipam_models.V1Allocation) []interface{} {
	if r.Spec == nil {
		return []interface{}{}
	}

	template := map[string]interface{}{
		"allocation_class": r.Spec.AllocationClass,
		"prefix_length":    int(r.Spec.PrefixLength),
		"ip_family":        string(*r.Spec.IPFamily),
		"reclaim_policy":   flattenComputeNetworkAllocationReclaimPolicy(r.Spec.ReclaimPolicy),
	}

	if len(r.Spec.Selectors) > 0 {
		template["selectors"] = flattenComputeMetaV1MatchExpressionsOrdered(r.Spec.Selectors)
	}

	return []interface{}{template}
}
