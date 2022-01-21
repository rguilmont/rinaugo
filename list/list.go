package list

// FoldLeft traverse the list from the first element to the last, and apply f transformation to type B to each element. Then in f you can make
//  whatever operation to the accunulator B.
func FoldLeft[A, B any](list []A) func(B, func(B, A) B) B {
	if list == nil {
		return nil
	}
	return func(init B, transform func(B, A) B) B {
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
