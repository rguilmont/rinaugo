package list

// Split a `list` of `A` into multiple list of `A` with length of `size`. size must be > 0
func Split[A any](list []A, size uint) [][]A {
	if size == 0 {
		panic("can't split list with size 0")
	}
	fold := FoldLeft[A, [][]A](list)

	res := [][]A{}

	i := 0
	j := -1
	// This ensure the compiler that we do pass a positive value
	s := int(size)
	action := func(acc [][]A, step A) [][]A {

		if i%s == 0 {
			j += 1
			if len(list)-i < s {
				acc = append(acc, make([]A, len(list)-i))
			} else {
				acc = append(acc, make([]A, s))
			}
		}
		acc[j][i%s] = step
		i += 1
		return acc
	}

	return fold(res, action)
}
