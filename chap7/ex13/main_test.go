package main

import (
	"testing"

	"github.com/bravmi/gopl/chap7/eval"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	var tests = []struct {
		expr string
		want string
	}{
		{"x", "x"},
		{"3.141", "3.141"},
		{"-x", "-x"},
		{"x + y", "x + y"},
		{"pow(x, y)", "pow(x, y)"},
	}
	for _, test := range tests {
		testname := test.expr
		t.Run(testname, func(t *testing.T) {
			expr, err := eval.Parse(test.expr)
			assert.NoError(t, err)
			got := expr.String()
			assert.Equal(t, test.want, got)
		})
	}
}
