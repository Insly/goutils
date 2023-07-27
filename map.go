package goutils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"hash"
	"strconv"
)

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

func MergeNestedMaps(base, input map[string]interface{}) map[string]interface{} {
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
				result[k] = MergeNestedMaps(baseSubMap, sub)
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

func GetMapNestedValue(json map[string]interface{}, nestedKey []string) (val interface{}, ok bool) {
	count := len(nestedKey)
	if count == 0 || json == nil {
		return nil, false
	}

	key := nestedKey[0]

	val, ok = json[key]
	if !ok || (count == 1) {
		return val, ok
	}

	if subJson, ok := val.(map[string]interface{}); ok {
		return GetMapNestedValue(subJson, nestedKey[1:])
	} else {
		return nil, false
	}
}

func ToString(list map[string]interface{}, key string) *string {
	val, exists := list[key]
	if exists {
		if r, ok := val.(string); ok {
			return &r
		}
	}
	return nil
}

func ForceString(list map[string]interface{}, key string, defaultVal string) string {
	r := ToString(list, key)
	if r == nil {
		return defaultVal
	}
	return *r
}

func ToInt(list map[string]interface{}, key string) *int {
	val, exists := list[key]
	if exists {
		switch v := val.(type) {
		case int:
			return &v
		case float64:
			i := int(v)
			return &i
		case float32:
			i := int(v)
			return &i
		case int64:
			i := int(v)
			return &i
		case int32:
			i := int(v)
			return &i
		case string:
			if i, err := strconv.Atoi(v); err == nil {
				return &i
			}
		}
	}
	return nil
}

func ForceInt(list map[string]interface{}, key string, defaultVal int) int {
	r := ToInt(list, key)
	if r == nil {
		return defaultVal
	}
	return *r
}

func ToBool(list map[string]interface{}, key string) *bool {
	val, exists := list[key]
	if exists {
		if r, ok := val.(bool); ok {
			return &r
		}
	}
	return nil
}

func ForceBool(list map[string]interface{}, key string, defaultVal bool) bool {
	r := ToBool(list, key)
	if r == nil {
		return defaultVal
	}
	return *r
}

func ToFloat(list map[string]interface{}, key string) *float64 {
	val, exists := list[key]
	if exists {
		switch v := val.(type) {
		case float64:
			return &v
		case float32:
			i := float64(v)
			return &i
		case int64:
			i := float64(v)
			return &i
		case int32:
			i := float64(v)
			return &i
		case int:
			i := float64(v)
			return &i
		case string:
			if i, err := strconv.ParseFloat(v, 64); err == nil {
				return &i
			}
		}
	}
	return nil
}

func ForceFloat(list map[string]interface{}, key string, defaultVal float64) float64 {
	r := ToFloat(list, key)
	if r == nil {
		return defaultVal
	}
	return *r
}

// MapToSlice extracts map values into a slice
func MapToSlice[key comparable, value any](input map[key]value) []value {
	slice := make([]value, 0, len(input))
	for _, v := range input {
		slice = append(slice, v)
	}

	return slice
}

func UniqueHash[key comparable, value any](input map[key]value, cryptoAlgo string) string {
	v, err := json.Marshal(input)
	if err != nil {
		return ""
	}

	var h hash.Hash

	switch cryptoAlgo {
	case "sha1":
		h = sha1.New()
	case "sha", "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		h = md5.New()
	}
	h.Write(v)

	return fmt.Sprintf("%x", h.Sum(nil))
}
