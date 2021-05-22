package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveNQueens(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]string{
		{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."},
	}, solveNQueens(4))
	assert.Equal([][]string{
		{"Q"},
	}, solveNQueens(1))
}
