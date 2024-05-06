package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingWriter(t *testing.T) {
	cw, pc := CountingWriter(os.Stdout)
	_, err := cw.Write([]byte("hello world"))
	assert.NoError(t, err)
	assert.Equal(t, int64(11), *pc)
}
