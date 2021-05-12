package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPrimes(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, countPrimes(10))
	assert.Equal(0, countPrimes(0))
	assert.Equal(0, countPrimes(1))
	assert.Equal(0, countPrimes(2))
	assert.Equal(1, countPrimes(3))
}
