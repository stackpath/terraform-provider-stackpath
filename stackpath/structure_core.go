package stackpath

import "github.com/stackpath/terraform-provider-stackpath/stackpath/internal/models"

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
