package main

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShaDiff(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	assert.Equal(t, 125, shaDiff(c1, c2))
}

func TestByteDiff(t *testing.T) {
	assert.Equal(t, 2, byteDiff([]byte{0}, []byte{6}))
	assert.Equal(t, 7, byteDiff([]byte{1, 2, 3}, []byte{4, 5, 6}))
}
