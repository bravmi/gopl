package main

import (
	"fmt"
	"unicode/utf8"
)

func rev(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverseUTF8(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	rev(b)
}

func main() {
	b := []byte("你好 世界")
	fmt.Println("b:")
	fmt.Println(string(b))
	reverseUTF8(b)
	fmt.Println(string(b))
}
