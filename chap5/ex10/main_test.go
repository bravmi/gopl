package main

import (
	"testing"
)

func TestTopoSort(t *testing.T) {
	order := topoSort(prereqs)
	seen := make(map[string]bool)
	for _, item := range order {
		seen[item] = true
		for _, prereq := range prereqs[item] {
			if !seen[prereq] {
				t.Errorf("item %q appears before its prerequisite %q", item, prereq)
			}
		}
	}
}
