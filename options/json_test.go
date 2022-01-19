package options

import (
	"encoding/json"
	"testing"
)

type Test2 struct {
	V int
}

type test struct {
	A OptionalJson[string]
	B OptionalJson[int]
	C OptionalJson[Test2]
	D OptionalJson[[]OptionalJson[string]]
}

func TestJsonUnmarshilMustBeIdempotent(t *testing.T) {
	js := `{"A":"first value","B":null,"C":{"V":22},"D":[null,null,null,"chien",null]}`

	t.Run("Testing IsEmpty", func(t *testing.T) {
		value := test{}
		json.Unmarshal([]byte(js), &value)
		d2, _ := json.Marshal(value)
		t.Logf("Unmarshaled : %v", string(d2))
		if js != string(d2) {
			t.Errorf("Expected %v, got %v", js, d2)
		}
	})

}
