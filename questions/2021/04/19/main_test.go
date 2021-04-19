package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum4(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(7, combinationSum4([]int{
		1, 2, 3,
	}, 4))
	assert.Zero(combinationSum4([]int{
		9,
	}, 3))
}
