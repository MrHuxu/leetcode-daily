package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{
		3, 4,
	}, searchRange([]int{
		5, 7, 7, 8, 8, 10,
	}, 8))
}
