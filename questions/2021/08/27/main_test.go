package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLUSlength(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, findLUSlength([]string{"aba", "cdc", "eae"}))
	assert.Equal(-1, findLUSlength([]string{"aaa", "aaa", "aa"}))
}
