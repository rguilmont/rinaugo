package options

import "fmt"

type nothing[A any] struct{}
type some[A any] struct {
	value A
}

func (n nothing[A]) a722fcf4_72e8_11ec_a2ee_3709e58e49d7() {}
func (s some[A]) a722fcf4_72e8_11ec_a2ee_3709e58e49d7()    {}

type Option[A any] interface {
	a722fcf4_72e8_11ec_a2ee_3709e58e49d7()
}

func NewOption[A any](content *A) Option[A] {
	if content == nil {
		return nothing[A]{}
	}
	return some[A]{*content}
}

func Some[A any](content A) Option[A] {
	return some[A]{content}
}

func Nothing[A any]() Option[A] {
	return nothing[A]{}
}

// Returns true if the option is nothing{}, false otherwise.
func IsEmpty[A any](o Option[A]) bool {
	fmt.Printf("%T %+v\n", o, o)
	if _, ok := o.(nothing[A]); ok {
		return true
	}
	return false
}

// Returns true if the option is an some, false otherwise.
func IsDefined[A any](o Option[A]) bool {
	return !IsEmpty(o)
}

// Get unwrap the Option, returning a pointer to a value, nil otherwise.
func Get[A any](o Option[A]) *A {
	if IsEmpty(o) {
		return nil
	}
	v := o.(some[A]).value
	return &v
}

// GetOrElse returns the value if the option is a some, def otherwise.
func GetOrElse[A any](o Option[A], def A) A {
	if IsEmpty(o) {
		return def
	}
	return o.(some[A]).value
}

func Map[A any, B any](o Option[A], f func(a A) B) Option[B] {
	if IsEmpty(o) {
		return nothing[B]{}
	}
	some := o.(some[A])
	converted := f(some.value)
	return NewOption(&converted)
}

func FlatMap[A any, B any](o Option[A], f func(a A) Option[B]) Option[B] {
	if IsEmpty(o) {
		return nothing[B]{}
	}
	some := o.(some[A])
	converted := f(some.value)
	return converted
}

type Zipped[A, B any] struct {
	a A
	b B
}

func Zip[A, B any](o Option[A], o2 Option[B]) Option[Zipped[A, B]] {
	return FlatMap(o, func(opt1 A) Option[B] {
		return Map(o2, func(opt2 B) Zipped[A, B] {
			return Zipped[A, B]{opt1, opt2}
		})
	})
}
