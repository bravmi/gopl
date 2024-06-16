// taken from:
// https://github.com/torbiak/gopl/blob/master/ex7.14/README

package main

import (
	"fmt"

	"github.com/bravmi/gopl/chap7/eval"
)

func main() {
	expr, _ := eval.Parse("x!")
	fmt.Println(expr.String())
	fmt.Println(expr.Eval(eval.Env{"x": 5}))
}
