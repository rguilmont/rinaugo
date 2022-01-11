package options

import "fmt"

type Nothing[A any] struct{}
type Some[A any] struct {
	value A
}

func test() {

}

func (n Nothing[A]) a722fcf4_72e8_11ec_a2ee_3709e58e49d7() {}
func (s Some[A]) a722fcf4_72e8_11ec_a2ee_3709e58e49d7()    {}

type Option[A any] interface {
	a722fcf4_72e8_11ec_a2ee_3709e58e49d7()
}

func NewOption[A any](content *A) Option[A] {
	if content == nil {
		return Nothing[A]{}
	}
	return Some[A]{*content}
}

// Returns true if the option is Nothing{}, false otherwise.
func IsEmpty[A any](o Option[A]) bool {
	fmt.Printf("%T %+v\n", o, o)
	if _, ok := o.(Nothing[A]); ok {
		return true
	}
	return false
}

// Returns true if the option is an Some, false otherwise.
func IsDefined[A any](o Option[A]) bool {
	return !IsEmpty(o)
}

// Get unwrap the Option, returning a pointer to a value, nil otherwise.
func Get[A any](o Option[A]) *A {
	if IsEmpty(o) {
		return nil
	}
	v := o.(Some[A]).value
	return &v
}

// GetOrElse returns the value if the option is a Some, def otherwise.
func GetOrElse[A any](o Option[A], def A) A {
	if IsEmpty(o) {
		return def
	}
	return o.(Some[A]).value
}

func Map[A any, B any](o Option[A], f func(a A) B) Option[B] {
	if IsEmpty(o) {
		return Nothing[B]{}
	}
	some := o.(Some[A])
	converted := f(some.value)
	return NewOption(&converted)
}

func FlatMap[A any, B any](o Option[A], f func(a A) Option[B]) Option[B] {
	if IsEmpty(o) {
		return Nothing[B]{}
	}
	some := o.(Some[A])
	converted := f(some.value)
	return converted
}

type Zipped[A, B any] struct {
	a A
	b B
}

func Zip[A, B any](o Option[A], o2 Option[B]) Option[Zipped[A, B]] {
	if Get(o) == nil || Get(o2) == nil {
		return Nothing[Zipped[A, B]]{}
	}
	z := Zipped[A, B]{
		a: *Get(o),
		b: *Get(o2),
	}
	return NewOption(&z)
}
