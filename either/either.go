package either

// Either represents a value of one of two possible types (a disjoint union).
// A useful case of Either is to transform common functions in Go that returns 2 values ( usually pointer or something and an error )
// to Wrap it. For instance :
//  func() (*int, error)
// can be transformed into
//  func() Either[int, error].
// Not only it open to solutions to avoid the infamous `if err != nil {...` pattern, but it also
//  allows to map, and compose. It's a type used a lot by the functions package.
// Either is right-biased, since it was inspired by Scala Either. So as stated in the scala code :
//  `Right` is assumed to be the default case to operate on.
//  If it is `Left`, operations like `map` and `flatMap` return the `Left` value unchanged.
type Either[A, B any] interface {
	// This function is a hidden one that simply allow to "control" what implement this interface.
	//  it's like
	a722fcf4_72e8_11ec_a2ee_3709e58e49d1()
	IsLeft() bool
	IsRight() bool
	// If you need to get the value directly, and check weither it's a left or a right, a good pattern is :
	// if ok, left := e.Left();ok {
	// 	 ...
	//} else {
	//  _, right := e.Right()
	//	...
	//}
	Left() (bool, *B)
	Right() (bool, *A)
}

type left[A, B any] struct {
	Value B
}

func (left[A, B]) IsLeft() bool        { return true }
func (left[A, B]) IsRight() bool       { return false }
func (l left[A, B]) Left() (bool, *B)  { return true, &l.Value }
func (l left[A, B]) Right() (bool, *A) { return false, nil }

type right[A, B any] struct {
	Value A
}

func (right[A, B]) IsLeft() bool        { return false }
func (right[A, B]) IsRight() bool       { return true }
func (r right[A, B]) Left() (bool, *B)  { return false, nil }
func (r right[A, B]) Right() (bool, *A) { return true, &r.Value }

func (l left[A, B]) a722fcf4_72e8_11ec_a2ee_3709e58e49d1()  {}
func (r right[A, B]) a722fcf4_72e8_11ec_a2ee_3709e58e49d1() {}

func Left[A, B any](l B) Either[A, B] {
	return left[A, B]{l}
}

func Right[A, B any](r A) Either[A, B] {
	return right[A, B]{r}
}

// Fold Applies fa if Either is a right and fb if Either is a left. Both returns a C.
func Fold[A, B, C any](e Either[A, B], fa func(A) C, fb func(B) C) C {
	if e.IsRight() {
		return fa(e.(right[A, B]).Value)
	}
	return fb(e.(left[A, B]).Value)
}

// FlatMap expect a func of A => Either[A1, B]. Function is ran when Either is a right.
func FlatMap[A, A1, B any](e Either[A, B], f func(A) Either[A1, B]) Either[A1, B] {

	// Happy path
	if e.IsRight() {
		v := e.(right[A, B]).Value
		return f(v)
	}
	// e is Left, return e
	return Left[A1, B](e.(left[A, B]).Value)
}

// Map expect a func of A => A1. Function is ran when Either is a right.
func Map[A, A1, B any](e Either[A, B], f func(A) A1) Either[A1, B] {
	// Happy path
	if e.IsRight() {
		v := e.(right[A, B]).Value
		return Right[A1, B](f(v))
	}
	// e is Left, return e
	return Left[A1, B](e.(left[A, B]).Value)
}

func Flatten[A, B any](e Either[Either[A, B], B]) Either[A, B] {
	if ok, left := e.Left(); ok {
		return Left[A, B](*left)
	} else {
		_, right := e.Right()
		return *right
	}
}
