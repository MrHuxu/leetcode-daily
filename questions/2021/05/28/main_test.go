package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumUniqueSubarray(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(17, maximumUniqueSubarray([]int{
		4, 2, 4, 5, 6,
	}))
	assert.Equal(8, maximumUniqueSubarray([]int{
		5, 2, 1, 2, 5, 2, 1, 2, 5,
	}))
}
