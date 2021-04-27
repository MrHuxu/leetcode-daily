package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPowerOfThree(t *testing.T) {
	assert := assert.New(t)

	assert.True(isPowerOfThree(27))
	assert.False(isPowerOfThree(0))
	assert.True(isPowerOfThree(9))
	assert.False(isPowerOfThree(45))
}
