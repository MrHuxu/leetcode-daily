package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSubstrings(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, countSubstrings("abc"))
	assert.Equal(6, countSubstrings("aaa"))
}
