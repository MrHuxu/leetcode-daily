package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, minOperations(3))
	assert.Equal(4, minOperations(4))
	assert.Equal(9, minOperations(6))
}
