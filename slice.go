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

func UniqueElements(slice []interface{}) []interface{} {
	s := make([]interface{}, 0, len(slice))
	exists := make(map[interface{}]bool)
	for _, value := range slice {
		if _, ok := exists[value]; !ok {
			exists[value] = true
			s = append(s, value)
		}
	}
	return s
}
