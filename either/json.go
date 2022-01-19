package either

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Wrapper to allow parsing a json value and convert it to an option
// ( An option being an interface, it can't implement JSON Unmarshaller interface... )
type EitherJson[A, B any] struct {
	right  A
	left   B
	isLeft bool
}

func (o *EitherJson[A, B]) UnmarshalJSON(b []byte) error {
	left := new(B)
	right := new(A)

	// Umpack in left first
	err1 := json.Unmarshal(b, right)
	if err1 != nil {
		// If impossible, try to unpack in right
		err2 := json.Unmarshal(b, left)
		// If it fails too, throw the error
		if err2 != nil {
			return fmt.Errorf("could not unmarshal in %T or %T : %v %v", reflect.TypeOf(left), reflect.TypeOf(right), err1, err2)
		}
		o.left = *left
		o.isLeft = true
		return nil
	}
	o.right = *right
	o.isLeft = false
	return nil
}

func (o EitherJson[A, B]) String() string {
	if o.isLeft {
		return fmt.Sprintf("%v", o.left)
	}
	return fmt.Sprintf("%v", o.right)
}

func (o EitherJson[A, B]) Get() Either[A, B] {
	if o.isLeft {
		return Left[A, B](o.left)
	}
	return Right[A, B](o.right)
}

// The marshalling must be done by the Left and Right struct
func (l left[A, B]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.Value)
}

func (r right[A, B]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Value)
}

// Just in case, marshaling OptionalJson should simply call the underlying method of Marshal
func (e EitherJson[A, B]) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Get())
}
