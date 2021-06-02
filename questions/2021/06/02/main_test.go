package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isInterleave(t *testing.T) {
	assert := assert.New(t)

	assert.True(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	assert.False(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	assert.True(isInterleave("", "", ""))
}
