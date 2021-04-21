package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumTotal(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(11, minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
	assert.Equal(-10, minimumTotal([][]int{
		{-10},
	}))
}
