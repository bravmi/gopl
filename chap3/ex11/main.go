// usage: go run main.go 1 12 123 1234 1234567890
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	re := regexp.MustCompile(`([+-])?(\d+)([.]\d+)?`)
	m := re.FindStringSubmatch(s)
	if m == nil {
		log.Fatal("invalid string")
	}
	buf := bytes.Buffer{}

	if sign := m[1]; sign != "" {
		buf.WriteString(sign)
	}

	if s := m[2]; true {
		r := len(s) % 3
		if r == 0 {
			r = 3
		}
		buf.WriteString(s[:r])
		for i := r; i < len(s); i += 3 {
			buf.WriteByte(',')
			buf.WriteString(s[i : i+3])
		}
	}

	if s := m[3]; s != "" {
		buf.WriteString(s)
	}

	return buf.String()
}
