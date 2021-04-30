package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
	assert.Equal(1, uniquePathsWithObstacles([][]int{
		{0, 1},
		{0, 0},
	}))
}
