package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxForm(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, findMaxForm(
		[]string{"10", "0001", "111001", "1", "0"}, 5, 3,
	))

	assert.Equal(2, findMaxForm(
		[]string{"10", "0", "1"}, 1, 1,
	))

	assert.Equal(4, findMaxForm(
		[]string{"0", "0", "1", "1"}, 2, 2,
	))

	assert.Equal(3, findMaxForm(
		[]string{"111", "1000", "1000", "1000"}, 9, 3,
	))

	assert.Equal(3, findMaxForm(
		[]string{"0001", "0101", "1000", "1000"}, 9, 3,
	))

	assert.Equal(17, findMaxForm(
		[]string{"0", "11", "1000", "01", "0", "101", "1", "1", "1", "0", "0", "0", "0", "1", "0", "0110101", "0", "11", "01", "00", "01111", "0011", "1", "1000", "0", "11101", "1", "0", "10", "0111"}, 9, 80,
	))
}
