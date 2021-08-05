package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stoneGame(t *testing.T) {
	assert := assert.New(t)

	assert.True(stoneGame([]int{5, 3, 4, 5}))
}
