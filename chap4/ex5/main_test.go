package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquash1(t *testing.T) {
	strs := []string{"a", "a", "b", "c", "c", "c", "d", "d", "e"}
	want := []string{"a", "b", "c", "d", "e"}
	assert.Equal(t, want, squash(strs))
}

func TestSquash2(t *testing.T) {
	strs := []string{"a", "a", "a"}
	want := []string{"a"}
	assert.Equal(t, want, squash(strs))
}

func TestSquash3(t *testing.T) {
	strs := []string{"a", "b", "c"}
	want := []string{"a", "b", "c"}
	assert.Equal(t, want, squash(strs))
}
