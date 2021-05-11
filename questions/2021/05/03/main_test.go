package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runningSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{
		1, 3, 6, 10,
	}, runningSum([]int{
		1, 2, 3, 4,
	}))
	assert.Equal([]int{
		1, 2, 3, 4, 5,
	}, runningSum([]int{
		1, 1, 1, 1, 1,
	}))
	assert.Equal([]int{
		3, 4, 6, 16, 17,
	}, runningSum([]int{
		3, 1, 2, 10, 1,
	}))
}
