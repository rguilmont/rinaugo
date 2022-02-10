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

func TestForEach(t *testing.T) {
	have := []string{"hello", "there", "my", "name", "is", "JSB"}
	accumulator := []string{}

	want := []string{"HELLO", "THERE", "MY", "NAME", "IS", "JSB"}

	ForEach(have, func(e string) {
		accumulator = append(accumulator, strings.ToUpper(e))
	})
	if !reflect.DeepEqual(accumulator, want) {
		t.Errorf("Expected %v, got %v", want, accumulator)
	}
}

func TestNilSafety(t *testing.T) {
	var have []string

	got := FlatMap(have, func(e string) []string {
		return Map([]byte(e), func(s byte) string { return strings.ToUpper(string(s)) })
	})
	if got != nil {
		t.Errorf("Expected %v, got %v", "<nil>", got)
	}
	got = Map(have, strings.ToUpper)
	if got != nil {
		t.Errorf("Expected %v, got %v", "<nil>", got)
	}
	got2 := FoldLeft[string, int](have)(0, func(acc int, s string) int {
		acc += len(s)
		return acc
	})
	if got2 != 0 {
		t.Errorf("Expected %v, got %v", "<nil>", got)
	}
}

func TestFlatten(t *testing.T) {
	have := [][]int{{9, 8, 7, 6, 5, 1, 2, 3, 4, 5}, {100, 101, 102, 103}, {3000}, {}}
	want := []int{9, 8, 7, 6, 5, 1, 2, 3, 4, 5, 100, 101, 102, 103, 3000}
	got := Flatten(have)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", have, got)
	}
}

func TestSort(t *testing.T) {
	have := []int{9, 8, 7, 6, 5, 1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	got := Sort(have, func(a, b int) bool { return a < b })

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", have, got)
	}
}

func TestSplit(t *testing.T) {
	have := []int{9, 8, 7, 6, 5, 1, 2, 3, 4, 5}
	want := [][]int{{9, 8, 7}, {6, 5, 1}, {2, 3, 4}, {5}}

	got := Split(have, 3)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", have, got)
	}
}

func TestSplitWithNil(t *testing.T) {
	var have []int = nil
	var want [][]int = nil

	got := Split(have, 3)

	if len(got) != 0 {
		t.Errorf("Expected %T, got %T", want, got)
	}
}

func TestSplit2(t *testing.T) {
	have := []int{9, 8, 7, 6}
	want := [][]int{{9, 8, 7, 6}}

	got := Split(have, 30)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", have, got)
	}
}

func TestSplit3(t *testing.T) {
	have := []int{9, 8, 7, 6}
	want := [][]int{{9, 8, 7, 6}}

	got := Split(have, 0)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", have, got)
	}
}
