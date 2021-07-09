package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLength(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, findLength([]int{
		1, 2, 3, 2, 1,
	}, []int{
		3, 2, 1, 4, 7,
	}))
	assert.Equal(5, findLength([]int{
		0, 0, 0, 0, 0,
	}, []int{
		0, 0, 0, 0, 0,
	}))
}
