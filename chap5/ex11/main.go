// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
// Usage:
// go run main.go
package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
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

//!-table

// !+main
func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) (order []string, err error) {
	seen := make(map[string]bool)
	var visitAll func(items []string, history map[string]int) error

	visitAll = func(items []string, history map[string]int) error {
		for _, item := range items {
			if seen[item] {
				start, ok := history[item]
				if !ok {
					break
				}
				cycle := []string{}
				for s, i := range history {
					if i >= start {
						cycle = append(cycle, s)
					}
				}
				cycle = append(cycle, item)
				return fmt.Errorf("cycle: %s", strings.Join(cycle, " -> "))
			} else {
				seen[item] = true
				history[item] = len(history)
				err := visitAll(m[item], history)
				if err != nil {
					return err
				}
				order = append(order, item)
				delete(history, item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	history := make(map[string]int)
	err = visitAll(keys, history)
	return order, err
}

//!-main
