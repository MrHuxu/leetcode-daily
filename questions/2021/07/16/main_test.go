package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fourSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1},
	}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
