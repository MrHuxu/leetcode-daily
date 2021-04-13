package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constructArray(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 2, 3}, constructArray(3, 1))
	assert.Equal([]int{1, 3, 2}, constructArray(3, 2))
}
