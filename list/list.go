package list

import "sort"

// FoldLeft traverse the list from the first element to the last, and apply f transformation to type B to each element. Then in f you can make
//  whatever operation to the accunulator B.
func FoldLeft[A, B any](list []A) func(B, func(B, A) B) B {

	return func(init B, transform func(B, A) B) B {
		if list == nil {
			return init
		}
		step := init
		for _, elem := range list {
			step = transform(step, elem)
		}
		return step
	}
}

func Map[A, B any](list []A, f func(A) B) []B {
	if list == nil {
		return nil
	}
	res := make([]B, len(list))
	for i, elem := range list {
		res[i] = f(elem)
	}
	return res
}

func FlatMap[A, B any](list []A, f func(A) []B) []B {
	if list == nil {
		return nil
	}
	res := []B{}
	for _, elem := range list {
		res = append(res, f(elem)...)
	}
	return res
}

func ForEach[A any](list []A, f func(A)) {
	maped := func(a A) interface{} {
		f(a)
		return nil
	}
	Map(list, maped)
}

func Filter[A any](list []A, f func(A) bool) []A {
	if list == nil {
		return nil
	}
	res := []A{}
	for _, elem := range list {
		if f(elem) {
			res = append(res, elem)
		}
	}
	return res
}

func Flatten[A any](list [][]A) []A {
	if list == nil {
		return nil
	}
	size := 0
	for _, l := range list {
		size += len(l)
	}
	res := make([]A, size)
	i := 0
	for _, l := range list {
		for _, elem := range l {
			res[i] = elem
			i += 1
		}
	}
	return res
}

// Struct to implement https://pkg.go.dev/sort#Interface
type data[A any] struct {
	value []A
	less  func(a, b A) bool
}

func (d data[A]) Len() int {
	return len(d.value)
}

func (d data[A]) Less(i, j int) bool {
	return d.less(d.value[i], d.value[j])
}

func (d *data[A]) Swap(i, j int) {
	a := d.value[i]
	d.value[i] = d.value[j]
	d.value[j] = a
}

func Sort[A any](list []A, lessFunc func(a, b A) bool) []A {
	copied := make([]A, len(list))
	copy(copied, list)
	d := data[A]{
		value: copied,
		less:  lessFunc,
	}

	sort.Sort(&d)

	return d.value
}
