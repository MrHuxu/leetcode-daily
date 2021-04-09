package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	assert := assert.New(t)

	assert.True(isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}))

	assert.False(isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}))
}
