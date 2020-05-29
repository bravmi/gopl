package ex1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKelvin(t *testing.T) {
	assert.Equal(t, CToK(AbsoluteZeroC), Kelvin(0))
	assert.Equal(t, KToC(Kelvin(0)), AbsoluteZeroC)
}
