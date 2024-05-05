package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingWriter(t *testing.T) {
	cw, pc := CountingWriter(os.Stdout)
	cw.Write([]byte("hello world"))
	assert.Equal(t, int64(11), *pc)
}
