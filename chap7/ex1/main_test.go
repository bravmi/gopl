package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCounter(t *testing.T) {
	var c WordCounter
	_, err := c.Write([]byte("hello world"))
	assert.NoError(t, err)
	assert.Equal(t, WordCounter(2), c)
}

func TestLineCounter(t *testing.T) {
	var c LineCounter
	_, err := c.Write([]byte("hello\nworld"))
	assert.NoError(t, err)
	assert.Equal(t, LineCounter(2), c)
}
