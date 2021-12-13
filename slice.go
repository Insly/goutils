package goutils

func ElementExists(slice []interface{}, el interface{}) bool {
	for _, v := range slice {
		if v == el {
			return true
		}
	}

	return false
}

func MergeSlices(base, input []interface{}) []interface{} {
	result := make([]interface{}, len(base))
	copy(result, base)
	for _, v := range input {
		if !ElementExists(result, v) {
			result = append(result, v)
		}
	}

	return result
}
