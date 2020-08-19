package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquashSpaces1(t *testing.T) {
	b := []byte("hello\tworld")
	want := []byte("hello world")
	assert.Equal(t, want, squashSpaces(b))
}

func TestSquashSpaces2(t *testing.T) {
	b := []byte("abc\r  \n\rdef")
	want := []byte("abc def")
	assert.Equal(t, want, squashSpaces(b))
}
