package goutils

import (
	"reflect"
	"testing"
)

type dummyStruct struct {
	Field string                 `json:"field"`
	Slice []interface{}          `json:"slice"`
	Map   map[string]interface{} `json:"map"`
}

func TestStructToMap(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[string]interface{}
		wantErr    bool
	}{
		{
			name: "Positive case",
			args: args{input: &dummyStruct{
				Field: "string field",
				Slice: []interface{}{"item1", "item2"},
				Map:   map[string]interface{}{"key1": "value1", "key2": "value2"},
			}},
			wantResult: map[string]interface{}{
				"field": "string field",
				"slice": []interface{}{"item1", "item2"},
				"map":   map[string]interface{}{"key1": "value1", "key2": "value2"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := StructToMap(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("StructToMap() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
