package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoReturn(t *testing.T) {
	assert.True(t, noReturn())
}
