package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	assert.Equal(t, 3, x.Len())
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Remove(1)
	x.Remove(9)
	assert.Equal(t, "{144}", x.String())
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Clear()
	assert.Equal(t, "{}", x.String())
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y := x.Copy()
	assert.Equal(t, x.String(), y.String())
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 144, 9)
	assert.Equal(t, "{1 9 144}", x.String())
}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)
	x.IntersectWith(&y)
	assert.Equal(t, "{9}", x.String())
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)
	x.DifferenceWith(&y)
	assert.Equal(t, "{1 144}", x.String())
}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)
	x.SymmetricDifference(&y)
	assert.Equal(t, "{1 42 144}", x.String())
}
