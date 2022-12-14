package goutils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
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

func TestGetMapNestedValue(t *testing.T) {
	type args struct {
		json      map[string]interface{}
		nestedKey []string
	}
	tests := []struct {
		name string
		args args
		want interface{}
		ok   bool
	}{
		{
			name: "Positive case",
			args: args{
				json:      map[string]interface{}{"key1": "val1", "key2": map[string]interface{}{"subkey2": "subval2"}},
				nestedKey: []string{"key2", "subkey2"},
			},
			want: "subval2",
			ok:   true,
		},
		{
			name: "Negative case",
			args: args{
				json:      map[string]interface{}{"key1": "val1", "key2": map[string]interface{}{"subkey2": "subval2"}},
				nestedKey: []string{"key2", "subkey3"},
			},
			want: nil,
			ok:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r, ok := GetMapNestedValue(tt.args.json, tt.args.nestedKey); (ok != tt.ok) || !reflect.DeepEqual(r, tt.want) {
				t.Errorf("TestGetNestedValue() = %v, %t, want %v, %t", r, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestMapToSlice(t *testing.T) {
	testMap := map[string]string{
		"one":   "two",
		"three": "four",
	}

	slice := MapToSlice(testMap)
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})

	assert.Equal(t, []string{"two", "four"}, slice)
}
