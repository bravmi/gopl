package main

import (
	"fmt"
	"crypto/sha256"
	"github.com/bravmi/gopl/chap2/popcount"
)

func byteDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popcount.PopCountClear(uint64(b[i]))
		case i >= len(b):
			count += popcount.PopCountClear(uint64(a[i]))
		default:
			count += popcount.PopCountClear(uint64(a[i] ^ b[i]))
		}
	}
	return count
}

func shaDiff(c1, c2 [32]byte) int {
	return byteDiff(c1[:], c2[:])
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n%d\n", c1, c2, c1 == c2, c1, shaDiff(c1, c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	// 125
}
