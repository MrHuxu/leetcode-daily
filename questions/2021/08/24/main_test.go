package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_complexNumberMultiply(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("0+2i", complexNumberMultiply("1+1i", "1+1i"))
	assert.Equal("0+-2i", complexNumberMultiply("1+-1i", "1+-1i"))
	assert.Equal("0+0i", complexNumberMultiply("1+-1i", "0+0i"))
}
