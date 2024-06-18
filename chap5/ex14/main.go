// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// usage:
// go run main.go
package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"linear algebra":        {"calculus"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// !+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		fmt.Printf("worklist: %q\n", worklist)
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				more := f(item)
				worklist = append(worklist, more...)
			}
		}
	}
}

//!-breadthFirst

// not a very good random distribution
func randomKey(m map[string][]string) string {
	for k := range m {
		return k
	}
	return ""
}

// !+main
func main() {
	worklist := []string{randomKey(prereqs)}
	deps := func(item string) []string {
		fmt.Printf("\t%q -> %q\n", item, prereqs[item])
		return prereqs[item]
	}
	breadthFirst(deps, worklist)
}

//!-main
