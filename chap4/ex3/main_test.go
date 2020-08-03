package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	a := &[]int{1, 2, 3}
	reverse(a)
	want := []int{3, 2, 1}
	assert.Equal(t, want, *a)
}
