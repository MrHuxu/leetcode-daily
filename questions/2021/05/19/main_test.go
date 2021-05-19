package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minMoves2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, minMoves2([]int{1, 2, 3}))
	assert.Equal(16, minMoves2([]int{1, 10, 2, 9}))
	assert.Equal(14, minMoves2([]int{1, 0, 0, 8, 6}))
}
