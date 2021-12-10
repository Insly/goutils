package goutils

func MergeMaps(base, input map[string]interface{}) map[string]interface{} {
	if base == nil {
		return CopyMap(input)
	}
	if input == nil {
		return CopyMap(base)
	}

	result := CopyMap(base)
	for k, v := range input {
		switch sub := v.(type) {
		case map[string]interface{}:
			if baseSubMap, isMap := result[k].(map[string]interface{}); isMap {
				result[k] = MergeMaps(baseSubMap, sub)
			} else {
				result[k] = v
			}
		case []interface{}:
			if baseSubSlice, isSlice := result[k].([]interface{}); isSlice {
				result[k] = MergeSlices(baseSubSlice, sub)
			} else {
				result[k] = v
			}
		default:
			result[k] = v
		}
	}

	return result
}

func CopyMap(input map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{}, len(input))
	for k, v := range input {
		sub, ok := v.(map[string]interface{})
		if ok {
			res[k] = CopyMap(sub)
		} else {
			res[k] = v
		}
	}

	return res
}
