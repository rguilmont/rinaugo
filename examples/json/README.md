# JSON serialization

This little example shows how we can use options to add nil safety with json Marshalling/Unmarshalling.

Run it with `go run` as long as you have go 1.18.

The idea behind is to _enforce_ that a field exist by the compiler before a developer can play with it. It happened that we forgot to check that a field was not nil before trying to use its value, which leads ofc to Nil pointer exception :)