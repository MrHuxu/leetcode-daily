package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	assert := assert.New(t)

	tree1 := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 5},
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
	}
	assert.Equal(tree1, lowestCommonAncestor(tree1, tree1.Left, tree1.Right))
	assert.Equal(tree1.Left, lowestCommonAncestor(tree1, tree1.Left, tree1.Left.Right))
}
