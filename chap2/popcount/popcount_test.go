// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testNumber uint64 = 0x1234567890ABCDEF

func TestPopCountTable(t *testing.T) {
	assert.Equal(t, 32, PopCountTable(testNumber))
}

func TestPopCountTableLoop(t *testing.T) {
	assert.Equal(t, PopCountTable(testNumber), PopCountTableLoop(testNumber))
}

func TestPopCountShift(t *testing.T) {
	assert.Equal(t, PopCountTable(testNumber), PopCountShift(testNumber))
}

func TestPopCountClear(t *testing.T) {
	assert.Equal(t, PopCountTable(testNumber), PopCountClear(testNumber))
}

func BenchmarkPopCountTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(testNumber)
	}
}

func BenchmarkPopCountTableLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTableLoop(testNumber)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(testNumber)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(testNumber)
	}
}
