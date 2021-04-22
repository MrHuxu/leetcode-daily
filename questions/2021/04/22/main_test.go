package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_leastBricks(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, leastBricks([][]int{
		{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1},
	}))
}
