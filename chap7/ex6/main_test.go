package main

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCelsiusFlag(t *testing.T) {
	os.Args = []string{"main", "-tempC", "30C"}
	flag.Parse()
	assert.Equal(t, Celsius(30.0), *tempC)
}

func TestCelsiusFlagF(t *testing.T) {
	os.Args = []string{"main", "-tempC", "86F"}
	flag.Parse()
	assert.InDelta(t, 30.0, float64(*tempC), 0.0001)
}

func TestCelsiusFlagK(t *testing.T) {
	os.Args = []string{"main", "-tempC", "303.15K"}
	flag.Parse()
	assert.InDelta(t, 30.0, float64(*tempC), 0.0001)
}
