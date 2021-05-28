package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProduct(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(16, maxProduct([]string{
		"abcw", "baz", "foo", "bar", "xtfn", "abcdef",
	}))
	assert.Equal(4, maxProduct([]string{
		"a", "ab", "abc", "d", "cd", "bcd", "abcd",
	}))
	assert.Equal(0, maxProduct([]string{
		"a", "aa", "aaa", "aaaa",
	}))
}
