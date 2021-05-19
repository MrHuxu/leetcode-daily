package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCameraCover(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, minCameraCover(&TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{},
			Right: &TreeNode{},
		},
	}))
	assert.Equal(2, minCameraCover(&TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left: &TreeNode{
					Right: &TreeNode{},
				},
			},
		},
	}))
	assert.Equal(2, minCameraCover(&TreeNode{
		Left: &TreeNode{},
		Right: &TreeNode{
			Right: &TreeNode{},
		},
	}))
	assert.Equal(2, minCameraCover(&TreeNode{
		Left: &TreeNode{
			Right: &TreeNode{
				Left: &TreeNode{
					Right: &TreeNode{
						Left: &TreeNode{},
					},
				},
			},
		},
	}))
}
