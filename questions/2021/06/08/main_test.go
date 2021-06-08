package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}, buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
