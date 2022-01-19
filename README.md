# RinauGo

**This repository is for the moment, as the state of golang 1.18, experimental.**

RinauGo adds some ADT to Golang, and functional programming capabilities.

The goal is not to transform the language, which is very imperative-oriented, but rather to enjoy some nice FP functions when it can be useful and add safety.

It relies on Generics, so require go >= 1.18.

## Why RinauGo

Because of [Nicolas Rinaudo](https://github.com/nrinaudo), colleague and Scala/FP advocate at our company, that doesn't really like golang as a language. That was mainly before the generics, maybe... 😉

## Options

The main goal of Options is being able to represent to presence ( or absence ) or value. This is already possible in golang, using pointers, but there's no Nil safety ensured by the compiler, which can leads to nil pointer exception at runtime. This is never very pleasant :)

## Either

Either is a disjoint union.
A useful case of Either is to transform common functions in Go that returns 2 values ( usually pointer or something and an error )
to Wrap it. For instance :
```
 func() (*int, error)
```
can be transformed into
```
 func() Either[int, error].
```
Not only it open to solutions to avoid the infamous `if err != nil {...` pattern, but it also
 allows to map, and compose. It's a type used a lot by the functions package.
Either is right-biased, since it was inspired by Scala Either. So as stated in the scala code :
 `Right` is assumed to be the default case to operate on.
 If it is `Left`, operations like `map` and `flatMap` return the `Left` value unchanged.

Still WIP :)

## Function

This one allows to run "Effects" and compose functions. There are helper functions to transform a function that returns 2 variables ( usually a pointer and an error ) into an `Effect[A,B]`. You can combine Effects function and run even run them in the background with `ParRun`, which returns a `chan(Either[A,B])`.

 See `examples/csv` to see how it works and why it's useful.

 Still very WIP too :)
