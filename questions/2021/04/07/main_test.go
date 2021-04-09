package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_halvesAreAlike(t *testing.T) {
	assert := assert.New(t)

	assert.True(halvesAreAlike("book"))
	assert.False(halvesAreAlike("textbook"))
	assert.False(halvesAreAlike("MerryChristmas"))
	assert.True(halvesAreAlike("AbCdEfGh"))
}
