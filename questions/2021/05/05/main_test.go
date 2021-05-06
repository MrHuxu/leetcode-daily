package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, jump([]int{2, 3, 1, 1, 4}))
	assert.Equal(2, jump([]int{2, 3, 0, 1, 4}))
}
