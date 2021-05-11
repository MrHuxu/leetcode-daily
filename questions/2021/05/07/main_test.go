package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDistance(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, minDistance(
		"sea", "eat",
	))
	assert.Equal(4, minDistance(
		"leetcode", "etco",
	))
}
