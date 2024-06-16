package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bravmi/gopl/chap7/eval"
)

func parseEnv(scanner *bufio.Scanner) eval.Env {
	env := eval.Env{}
	scanner.Scan()
	envStr := scanner.Text()
	for _, s := range strings.Fields(envStr) {
		s = strings.Trim(s, ",")
		parts := strings.Split(s, "=")
		if len(parts) != 2 {
			fmt.Println("invalid format:", s)
			continue
		}
		ident := parts[0]
		val, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Println("failed to parse value for", ident, ":", err)
			continue
		}
		env[eval.Var(ident)] = val
	}
	return env
}

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
	env := parseEnv(scanner)
	fmt.Println(env)

	fmt.Println("Evaluating expression...")
	result := expr.Eval(env)
	fmt.Printf("%s = %.6g\n", expr, result)
}
