package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComma(t *testing.T) {
	assert.Equal(t, true, anagram("abc", "cba"))
	assert.Equal(t, false, anagram("abc", "cbac"))
	assert.Equal(t, true, anagram("", ""))
	assert.Equal(t, false, anagram("a", ""))
}
