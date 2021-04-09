package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestValidParentheses(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, longestValidParentheses("(()"))
	assert.Equal(4, longestValidParentheses(")()())"))
	assert.Equal(2, longestValidParentheses("()(()"))
	assert.Zero(longestValidParentheses(""))
}
