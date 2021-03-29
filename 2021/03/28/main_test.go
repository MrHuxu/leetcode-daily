package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_originalDigits(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("012", originalDigits("owoztneoer"))
	assert.Equal("45", originalDigits("fviefuro"))
	assert.Equal("00", originalDigits("zerozero"))
}
