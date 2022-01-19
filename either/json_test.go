package either

import (
	"encoding/json"
	"testing"
)

type Test2 struct {
	V int
}

type test struct {
	A EitherJson[string, int]
	B EitherJson[string, EitherJson[int, float64]]
}

func TestJsonUnmarshilMustBeIdempotent(t *testing.T) {
	js := `{"A":"first value","B":11}`

	t.Run("Testing IsEmpty", func(t *testing.T) {
		value := test{}
		json.Unmarshal([]byte(js), &value)
		d2, _ := json.Marshal(value)

		t.Log(value.B.Get())

		t.Logf("Unmarshaled : %v", string(d2))
		if js != string(d2) {
			t.Errorf("Expected %v, got %v", js, string(d2))
		}
	})

}
