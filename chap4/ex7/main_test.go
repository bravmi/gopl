package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse1(t *testing.T) {
	b := []byte("你好 世界")
	want := []byte("界世 好你")
	reverseUTF8(b)
	assert.Equal(t, want, b)
}

func TestReverse2(t *testing.T) {
	b := []byte("Räksmörgås")
	want := []byte("sågrömskäR")
	reverseUTF8(b)
	assert.Equal(t, want, b)
}


