package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCounter(t *testing.T) {
	var c WordCounter
	c.Write([]byte("hello world"))
	assert.Equal(t, WordCounter(2), c)
}

func TestLineCounter(t *testing.T) {
	var c LineCounter
	c.Write([]byte("hello\nworld"))
	assert.Equal(t, LineCounter(2), c)
}
