package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkPossibility(t *testing.T) {
	assert := assert.New(t)

	assert.True(checkPossibility([]int{4, 2, 3}))
	assert.False(checkPossibility([]int{4, 2, 1}))
	assert.False(checkPossibility([]int{3, 4, 2, 3}))
	assert.True(checkPossibility([]int{5, 7, 1, 8}))
	assert.True(checkPossibility([]int{5, 7, 1}))
}
