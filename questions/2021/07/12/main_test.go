package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIsomorphic(t *testing.T) {
	assert := assert.New(t)

	assert.True(isIsomorphic("egg", "add"))
	assert.False(isIsomorphic("foo", "bar"))
	assert.True(isIsomorphic("paper", "title"))
	assert.False(isIsomorphic("badc", "baba"))
}
