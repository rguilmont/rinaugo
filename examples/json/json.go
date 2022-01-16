package main

import (
	"encoding/json"
	"fmt"

	"github.com/rguilmont/rinaugo/options"
)

func main() {

	type Test2 struct {
		V int
	}

	type test struct {
		A options.OptionalJson[string]
		B options.OptionalJson[int]
		C options.OptionalJson[Test2]
	}

	js := `
	{
		"A": "first value",
		"C": {
			"V": 22
		}
	}
	`
	t := test{}
	err := json.Unmarshal([]byte(js), &t)

	fmt.Println(err, t)

	t2 := struct {
		A options.Option[string]
		B options.Option[int]
		C options.Option[Test2]
	}{
		options.Some("22"),
		options.Some(22),
		options.Nothing[Test2](),
	}

	js2, err := json.Marshal(t2)
	fmt.Println(string(js2), err)
}
