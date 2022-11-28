package goutils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
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

func TestSliceHasElement(t *testing.T) {
	t.Run("Test nil slice has no element", func(t *testing.T) {
		assert.False(t, SliceContainsElement(nil, "none"))
	})

	t.Run("Test string slice has element", func(t *testing.T) {
		slice := []string{"foo", "bar", "baz"}
		assert.True(t, SliceContainsElement(slice, "bar"))
		assert.True(t, SliceContainsElement(slice, "foo"))
		assert.False(t, SliceContainsElement(slice, "none"))
		assert.False(t, SliceContainsElement(nil, "none"))
	})

	t.Run("Test int slice has element", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		assert.True(t, SliceContainsElement(slice, 1))
		assert.True(t, SliceContainsElement(slice, 2))
		assert.False(t, SliceContainsElement(slice, 99))
	})

	t.Run("Test custom slice has element", func(t *testing.T) {
		slice := []struct {
			Name string
		}{
			{
				Name: "foo",
			},
			{
				Name: "bar",
			},
		}
		assert.True(t, SliceContainsElement(slice, struct{ Name string }{Name: "foo"}))
		assert.True(t, SliceContainsElement(slice, struct{ Name string }{Name: "bar"}))
		assert.False(t, SliceContainsElement(slice, struct{ Name string }{Name: "baz"}))
	})
}

func TestSliceHasElementFunc(t *testing.T) {
	t.Run("Test nil slice has no element with func", func(t *testing.T) {
		var nilSlice []string
		assert.False(t, SliceContainsElementFunc(nilSlice, func(idx int, el string) bool {
			return false
		}))
	})

	t.Run("Test string slice has element with func", func(t *testing.T) {
		slice := []string{"first", "second", "third"}

		// Slice has element which contains "ir"
		assert.True(t, SliceContainsElementFunc(slice, func(idx int, el string) bool {
			return strings.Contains(el, "ir")
		}))

		// Slice is missing element which contains "foo"
		assert.False(t, SliceContainsElementFunc(slice, func(idx int, el string) bool {
			return strings.Contains(el, "foo")
		}))
	})

	t.Run("Test custom slice has element with func", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		slice := []Person{
			{
				Name: "Foo",
				Age:  10,
			},
			{
				Name: "Foo",
				Age:  99,
			},
		}

		// Slice has person with name "Foo" and age "99"
		assert.True(t, SliceContainsElementFunc(slice, func(idx int, el Person) bool {
			return el.Name == "Foo" && el.Age == 99
		}))

	})

	t.Run("Test pointer slice has element with func", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		slice := []*Person{
			nil,
			nil,
			nil,
			nil,
			{
				Name: "Foo",
				Age:  10,
			},
		}

		// Slice has person with name "Foo" and age "99"
		assert.True(t, SliceContainsElementFunc(slice, func(idx int, el *Person) bool {
			return el != nil && el.Name == "Foo"
		}))
	})
}

func TestSliceFilterFunc(t *testing.T) {
	t.Run("Test filtering string slice with func", func(t *testing.T) {
		slice := []string{"A-string", "B-string", "C-string"}

		// Filter out everything not starting with "B"
		newSlice := SliceFilterFunc(slice, func(el string) bool {
			return strings.HasPrefix(el, "B")
		})
		if assert.Len(t, newSlice, 1) {
			assert.Equal(t, "B-string", newSlice[0])
		}
	})

	t.Run("Test filtering custom structs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		slice := []Person{
			{
				Name: "Foo",
				Age:  10,
			},
			{
				Name: "Foo",
				Age:  99,
			},
		}
		// Filter out all persons under age 51
		newSlice := SliceFilterFunc(slice, func(el Person) bool {
			return el.Age > 50
		})
		if assert.Len(t, newSlice, 1) {
			assert.Equal(t, 99, newSlice[0].Age)
		}
	})
}
