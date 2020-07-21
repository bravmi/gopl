package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	sha := flag.String("sha", "256", "which sha: 256 / 384 / 512")
	flag.Parse()
	for _, s := range flag.Args() {
		b := []byte{}
		switch *sha {
		case "384":
			c := sha512.Sum384([]byte(s))
			b = c[:]
		case "512":
			c := sha512.Sum512([]byte(s))
			b = c[:]
		case "256":
			fallthrough
		default:
			c := sha256.Sum256([]byte(s))
			b = c[:]
		}
		fmt.Printf("%x\n", b)
	}
}
