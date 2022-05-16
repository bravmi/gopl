// usage: go run main.go abc cba abc cbac
package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	if len(os.Args)%2 != 1 {
		log.Fatal("uneven number of arguments")
	}
	for i := 1; i < len(os.Args)-1; i++ {
		fmt.Printf("  %v\n", anagram(os.Args[i], os.Args[i+1]))
	}
}

func counter(s string) map[rune]int {
	c := make(map[rune]int)
	for _, r := range s {
		c[r]++
	}
	return c
}

func anagram(s1, s2 string) bool {
	c1 := counter(s1)
	c2 := counter(s2)
	return reflect.DeepEqual(c1, c2)
}
