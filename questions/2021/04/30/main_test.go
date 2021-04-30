package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_powerfulIntegers(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{
		2, 4, 10, 3, 5, 7, 9,
	}, powerfulIntegers(2, 3, 10))
	assert.Equal([]int{
		2, 6, 4, 8, 10, 14,
	}, powerfulIntegers(3, 5, 15))
	assert.Equal([]int{
		2, 3, 5, 9,
	}, powerfulIntegers(2, 1, 10))
}
