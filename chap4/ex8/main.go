// usage: go run main.go -top 10 < main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func mostCommon(counts map[rune]int, n int) []rune {
	keys := []rune{}
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return counts[keys[i]] > counts[keys[j]]
	})
	return keys[:n]
}

func main() {
	top := flag.Int("top", 20, "")
	flag.Parse()

	counts := map[rune]int{}        // counts of Unicode characters
	categories := map[string]int{}  // counts of Unicode categories
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for cat, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				categories[cat]++
			}
		}
		counts[r]++
		utflen[n]++
	}

	// runes
	fmt.Printf("%-5s\tcount\n", "rune")
	for _, r := range mostCommon(counts, *top) {
		n := counts[r]
		fmt.Printf("%-5q\t%d\n", r, n)
	}
	// lengths
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	// categories
	fmt.Printf("\n%-20s\tcount\n", "category")
	catNames := []string{}
	for cat := range categories {
		catNames = append(catNames, cat)
	}
	sort.Slice(catNames, func(i, j int) bool {
		return catNames[i] <= catNames[j]
	})
	for _, cat := range catNames {
		n := categories[cat]
		fmt.Printf("%-20s\t%d\n", cat, n)
	}
	// invalids
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
