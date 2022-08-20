# go-maths

[![PkgGoDev](https://pkg.go.dev/badge/github.com/jabolopes/go-maths)](https://pkg.go.dev/github.com/jabolopes/go-maths)

This package provides a generic library for math.

## Installation

```sh
$ go get github.com/jabolopes/go-maths
```

You can use `go get -u` to update the package. If you are using Go modules, you
can also just import the package and it will be automatically downloaded on the
first compilation.

## Examples

```go
math.Abs(2)
math.Abs(3.0)
math.Abs(uint32(4))

math.Neg(2)
math.Neg(3.0)
math.Neg(uint32(4))
```
