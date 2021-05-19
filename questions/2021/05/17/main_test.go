package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestStrChain(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, longestStrChain([]string{
		"a", "b", "ba", "bca", "bda", "bdca",
	}))
	assert.Equal(5, longestStrChain([]string{
		"xbc", "pcxbcf", "xb", "cxbc", "pcxbc",
	}))
	assert.Equal(4, longestStrChain([]string{
		"bdca", "bda", "ca", "dca", "a",
	}))
}
