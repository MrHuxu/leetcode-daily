package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumMatrix(t *testing.T) {
	assert := assert.New(t)

	nm := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	assert.Equal(8, nm.SumRegion(2, 1, 4, 3))
	assert.Equal(11, nm.SumRegion(1, 1, 2, 2))
	assert.Equal(12, nm.SumRegion(1, 2, 2, 4))
}
