package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	}, removeNthFromEnd(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}, 2))

	assert.Nil(removeNthFromEnd(&ListNode{
		Val: 1,
	}, 1))

	assert.Equal(&ListNode{
		Val: 1,
	}, removeNthFromEnd(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}, 1))
}
