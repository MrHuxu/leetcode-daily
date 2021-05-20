package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		{3},
		{9, 20},
		{15, 7},
	}, levelOrder(&TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}))

	assert.Equal([][]int{{1}}, levelOrder(&TreeNode{Val: 1}))

	assert.Empty(levelOrder(nil))
}
