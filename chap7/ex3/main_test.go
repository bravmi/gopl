package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeString(t *testing.T) {
	root := &tree{value: 1}
	add(root, 2)
	add(root, 3)
	add(root, 4)
	add(root, 5)
	assert.Equal(t, "[1 2 3 4 5]", root.String())
}
