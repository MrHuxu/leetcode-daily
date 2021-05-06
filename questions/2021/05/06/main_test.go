package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedListToBST(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		&TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val: -10,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
		sortedListToBST(
			&ListNode{
				Val: -10,
				Next: &ListNode{
					Val: -3,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		),
	)

	assert.Nil(sortedListToBST(nil))
	assert.Equal(
		&TreeNode{
			Val: 0,
		},
		sortedListToBST(&ListNode{Val: 0}),
	)
	assert.Equal(
		&TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 1},
		},
		sortedListToBST(&ListNode{
			Val:  1,
			Next: &ListNode{Val: 3},
		}),
	)
}
