package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	rotate(matrix)
	assert.Equal([][]int{
		{7, 4, 1}, {8, 5, 2}, {9, 6, 3},
	}, matrix)
}
