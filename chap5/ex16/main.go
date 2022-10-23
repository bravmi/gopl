package main

import (
	"fmt"
	"strings"
)

func join(sep string, elems ...string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}

	var b strings.Builder
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

func main() {
	fmt.Println(join(", ", "a", "b", "c"))
}
