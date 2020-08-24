// usage: go run main.go -top 10 < main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func mostCommon(counts map[string]int, n int) []string {
	keys := []string{}
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

	counts := map[string]int{}

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		word := in.Text()
		counts[word]++
	}
	if err := in.Err(); err != nil {
		log.Fatalf("wordfreq: %v\n", err)
	}

	// words
	fmt.Printf("%-15s\tcount\n", "word")
	for _, w := range mostCommon(counts, *top) {
		n := counts[w]
		fmt.Printf("%-15s\t%d\n", w, n)
	}
}
