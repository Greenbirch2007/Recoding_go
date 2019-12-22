package main

import (
	"fmt"
	"go/ast"
)

type A struct {
	Name string
	id int64
}

type B struct {
	A
	Name string
	num int
}

type C struct {
	A
	B
}
func main()  {

	c := C{}
	c.A.Name = "xx"
	c.B.Name = "b"\
	c.
}