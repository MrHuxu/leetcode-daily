package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxScore(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(12, maxScore([]int{
		1, 2, 3, 4, 5, 6, 1,
	}, 3))
	assert.Equal(4, maxScore([]int{
		2, 2, 2,
	}, 2))
	assert.Equal(55, maxScore([]int{
		9, 7, 7, 9, 7, 7, 9,
	}, 7))
}
