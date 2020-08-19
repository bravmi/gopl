package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(b []byte) []byte {
	space := false
	i := 0
	for j := 0; j < len(b); {
		r, size := utf8.DecodeRune(b[j:])
		if unicode.IsSpace(r) {
			j += size
			space = true
			continue
		}
		if space {
			// size = utf8.EncodeRune(b[i:], rune(' '))
			// i += size
			b[i] = byte(' ')
			i++
			space = false
		}
		utf8.EncodeRune(b[i:], r)
		i += size
		j += size
	}
	return b[:i]
}

func main() {
	b := []byte("hello\tworld")
	fmt.Println("squashSpaces:")
	fmt.Println(string(squashSpaces(b)))
}
