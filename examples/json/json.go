package main

import (
	"encoding/json"
	"fmt"

	"github.com/rguilmont/rinaugo/either"
	"github.com/rguilmont/rinaugo/options"
)

func main() {

	type errorMessage string

	type Test2 struct {
		V int
	}

	type test struct {
		A options.OptionalJson[string]
		B options.OptionalJson[int]
		C options.OptionalJson[Test2]
		D either.EitherJson[float64, errorMessage]
	}

	js := `
	{
		"A": "first value",
		"C": {
			"V": 22
		},
		"D": "Error message: it failed"
	}
	`
	t := test{}
	err := json.Unmarshal([]byte(js), &t)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed json : %+v\n", t)

	// Deal with the options
	a := options.Map(t.A.Get(), func(s string) int {
		return len(s)
	})

	d := either.Map(t.D.Get(), func(f float64) float64 {
		return f * 2
	})

	t2 := struct {
		A options.Option[int]
		B options.Option[int]
		C options.Option[Test2]
		D either.Either[float64, errorMessage]
	}{
		a,
		options.Some(22),
		options.Nothing[Test2](),
		d,
	}

	js2, err := json.Marshal(t2)
	fmt.Println(string(js2), "Error while marshalling: ", err)
}
