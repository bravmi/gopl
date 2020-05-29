package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, lineCount := range counts {
		n := 0
		for _, count := range lineCount {
			n += count
		}
		if n <= 1 {
			continue
		}
		fmt.Printf("%d\t%s\t[total]\n", n, line)
		for fname, count := range lineCount {
			fmt.Printf("\t%d\t[%s]\n", count, fname)
		}
	}
}

func countLines(f *os.File, fname string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if _, ok := counts[line]; !ok {
			counts[line] = make(map[string]int)
		}
		counts[line][fname]++
	}
}
