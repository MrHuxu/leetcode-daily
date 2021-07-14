package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPeakElement(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, findPeakElement([]int{1, 2, 3, 1}))
	assert.Equal(5, findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	assert.Equal(0, findPeakElement([]int{1}))
	assert.Equal(1, findPeakElement([]int{1, 2}))
}
