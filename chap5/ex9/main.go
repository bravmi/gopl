package main

import (
	"fmt"
	"regexp"
	"strings"
)

var shellPat = regexp.MustCompile(`\$[a-zA-Z_][a-zA-Z0-9_]*`)

func expand(s string, f func(string) string) string {
	wrapper := func(s string) string {
		return f(s[1:])
	}
	return shellPat.ReplaceAllStringFunc(s, wrapper)
}

func main() {
	s := "Hello, $foo. Do you know $bar?"
	fmt.Printf("expanded: %s\n", expand(s, strings.ToUpper))
}
