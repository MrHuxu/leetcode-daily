package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findAndReplacePattern(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{
		"mee", "aqq",
	}, findAndReplacePattern([]string{
		"abc", "deq", "mee", "aqq", "dkd", "ccc",
	}, "abb"))
	assert.Equal([]string{
		"a", "b", "c",
	}, findAndReplacePattern([]string{
		"a", "b", "c",
	}, "a"))
}
