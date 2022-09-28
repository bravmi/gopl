package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	s := "Hello, $foo. Do you know $bar?"
	expected := "Hello, FOO. Do you know BAR?"
	actual := expand(s, strings.ToUpper)
	assert.Equal(t, expected, actual)
}
