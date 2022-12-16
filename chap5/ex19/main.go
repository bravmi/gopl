package main

import "fmt"

func noReturn() (ret bool) {
	defer func() {
		if p := recover(); p != nil {
			ret = true
		}
	}()
	panic(true)
}

func main() {
	fmt.Printf("noReturn() = %v\n", noReturn())
}
