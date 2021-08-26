package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidSerialization(t *testing.T) {
	assert := assert.New(t)

	assert.True(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	assert.False(isValidSerialization("1,#"))
	assert.False(isValidSerialization("9,#,#,1"))
}
