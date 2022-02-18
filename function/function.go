package function

import (
	"github.com/rguilmont/rinaugo/either"
)

// TODO : add doc and cleanup...

type Func0[A, B any] func() (A, B)
type Func1[A, B, C any] func(C) (A, B)
type Func2[A, B, C, D any] func(C, D) (A, B)
type Func3[A, B, C, D, E any] func(C, D, E) (A, B)
type Func4[A, B, C, D, E, F any] func(C, D, E, F) (A, B)
type Func5[A, B, C, D, E, F, G any] func(C, D, E, F, G) (A, B)
type Func6[A, B, C, D, E, F, G, H any] func(C, D, E, F, G, H) (A, B)
type Func7[A, B, C, D, E, F, G, H, I any] func(C, D, E, F, G, H, I) (A, B)
type Func8[A, B, C, D, E, F, G, H, I, J any] func(C, D, E, F, G, H, I, J) (A, B)
type Func9[A, B, C, D, E, F, G, H, I, J, K any] func(C, D, E, F, G, H, I, J, K) (A, B)
type Func10[A, B, C, D, E, F, G, H, I, J, K, L any] func(C, D, E, F, G, H, I, J, K, L) (A, B)
type Func11[A, B, C, D, E, F, G, H, I, J, K, L, M any] func(C, D, E, F, G, H, I, J, K, L, M) (A, B)

type Effect[A, B any] func() either.Either[A, B] // Don't have better name yet :(
type EffectL[A, B any] func() []either.Either[A, B]

// ParRun allows to run the function Effect in parallel, and returns a channel of either
func (e Effect[A, B]) ParRun() chan (either.Either[A, B]) {
	c := make(chan (either.Either[A, B]))
	go func() {
		c <- e()
		close(c)
	}()
	return c
}

// ParRun allows to run the function Effect in parallel, and returns a channel of either
func (e EffectL[A, B]) ParRun() chan ([]either.Either[A, B]) {
	c := make(chan ([]either.Either[A, B]))
	go func() {
		c <- e()
		close(c)
	}()
	return c
}

func Map[A, B, A1 any](e Effect[A, B], f func(a A) A1) Effect[A1, B] {
	return func() either.Either[A1, B] {
		res := either.Map(e(), f)
		return res
	}
}

func FlatMap[A, B, A1 any](e Effect[A, B], f func(a A) Effect[A1, B]) Effect[A1, B] {
	return func() either.Either[A1, B] {
		res := e()
		if ok, right := res.Right(); ok {
			return f(*right)()
		} else {
			_, left := res.Left()
			return either.Left[A1, B](*left)
		}
	}
}

func toEither[A, B any](a A, b B) either.Either[A, B] {
	if interface{}(b) != interface{}(nil) {
		return either.Left[A, B](b)
	}
	return either.Right[A, B](a)
}

func WrapFunc0[A, B any](f Func0[A, B]) Effect[A, B] {
	return func() either.Either[A, B] {
		return toEither(f())
	}
}

func WrapFunc1[A, B, C any](f Func1[A, B, C], arg0 C) Effect[A, B] {
	return func() either.Either[A, B] {
		return toEither(f(arg0))
	}
}
