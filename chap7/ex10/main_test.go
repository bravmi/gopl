package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	s := sort.StringSlice{"a", "b", "c", "b", "a"}
	assert.True(t, IsPalindrome(s))
}

func TestIsNotPalindrome(t *testing.T) {
	s := sort.IntSlice{1, 2, 3, 4, 5}
	assert.False(t, IsPalindrome(s))
}
