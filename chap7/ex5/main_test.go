package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitReader(t *testing.T) {
	r := LimitReader(strings.NewReader("Hello, World!"), 5)
	b, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, "Hello", string(b))
}
