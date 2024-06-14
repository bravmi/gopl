package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bravmi/gopl/chap7/eval"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter an expression: ")
	scanner.Scan()
	exprStr := scanner.Text()

	fmt.Println("Parsing expression...")
	expr, err := eval.Parse(exprStr)
	if err != nil {
		fmt.Println("Failed to parse expression:", err)
		return
	}

	fmt.Println("Enter values for variables (e.g. x=1, y=2, ...):")
	env := eval.Env{}
	scanner.Scan()
	envStr := scanner.Text()
	for _, s := range strings.Fields(envStr) {
		pair := strings.Split(s, "=")
		if len(pair) != 2 {
			fmt.Println("Invalid variable assignment:", s)
			return
		}
		ident := pair[0]
		val, err := strconv.ParseFloat(pair[1], 64)
		if err != nil {
			fmt.Println("Failed to parse value:", err)
			return
		}
		env[eval.Var(ident)] = val
	}

	fmt.Println("Evaluating expression...")
	result := expr.Eval(env)
	fmt.Printf("%s = %.6g\n", expr, result)
}
