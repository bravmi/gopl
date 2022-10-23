package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		sep   string
		elems []string
		want  string
	}{
		{"", []string{"a", "b", "c"}, "abc"},
		{", ", []string{"a", "b", "c"}, "a, b, c"},
		{"", []string{"a"}, "a"},
		{"", []string{}, ""},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, join(test.sep, test.elems...))
		assert.Equal(t, test.want, strings.Join(test.elems, test.sep))
	}
}
