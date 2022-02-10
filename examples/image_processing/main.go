package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"time"

	"github.com/corona10/goimagehash"
	"github.com/rguilmont/rinaugo/either"
	"github.com/rguilmont/rinaugo/function"
	"github.com/rguilmont/rinaugo/list"
	"github.com/sirupsen/logrus"
)

// Alias the types because it's a bit tedious
type imgPhash = either.Either[*goimagehash.ImageHash, error]
type imgDownloaded = either.Either[image.Image, error]

func getImage(url string) function.Effect[image.Image, error] {
	download := function.WrapFunc1(http.DefaultClient.Get, url)
	logrus.Infoln("Downloading", url)
	return function.FlatMap(download, func(r *http.Response) function.Effect[image.Image, error] {
		logrus.Infoln("Transforming", url)
		return function.WrapFunc1(jpeg.Decode, io.Reader(r.Body))
	})
}

func getPHash(img image.Image) function.Effect[*goimagehash.ImageHash, error] {
	logrus.Infoln("Calculating Phash for an image")
	return function.WrapFunc1(goimagehash.PerceptionHash, img)
}

// This shows how to process a list of image in parralel and in a safe way, thanks to ADT and Functions ParRun
func main() {

	toProcess := []string{}
	for i := 0; i < 1000; i++ {
		toProcess = append(toProcess, fmt.Sprintf("https://picsum.photos/id/%v/800/600", i))
	}

	partitioned := list.Split(toProcess, 1000)

	fmt.Println("Start scheduling", time.Now())
	res := list.Map(partitioned, func(urls []string) chan ([]imgPhash) {
		var processing function.EffectL[*goimagehash.ImageHash, error] = func() []imgPhash {
			res := []imgPhash{}
			for _, u := range urls {
				var processResult function.Effect[*goimagehash.ImageHash, error] = function.FlatMap(getImage(u), getPHash)
				res = append(res, processResult())
			}
			return res
		}
		return processing.ParRun()
	})
	fmt.Println("Ended scheduling", time.Now())

	list.ForEach(res, func(c chan ([]imgPhash)) {
		for result := range c {
			for _, elem := range result {
				if ok, val := elem.Right(); ok {
					fmt.Println("SUCCESS : ", **val)
				} else {
					fmt.Println("ERROR :", elem)
				}
			}
		}
	})

}
