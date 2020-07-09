package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComma(t *testing.T) {
	assert.Equal(t, "1", comma("1"))
	assert.Equal(t, "12", comma("12"))
	assert.Equal(t, "123", comma("123"))
	assert.Equal(t, "1,234", comma("1234"))
	assert.Equal(t, "1,234,567,890", comma("1234567890"))

	assert.Equal(t, "1.123", comma("1.123"))
	assert.Equal(t, "12.123", comma("12.123"))
	assert.Equal(t, "123.123", comma("123.123"))
	assert.Equal(t, "1,234.123", comma("1234.123"))
	assert.Equal(t, "1,234,567,890.123", comma("1234567890.123"))
	assert.Equal(t, "-1,234,567,890.123", comma("-1234567890.123"))
	assert.Equal(t, "+1,234,567,890.123", comma("+1234567890.123"))
}
