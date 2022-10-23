package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 3, min(3))
	assert.Equal(t, 1, min(1, 2, 3, 4))
	vals := []int{1, 2, 3, 4}
	assert.Equal(t, 1, min(9, vals...))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 3, max(3))
	assert.Equal(t, 4, max(1, 2, 3, 4))
	vals := []int{1, 2, 3, 4}
	assert.Equal(t, 9, max(9, vals...))
}
