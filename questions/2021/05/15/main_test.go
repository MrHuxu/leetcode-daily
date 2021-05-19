package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isNumber(t *testing.T) {
	assert := assert.New(t)

	assert.True(isNumber("0"))
	assert.False(isNumber("e"))
	assert.False(isNumber("."))
	assert.True(isNumber(".1"))
	assert.True(isNumber("1E9"))
}
