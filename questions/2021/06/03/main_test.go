package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, maxArea(5, 4, []int{1, 2, 4}, []int{1, 3}))
	assert.Equal(6, maxArea(5, 4, []int{3, 1}, []int{1}))
	assert.Equal(9, maxArea(5, 4, []int{3}, []int{3}))
}
