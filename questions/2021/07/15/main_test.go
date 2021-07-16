package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_triangleNumber(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, triangleNumber([]int{2, 2, 3, 4}))
	assert.Equal(4, triangleNumber([]int{4, 2, 3, 4}))
	assert.Equal(19, triangleNumber([]int{48, 66, 61, 46, 94, 75}))
}
