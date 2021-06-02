package stackpath

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/workload/workload_models"
)

func convertComputeMatchExpression(data []interface{}) []*workload_models.V1MatchExpression {
	selectors := make([]*workload_models.V1MatchExpression, len(data))
	for i, s := range data {
		selector := s.(map[string]interface{})
		vals := make([]string, len(selector["values"].([]interface{})))
		for j, v := range selector["values"].([]interface{}) {
			vals[j] = v.(string)
		}
		selectors[i] = &workload_models.V1MatchExpression{
			Key:      selector["key"].(string),
			Operator: selector["operator"].(string),
			Values:   vals,
		}
	}
	return selectors
}

func convertToStringMap(data map[string]interface{}) workload_models.V1StringMapEntry {
	stringMap := make(workload_models.V1StringMapEntry, len(data))
	for k, v := range data {
		stringMap[k] = v.(string)
	}
	return stringMap
}

func convertToStringArray(data []interface{}) []string {
	s := make([]string, len(data))
	for i, c := range data {
		s[i] = c.(string)
	}
	return s
}

func convertWorkloadToIPAMMatchExpression(selectors []*workload_models.V1MatchExpression) []*ipam_models.NetworkMatchExpression {
	converted := make([]*ipam_models.NetworkMatchExpression, len(selectors))

	for i, selector := range selectors {
		converted[i] = &ipam_models.NetworkMatchExpression{
			Key:      selector.Key,
			Operator: selector.Operator,
			Values:   selector.Values,
		}
	}

	return converted
}

func convertIPAMToWorkloadMatchExpression(selectors []*ipam_models.NetworkMatchExpression) []*workload_models.V1MatchExpression {
	converted := make([]*workload_models.V1MatchExpression, len(selectors))

	for i, selector := range selectors {
		converted[i] = &workload_models.V1MatchExpression{
			Key:      selector.Key,
			Operator: selector.Operator,
			Values:   selector.Values,
		}
	}

	return converted
}

func convertWorkloadToIPAMStringMapEntry(mapEntries workload_models.V1StringMapEntry) ipam_models.NetworkStringMapEntry {
	converted := make(ipam_models.NetworkStringMapEntry, len(mapEntries))

	for k, v := range mapEntries {
		converted[k] = v
	}

	return converted
}

func convertIPAMToWorkloadStringMapEntry(mapEntries ipam_models.NetworkStringMapEntry) workload_models.V1StringMapEntry {
	converted := make(workload_models.V1StringMapEntry, len(mapEntries))

	for k, v := range mapEntries {
		converted[k] = v
	}

	return converted
}

// flattenComputeMatchExpressionsOrdered flattens the provided workload match expressions
// with respect to the order of any existing match expressions defined in the provided
// ResourceData. The prefix should be the flattened key of the list of match expressions
// in the ResourceData.
func flattenComputeMatchExpressionsOrdered(prefix string, data *schema.ResourceData, selectors []*workload_models.V1MatchExpression) []interface{} {
	ordered := make(map[string]int, len(selectors))
	for i, selector := range selectors {
		ordered[selector.Key] = i
	}
	flattened := make([]interface{}, len(selectors))
	for _, selector := range selectors {
		data := map[string]interface{}{
			"key":      selector.Key,
			"operator": selector.Operator,
			"values":   flattenStringArray(selector.Values),
		}
		flattened[ordered[selector.Key]] = data
	}
	return flattened
}

// flattenComputeMatchExpressions flattens the provided workload match expressions
// as given with no respect to ordering. If the order of the resulting match expressions
// is important, eg when using for diff logic, then flattenComputeMatchExpressionsOrdered
// should be used.
func flattenComputeMatchExpressions(selectors []*workload_models.V1MatchExpression) []interface{} {
	flattened := make([]interface{}, len(selectors))
	for i, selector := range selectors {
		flattened[i] = map[string]interface{}{
			"key":      selector.Key,
			"operator": selector.Operator,
			"values":   flattenStringArray(selector.Values),
		}
	}
	return flattened
}

func flattenStringMap(stringMap workload_models.V1StringMapEntry) map[string]interface{} {
	m := make(map[string]interface{}, len(stringMap))
	for k, v := range stringMap {
		m[k] = v
	}
	return m
}

func flattenStringArray(arr []string) []interface{} {
	a := make([]interface{}, len(arr))
	for i, s := range arr {
		a[i] = s
	}
	return a
}
