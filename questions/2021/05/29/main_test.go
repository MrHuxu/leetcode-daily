package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_totalNQueens(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, totalNQueens(4))
	assert.Equal(1, totalNQueens(1))
}
