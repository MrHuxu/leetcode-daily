package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{
		"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
	}, letterCombinations("23"))
	assert.Empty(letterCombinations(""))
	assert.Equal([]string{
		"a", "b", "c",
	}, letterCombinations("2"))
}
