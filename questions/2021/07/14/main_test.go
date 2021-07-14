package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_customSortString(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("cbad", customSortString("cba", "abcd"))
	assert.Equal("eexvw", customSortString("exv", "xwvee"))
}
