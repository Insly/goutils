package goutils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCopyMap(t *testing.T) {
	type args struct {
		input map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Positive case",
			args: args{input: map[string]interface{}{"key": "value"}},
			want: map[string]interface{}{"key": "value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyMap(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyMap() = %v, want %v", got, tt.want)
			} else {
				if !assert.NotSame(t, tt.args.input, got) {
					t.Error("CopyMap(): input and result have the same pointer")
				}
			}
		})
	}
}

func TestMergeMaps(t *testing.T) {
	type args struct {
		base  map[string]interface{}
		input map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Positive case",
			args: args{
				base:  map[string]interface{}{"key1": "val1", "key2": "val2", "key3": map[string]interface{}{"subkey3": "subval3"}},
				input: map[string]interface{}{"key0": "val0", "key2": "val2", "key3": map[string]interface{}{"subkey4": "subval4"}},
			},
			want: map[string]interface{}{"key0": "val0", "key1": "val1", "key2": "val2", "key3": map[string]interface{}{"subkey3": "subval3", "subkey4": "subval4"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeMaps(tt.args.base, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
