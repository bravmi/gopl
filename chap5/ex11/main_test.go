package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopoSortPrereqs(t *testing.T) {
	_, err := topoSort(prereqs)
	assert.ErrorContains(t, err, "cycle: calculus -> linear algebra -> calculus")
}

func TestTopoSortSimple(t *testing.T) {
	m := map[string][]string{"A": {"B"}, "B": {"C"}, "C": {"A"}}
	_, err := topoSort(m)
	assert.ErrorContains(t, err, "cycle")
}

func TestTopoSortTim(t *testing.T) {
	m := map[string][]string{"t": {"v", "w"}, "v": {"s"}, "w": {"s"}}
	order, _ := topoSort(m)
	assert.Equal(t, []string{"s", "v", "w", "t"}, order)
}

func TestTopoSortSelf(t *testing.T) {
	m := map[string][]string{"s": {"s"}}
	_, err := topoSort(m)
	assert.ErrorContains(t, err, "cycle")
}
