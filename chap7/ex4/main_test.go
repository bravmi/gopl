package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReader(t *testing.T) {
	s := "hello world"
	b, err := io.ReadAll(NewReader(s))
	assert.NoError(t, err)
	assert.Equal(t, string(b), s, "b != s")
}

func TestNewReaderWithHTML(t *testing.T) {
	s := "<html><head></head><body><p>hello world</p></body></html>"
	b, err := io.ReadAll(NewReader(s))
	assert.NoError(t, err)
	assert.Equal(t, string(b), s, "b != s")
}
