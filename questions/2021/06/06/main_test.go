package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestConsecutive(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	assert.Equal(9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	assert.Equal(5, longestConsecutive([]int{-6, 8, -5, 7, -9, -1, -7, -6, -9, -7, 5, 7, -1, -8, -8, -2, 0}))
}
