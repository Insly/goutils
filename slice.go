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

// SliceContainsElement returns true if the slice contains element el
func SliceContainsElement[Element comparable](slice []Element, el Element) bool {
	for _, v := range slice {
		if v == el {
			return true
		}
	}

	return false
}

// SliceContainsElementFunc returns true if the comparisonFunc for an element returns true
func SliceContainsElementFunc[Element comparable](slice []Element, comparisonFunc func(el Element) bool) bool {
	for _, v := range slice {
		if comparisonFunc(v) {
			return true
		}
	}

	return false
}

// SliceFilterFunc returns new slice which pass the filterFunc
func SliceFilterFunc[Element comparable](slice []Element, filterFunc func(el Element) bool) []Element {
	filteredSlice := make([]Element, 0)
	for _, v := range slice {
		if filterFunc(v) {
			filteredSlice = append(filteredSlice, v)
		}
	}

	return filteredSlice
}
