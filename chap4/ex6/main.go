package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(b []byte) []byte {
	space := false
	j := 0
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		i += size
		if unicode.IsSpace(r) {
			space = true
			continue
		}
		if space {
			b[j] = byte(' ')
			j++
			space = false
		}
		utf8.EncodeRune(b[j:], r)
		j += size
	}
	return b[:j]
}

func main() {
	b := []byte("hello\tworld")
	fmt.Println("squashSpaces:")
	fmt.Println(string(squashSpaces(b)))
}
