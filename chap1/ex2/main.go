package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for i, arg := range os.Args {
		fmt.Printf("i = %v, arg = %v\n", i, arg)
		s += arg + sep
	}
	fmt.Println(s)
}
