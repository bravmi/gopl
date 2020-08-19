package main

import (
	"fmt"
	"unicode"
)

func squashSpaces(b []byte) []byte {
	out := make([]rune, 0)
	space := false
	for _, r := range string(b) {
		if unicode.IsSpace(r) {
			space = true
			continue
		}
		if space {
			out = append(out, ' ')
			space = false
		}
		out = append(out, r)
	}
	return []byte(string(out))
}

func main() {
	b := []byte("hello\tworld")
	fmt.Println("squashSpaces:")
	fmt.Println(string(squashSpaces(b)))
}
