package stackpath

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/internal/models"
)

func convertComputeMatchExpression(data []interface{}) []*models.V1MatchExpression {
	selectors := make([]*models.V1MatchExpression, len(data))
	for i, s := range data {
		selector := s.(map[string]interface{})
		vals := make([]string, len(selector["values"].([]interface{})))
		for j, v := range selector["values"].([]interface{}) {
			vals[j] = v.(string)
		}
		selectors[i] = &models.V1MatchExpression{
			Key:      selector["key"].(string),
			Operator: selector["operator"].(string),
			Values:   vals,
		}
	}
	return selectors
}

func convertToStringMap(data map[string]interface{}) models.V1StringMapEntry {
	stringMap := make(models.V1StringMapEntry, len(data))
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

// flattenComputeMatchExpressions flattens the provided workload match expressions
// with respect to the order of any existing match expressions defined in the provided
// ResourceData. The prefix should be the flattened key of the list of match expressions
// in the ResourceData.
func flattenComputeMatchExpressionsOrdered(prefix string, data *schema.ResourceData, selectors []*models.V1MatchExpression) []interface{} {
	ordered := make(map[string]int, data.Get(prefix+".#").(int))
	for i, d := range selectors {
		ordered[d.Key] = i
	}
	s := make([]interface{}, data.Get(prefix+".#").(int))
	for _, v := range selectors {
		data := map[string]interface{}{
			"key":      v.Key,
			"operator": v.Operator,
			"values":   flattenStringArray(v.Values),
		}
		if index, exists := ordered[v.Key]; exists {
			s[index] = data
		} else {
			s = append(s, data)
		}
	}
	return s
}

// flattenComputeMatchExpressions flattens the provided workload match expressions
// as given with no respect to ordering. If the order of the resulting match expressions
// is important, eg when using for diff logic, then flattenComputeMatchExpressionsOrdered
// should be used.
func flattenComputeMatchExpressions(selectors []*models.V1MatchExpression) []interface{} {
	s := make([]interface{}, len(selectors))
	for i, v := range selectors {
		s[i] = map[string]interface{}{
			"key":      v.Key,
			"operator": v.Operator,
			"values":   flattenStringArray(v.Values),
		}
	}
	return s
}

func flattenStringMap(stringMap models.V1StringMapEntry) map[string]interface{} {
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
