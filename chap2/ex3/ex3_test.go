// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package ex3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopCountTable(t *testing.T) {
	assert.Equal(t, 32, PopCountTable(0x1234567890ABCDEF))
}

func TestPopCountTableLoop(t *testing.T) {
	assert.Equal(t, PopCountTable(0x1234567890ABCDEF), PopCountTableLoop(0x1234567890ABCDEF))
}

func TestPopCountShift(t *testing.T) {
	assert.Equal(t, PopCountTable(0x1234567890ABCDEF), PopCountShift(0x1234567890ABCDEF))
}

func TestPopCountClear(t *testing.T) {
	assert.Equal(t, PopCountTable(0x1234567890ABCDEF), PopCountClear(0x1234567890ABCDEF))
}

func BenchmarkPopCountTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountTableLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTableLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(0x1234567890ABCDEF)
	}
}
