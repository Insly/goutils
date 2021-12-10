package goutils

import "encoding/json"

func StructToMap(input interface{}) (result map[string]interface{}, err error) {
	obj, err := json.Marshal(input)
	if err != nil {
		return
	}
	err = json.Unmarshal(obj, &result)
	return
}
