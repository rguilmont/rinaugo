package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/rguilmont/rinaugo/either"
	"github.com/rguilmont/rinaugo/function"
)

func readCSV(r *os.File) either.Either[[][]string, error] {
	csvReader := csv.NewReader(r)
	res := [][]string{}
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return either.Left[[][]string, error](err)
		}
		// do something with read line
		res = append(res, rec)
	}
	return either.Right[[][]string, error](res)
}

// Same as above, using more either function. As you can see it's not that different
//  as the above.
func readCSV2(r *os.File) function.Effect[[][]string, error] {
	return func() either.Either[[][]string, error] {
		csvReader := csv.NewReader(r)
		res := [][]string{}
		for {
			f := function.WrapFunc0(csvReader.Read)
			eith := f()
			if ok, err := eith.Left(); ok {
				if *err == io.EOF {
					break
				}
				return either.Left[[][]string, error](*err)
			}
			// do something with read line
			_, line := eith.Right()
			res = append(res, *line)
		}
		return either.Right[[][]string, error](res)
	}
}

func transform(s [][]string) []string {
	res := []string{}
	for _, line := range s {
		res = append(res, fmt.Sprintf("%v\n", line))
	}
	return res
}

func main() {
	ff1 := function.WrapFunc1(os.Open, "file.csv")
	ff2 := function.FlatMap(ff1, func(f *os.File) function.Effect[[][]string, error] {
		return func() either.Either[[][]string, error] {
			return readCSV(f)
		}
	})
	res := <-ff2.ParRun()

	if ok, left := res.Left(); ok {
		fmt.Println("Error occured during processing :", *left)

	} else {
		_, right := res.Right()
		fmt.Printf("%+v\n %T\n", *right, *right)
	}

	ff3 := function.WrapFunc1(os.Open, "file.csv")
	ff4 := function.FlatMap(ff3, readCSV2)
	ff5 := function.Map(ff4, transform)
	res2 := <-ff5.ParRun()

	if ok, left2 := res2.Left(); ok {
		fmt.Println("Error occured during processing :", *left2)
	} else {
		_, right2 := res2.Right()
		fmt.Printf("%+v\n %T\n", *right2, *right2)
	}

}
