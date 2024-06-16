package main

import (
	"fmt"
	"testing"

	"github.com/bravmi/gopl/chap7/eval"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	var tests = []struct {
		expr string
		want string
	}{
		{"x!", "x!"},
		{"x! + y", "x! + y"},
		{"+x! + y", "+x! + y"},
		{"+(x!) + y", "+x! + y"},
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

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  eval.Env
		want string
	}{
		{"x!", eval.Env{"x": 5}, "120"},
		{"x!", eval.Env{"x": 10}, "3.6288e+06"},
	}
	for _, test := range tests {
		testname := test.expr
		t.Run(testname, func(t *testing.T) {
			expr, err := eval.Parse(test.expr)
			assert.NoError(t, err)
			got := fmt.Sprintf("%.6g", expr.Eval(test.env))
			assert.Equal(t, test.want, got)
		})
	}
}
