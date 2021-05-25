package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evalRPN(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(9, evalRPN([]string{
		"2", "1", "+", "3", "*",
	}))
	assert.Equal(6, evalRPN([]string{
		"4", "13", "5", "/", "+",
	}))
	assert.Equal(22, evalRPN([]string{
		"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+",
	}))
}
