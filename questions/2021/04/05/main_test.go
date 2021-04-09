package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIdealPermutation(t *testing.T) {
	assert := assert.New(t)

	assert.True(isIdealPermutation([]int{1, 0, 2}))
	assert.False(isIdealPermutation([]int{1, 2, 0}))
}
