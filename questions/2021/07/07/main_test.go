package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallest(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(13, kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 8))
	assert.Equal(2, kthSmallest([][]int{
		{1, 2}, {3, 3},
	}, 2))
}
