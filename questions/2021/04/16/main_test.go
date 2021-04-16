package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abcd", removeDuplicates("abcd", 2))
	assert.Equal("aa", removeDuplicates("deeedbbcccbdaa", 3))
	assert.Equal("ps", removeDuplicates("pbbcggttciiippooaais", 2))
}
