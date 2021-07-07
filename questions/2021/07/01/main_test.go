package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_grayCode(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0, 2, 3, 1}, grayCode(2))
	assert.Equal([]int{0, 1}, grayCode(1))
}
