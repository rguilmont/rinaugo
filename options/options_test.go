package options

import (
	"fmt"
	"testing"
)

func pointer[A any](a A) *A {
	return &a
}

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		a    Option[int]
		want bool
	}{
		{NewOption(pointer(1)), false},
		{NewOption[int](nil), true},
	}

	for _, tt := range tests {
		t.Run("Testing IsEmpty", func(t *testing.T) {
			if !IsEmpty(tt.a) == tt.want {
				t.Errorf("got %v, want %v", tt.a, tt.want)
			}
		})
	}
}

func TestIsDefined(t *testing.T) {
	var tests = []struct {
		a Option[int]
	}{
		{NewOption(pointer(1))},
		{NewOption[int](nil)},
	}

	for _, tt := range tests {
		t.Run("Testing IsEmpty", func(t *testing.T) {
			if IsEmpty(tt.a) == IsDefined(tt.a) {
				t.Errorf("got %v, want %v", IsEmpty(tt.a), IsDefined(tt.a))
			}
		})
	}
}

func TestGet(t *testing.T) {
	var tests = []*string{
		pointer("chababadou"),
		nil,
	}

	for _, tt := range tests {
		t.Run("Testing IsEmpty", func(t *testing.T) {

			opt := NewOption(tt)
			v := Get(opt)
			if v == tt && !(v == nil && tt == nil) {
				t.Errorf("got %v, want %v", tt, Get(opt))
			}
		})
	}
}

func TestMap(t *testing.T) {
	var tests = []struct {
		a         Option[int]
		want      Option[string]
		transform func(int) string
	}{
		{
			NewOption(pointer(1)),
			NewOption(pointer("1")),
			func(i int) string { return fmt.Sprintf("%v", i) },
		},
	}

	for _, tt := range tests {
		t.Run("Testing IsEmpty", func(t *testing.T) {
			if Get(Map(tt.a, tt.transform)) == Get(tt.want) {
				t.Errorf("got %v, want %v", tt.a, tt.want)
			}
		})
	}
}

func TestZip(t *testing.T) {
	var tests = []struct {
		a    Option[int]
		b    Option[string]
		want Option[Zipped[int, string]]
	}{
		{
			NewOption(pointer(1)),
			NewOption(pointer("1")),
			NewOption(&Zipped[int, string]{1, "1"}),
		},
		{
			NewOption(pointer(1)),
			NewOption[string](nil),
			Nothing[Zipped[int, string]]{},
		},
	}

	for _, tt := range tests {
		t.Run("Testing IsEmpty", func(t *testing.T) {

			zipped := Zip(tt.a, tt.b)

			if zipped != tt.want {
				t.Errorf("got %v, want %v", zipped, tt.want)
			}
		})
	}
}
