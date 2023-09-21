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

func TestMergeNestedMaps(t *testing.T) {
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
				base:  map[string]interface{}{"key1": "val1", "key2": "val2", "key3": map[string]interface{}{"subkey3": "subval3", "subkey4": "subval4"}},
				input: map[string]interface{}{"key0": "val0", "key2": "val2", "key3": map[string]interface{}{"subkey3": "subval3updated"}},
			},
			want: map[string]interface{}{"key0": "val0", "key1": "val1", "key2": "val2", "key3": map[string]interface{}{"subkey3": "subval3updated", "subkey4": "subval4"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeNestedMaps(tt.args.base, tt.args.input); !reflect.DeepEqual(got, tt.want) {
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

func TestToBool(t *testing.T) {
	tests := []struct {
		name     string
		list     map[string]interface{}
		key      string
		expected *bool
	}{
		{
			name:     "BoolValue",
			list:     map[string]interface{}{"key1": true},
			key:      "key1",
			expected: boolPtr(true),
		},
		{
			name:     "IntValue1",
			list:     map[string]interface{}{"key2": 1},
			key:      "key2",
			expected: boolPtr(true),
		},
		{
			name:     "IntValue0",
			list:     map[string]interface{}{"key3": 0},
			key:      "key3",
			expected: boolPtr(false),
		},
		{
			name:     "StringValue1",
			list:     map[string]interface{}{"key4": "1"},
			key:      "key4",
			expected: boolPtr(true),
		},
		{
			name:     "StringValue0",
			list:     map[string]interface{}{"key5": "0"},
			key:      "key5",
			expected: boolPtr(false),
		},
		{
			name:     "FloatValue1",
			list:     map[string]interface{}{"key6": 1.0},
			key:      "key6",
			expected: boolPtr(true),
		},
		{
			name:     "FloatValue0",
			list:     map[string]interface{}{"key7": 0.0},
			key:      "key7",
			expected: boolPtr(false),
		},
		{
			name:     "KeyNotExists",
			list:     map[string]interface{}{"key8": 42},
			key:      "nonexistent",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToBool(tt.list, tt.key)
			if (result == nil && tt.expected != nil) || (result != nil && *result != *tt.expected) {
				t.Errorf("toBool() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func boolPtr(b bool) *bool {
	return &b
}
