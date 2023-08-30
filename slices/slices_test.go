package slices

import (
	"reflect"
	"testing"
)

func TestEqual(t *testing.T) {
	t.Parallel()
	type args[T comparable] struct {
		f []T
		s []T
	}
	type testCase[T comparable] struct {
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{args: struct {
			f []int
			s []int
		}{f: []int{1, 2}, s: []int{1, 2}}, want: true},
		{args: struct {
			f []int
			s []int
		}{f: []int{}, s: []int{}}, want: true},
		{args: struct {
			f []int
			s []int
		}{f: []int{1, 2}, s: []int{2, 1}}, want: false},
		{args: struct {
			f []int
			s []int
		}{f: []int{1, 1}, s: []int{1, 1, 1}}, want: false},
		{args: struct {
			f []int
			s []int
		}{f: []int{}, s: []int{0, 0, 0}}, want: false},
		{args: struct {
			f []int
			s []int
		}{f: []int{1, 2, 3, 4, 5, 6, 7, 8}, s: []int{1, 2, 3, 4, 5, 6, 7, 8}}, want: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Equal(tt.args.f, tt.args.s); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	t.Parallel()
	type args[T any] struct {
		size int
		f    func(i int) T
	}
	type testCase[T any] struct {
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{args: struct {
			size int
			f    func(i int) int
		}{size: 0, f: func(i int) int {
			return i
		}}, want: []int{}},

		{args: struct {
			size int
			f    func(i int) int
		}{size: 3, f: func(i int) int {
			return i
		}}, want: []int{0, 1, 2}},

		{args: struct {
			size int
			f    func(i int) int
		}{size: 4, f: func(i int) int {
			return i * 2
		}}, want: []int{0, 2, 4, 6}},

		{args: struct {
			size int
			f    func(i int) int
		}{size: 5, f: func(i int) int {
			return i
		}}, want: []int{0, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			if got := Generate(tt.args.size, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	t.Parallel()
	type testCase[T any] struct {
		args []T
		want []T
	}
	tests := []testCase[int]{
		{args: []int{}, want: []int{}},
		{args: []int{1}, want: []int{1}},
		{args: []int{0, 1, 2}, want: []int{2, 1, 0}},
		{args: []int{0, 1, 2, 3}, want: []int{3, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Reverse(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	t.Parallel()
	type args[T comparable] struct {
		s []T
		e T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{name: "same elements", args: struct {
			s []int
			e int
		}{s: []int{1, 2, 3}, e: 1}, want: 1},
		{name: "different elements", args: struct {
			s []int
			e int
		}{s: []int{1, 2, 3, 4, 5, 6, 7}, e: 6}, want: 1},
		{name: "different elements", args: struct {
			s []int
			e int
		}{s: []int{1, 2, 3, 4, 5, 6, 7, 6, 8, 9, 0, 6}, e: 6}, want: 3},
		{name: "no elements", args: struct {
			s []int
			e int
		}{s: []int{}, e: 10}, want: 0},
		{name: "different elements", args: struct {
			s []int
			e int
		}{s: []int{1, 2, 3, 4, 5, 9}, e: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	t.Parallel()
	type args[T comparable] struct {
		s []T
		e T
	}
	type testCase[T comparable] struct {
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{args: struct {
			s []int
			e int
		}{s: []int{1, 2, 3}, e: 1}, want: true},
		{args: struct {
			s []int
			e int
		}{s: []int{1, 2, 3, 4, 5, 6, 7}, e: 6}, want: true},
		{args: struct {
			s []int
			e int
		}{s: []int{1, 2, 3, 4, 5, 6, 7, 6, 8, 9, 0, 6}, e: 6}, want: true},
		{args: struct {
			s []int
			e int
		}{s: []int{}, e: 10}, want: false},
		{args: struct {
			s []int
			e int
		}{s: []int{1, 2, 3, 4, 5, 9}, e: 0}, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
