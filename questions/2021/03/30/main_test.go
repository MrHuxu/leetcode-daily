package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxEnvelopes(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, maxEnvelopes([][]int{
		{5, 4}, {6, 4}, {6, 7}, {2, 3},
	}))
	assert.Equal(1, maxEnvelopes([][]int{
		{1, 1}, {1, 1}, {1, 1},
	}))
	assert.Equal(4, maxEnvelopes([][]int{
		{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1},
	}))
	assert.Equal(3, maxEnvelopes([][]int{
		{46, 89}, {50, 53}, {52, 68}, {72, 45}, {77, 81},
	}))
}
