package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPossible(t *testing.T) {
	assert := assert.New(t)

	assert.True(isPossible([]int{9, 3, 5}))
	assert.False(isPossible([]int{1, 1, 1, 2}))
	assert.True(isPossible([]int{8, 5}))
	assert.True(isPossible([]int{1, 1000000000}))
	assert.False(isPossible([]int{2, 900000002}))
}
