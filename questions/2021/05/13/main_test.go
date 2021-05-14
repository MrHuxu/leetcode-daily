package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ambiguousCoordinates(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{
		"(1, 23)", "(1, 2.3)", "(12, 3)", "(1.2, 3)",
	}, ambiguousCoordinates("(123)"))
	assert.Equal([]string{
		"(0, 0.011)", "(0.001, 1)",
	}, ambiguousCoordinates("(00011)"))
	assert.Equal([]string{
		"(0, 123)", "(0, 1.23)", "(0, 12.3)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)",
	}, ambiguousCoordinates("(0123)"))
	assert.Equal([]string{"(10, 0)"}, ambiguousCoordinates("(100)"))
}
