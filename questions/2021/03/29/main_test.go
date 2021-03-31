package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_flipMatchVoyage(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{-1}, flipMatchVoyage(&TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}, []int{2, 1}))

	assert.Equal([]int{1}, flipMatchVoyage(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}, []int{1, 3, 2}))

	assert.Empty(flipMatchVoyage(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}, []int{1, 2, 3}))

	assert.Empty(flipMatchVoyage(&TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}, []int{1, 2}))
}
