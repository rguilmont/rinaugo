package list

import (
	"reflect"
	"strings"
	"testing"
)

func TestFoldLeftTest(t *testing.T) {
	have := []string{"hello", "there", "my", "name", "is", "JSB"}
	want := func() int {
		init := 0
		for _, e := range have {
			init += len(e)
		}
		return init
	}()

	got := FoldLeft[string, int](have)(0, func(acc int, s string) int {
		acc += len(s)
		return acc
	})

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestMap(t *testing.T) {
	have := []string{"hello", "there", "my", "name", "is", "JSB"}
	want := []string{"HELLO", "THERE", "MY", "NAME", "IS", "JSB"}

	got := Map(have, strings.ToUpper)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestFlatMap(t *testing.T) {
	have := []string{"hello", "there", "my", "name", "is", "JSB"}

	want := []string{"H", "E", "L", "L", "O", "T", "H", "E", "R", "E", "M", "Y", "N", "A", "M", "E", "I", "S", "J", "S", "B"}

	got := FlatMap(have, func(e string) []string {
		return Map([]byte(e), func(s byte) string { return strings.ToUpper(string(s)) })
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestWithNil(t *testing.T) {
	var have []string

	got := FlatMap(have, func(e string) []string {
		return Map([]byte(e), func(s byte) string { return strings.ToUpper(string(s)) })
	})
	if got != nil {
		t.Errorf("Expected %v, got %v", "<nil>", got)
	}
}
