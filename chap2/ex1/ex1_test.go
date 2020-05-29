package ex1

import (
	"testing"

	"github.com/bravmi/gopl/chap2/tempconv"
	"github.com/stretchr/testify/assert"
)

func TestKelvin(t *testing.T) {
	assert.Equal(t, tempconv.CToK(tempconv.AbsoluteZeroC), tempconv.Kelvin(0))
	assert.Equal(t, tempconv.KToC(tempconv.Kelvin(0)), tempconv.AbsoluteZeroC)
}
