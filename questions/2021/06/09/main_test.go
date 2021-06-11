package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxResult(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(7, maxResult([]int{
		1, -1, -2, 4, -7, 3,
	}, 2))
	assert.Equal(17, maxResult([]int{
		10, -5, -2, 4, 0, 3,
	}, 3))
	assert.Equal(0, maxResult([]int{
		1, -5, -20, 4, -1, 3, -6, -3,
	}, 2))
}
