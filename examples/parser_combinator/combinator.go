package parsercombinator

import (
	"fmt"
	"strings"

	"github.com/rguilmont/rinaugo/either"
)

// Draft of a parser combinator, using the ADTs declared in Rinaugo

type parserStep[A any] struct {
	parsed   A
	remaning string
}

type Tuple[A, B any] struct {
	a A
	b B
}

type none string

type parseResult[A any] either.Either[parserStep[A], error]

func parseResultSuccess[A any](v parserStep[A]) parseResult[A] {
	return either.Right[parserStep[A], error](v)
}

func parseResultError[A any](e error) parseResult[A] {
	return either.Left[parserStep[A], error](e)
}

type Parser[A any] func(toParse string) parseResult[A]

// Accept a string that start with expected
func String(expected string) Parser[string] {
	return func(s string) parseResult[string] {
		if strings.HasPrefix(s, expected) {
			res := parserStep[string]{
				parsed:   expected,
				remaning: s[len(expected):],
			}
			casted := either.Right[parserStep[string], error](res)
			return casted
		}
		return either.Left[parserStep[string], error](fmt.Errorf("expected %v, not found in %v", expected, s))
	}
}

func Token() Parser[string] {
	return func(s string) parseResult[string] {
		res := strings.Split(s, " ")
		if len(res) == 0 {
			return parseResultError[string](fmt.Errorf("Could not parse token %v", s))
		}
		step := parserStep[string]{
			parsed:   res[0],
			remaning: s[len(res[0]):],
		}

		return parseResultSuccess(step)
	}
}

func Number(expected int) Parser[int] {
	return func(s string) parseResult[int] {
		stringedNumber := fmt.Sprint(expected)
		if strings.HasPrefix(s, stringedNumber) {
			res := parserStep[int]{
				parsed:   expected,
				remaning: s[len(stringedNumber):],
			}
			casted := either.Right[parserStep[int], error](res)
			return casted
		}
		return either.Left[parserStep[int], error](fmt.Errorf("expected %v, not found in %v", expected, s))
	}
}

func Whitespace() Parser[string] {
	return String(" ")
}

func Concat[A, B any](p1 Parser[A], p2 Parser[B]) Parser[Tuple[A, B]] {
	return func(s string) parseResult[Tuple[A, B]] {
		res1 := p1(s)
		// This is a nightmare
		return either.FlatMap[parserStep[A], parserStep[Tuple[A, B]], error](res1, func(p parserStep[A]) either.Either[parserStep[Tuple[A, B]], error] {
			aa := p.parsed
			resultB := p2(p.remaning)

			mapped := either.Map[parserStep[B], parserStep[Tuple[A, B]], error](resultB, func(pp parserStep[B]) parserStep[Tuple[A, B]] {
				return parserStep[Tuple[A, B]]{
					parsed: Tuple[A, B]{
						a: aa,
						b: pp.parsed,
					},
					remaning: pp.remaning,
				}
			})
			return mapped
		})
	}
}

func Concat3[A, B, C any](p1 Parser[A], p2 Parser[B], p3 Parser[C]) Parser[Tuple[Tuple[A, B], C]] {
	c1 := Concat(p1, p2)
	c2 := Concat(c1, p3)
	return c2
}

func Concat4[A, B, C, D any](p1 Parser[A], p2 Parser[B], p3 Parser[C], p4 Parser[D]) Parser[Tuple[Tuple[Tuple[A, B], C], D]] {
	c1 := Concat3(p1, p2, p3)
	c2 := Concat(c1, p4)
	return c2
}

func Map[A, B any](p Parser[A], f func(A) B) Parser[B] {
	return func(s string) parseResult[B] {
		res := p(s)
		if ok, right := res.Right(); ok {
			return either.Right[parserStep[B], error](
				parserStep[B]{
					parsed:   f(right.parsed),
					remaning: right.remaning,
				})
		}
		_, left := res.Left()
		return either.Left[parserStep[B], error](*left)
	}
}

func Or[A, B any](p1 Parser[A], p2 Parser[B]) Parser[either.Either[A, B]] {
	return func(s string) parseResult[either.Either[A, B]] {
		res1 := p1(s)
		if ok, right := res1.Right(); ok {
			p := parseResultSuccess(parserStep[either.Either[A, B]]{
				parsed:   either.Right[A, B](right.parsed),
				remaning: right.remaning,
			})
			return p
		}
		_, err1 := res1.Left()

		res2 := p2(s)
		if ok, right := res2.Right(); ok {
			p := parseResultSuccess(parserStep[either.Either[A, B]]{
				parsed:   either.Left[A, B](right.parsed),
				remaning: right.remaning,
			})
			return p
		}
		_, err2 := res2.Left()

		return parseResultError[either.Either[A, B]](fmt.Errorf("Can't parse OR : got error (%v) and (%v)", *err1, *err2))
	}
}
