package options

import (
	"encoding/json"
	"fmt"
)

// Wrapper to allow parsing a json value and convert it to an option
// ( An option being an interface, it can't implement JSON Unmarshaller interface... )
type OptionalJson[A any] struct {
	value *A
}

func (o *OptionalJson[A]) UnmarshalJSON(b []byte) error {
	v := new(A)
	if b == nil || string(b) == "null" {
		return nil
	}
	err := json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	o.value = v
	return nil
}

func (o OptionalJson[A]) String() string {
	if o.value != nil {
		return fmt.Sprintf("%v", *o.value)
	}
	return "<nil>"
}

func (o OptionalJson[A]) Get() Option[A] {
	if o.value == nil {
		return Nothing[A]()
	}
	return Some(*o.value)
}

// The marshalling must be done by the Some and Nothing struct
func (n nothing[A]) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

func (s some[A]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

// Just in case, marshaling OptionalJson should simply call the underlying method of Marshal
func (o OptionalJson[A]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Get())
}
