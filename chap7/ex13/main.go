package main

import (
	"fmt"

	"github.com/bravmi/gopl/chap7/eval"
)

func main() {
	e, _ := eval.Parse("x + y")
	fmt.Println(e.String())
}
