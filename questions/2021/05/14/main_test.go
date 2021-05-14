package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_flatten(t *testing.T) {
	assert := assert.New(t)

	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	flatten(tree)
	assert.Equal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}, tree)

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 6},
		},
	}
	flatten(tree)
	assert.Equal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		},
	}, tree)
}
