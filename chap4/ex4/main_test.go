package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateLeft(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	rotateLeft(a, 2)
	want := []int{3, 4, 5, 1, 2}
	assert.Equal(t, want, a)
}

func TestRotateRight(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	rotateRight(a, 2)
	want := []int{4, 5, 1, 2, 3}
	assert.Equal(t, want, a)
}
