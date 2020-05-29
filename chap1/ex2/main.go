package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for i, arg := range os.Args {
		fmt.Printf("i = %v, arg = %v", i, arg)
		fmt.Println()
		s += arg + sep
	}
	fmt.Println(s)
}
