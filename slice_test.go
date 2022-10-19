package goutils

import (
	"reflect"
	"testing"
)

func TestElementExists(t *testing.T) {
	type args struct {
		slice []interface{}
		el    interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive case",
			args: args{
				slice: []interface{}{"1", "2", "3"},
				el:    "1",
			},
			want: true,
		},
		{
			name: "Negatice case",
			args: args{
				slice: []interface{}{"1", "2", "3"},
				el:    "0",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementExists(tt.args.slice, tt.args.el); got != tt.want {
				t.Errorf("ElementExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSlices(t *testing.T) {
	type args struct {
		base  []interface{}
		input []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "Positive case",
			args: args{
				base:  []interface{}{"1", "2", "3"},
				input: []interface{}{"0", "2"},
			},
			want: []interface{}{"1", "2", "3", "0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSlices(tt.args.base, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueElements(t *testing.T) {
	type args struct {
		slice []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "Positive case string",
			args: args{
				slice: []interface{}{"1", "2", "3", "2"},
			},
			want: []interface{}{"1", "2", "3"},
		},
		{
			name: "Positive case int",
			args: args{
				slice: []interface{}{1, 2, 3, 2},
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueElements(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
