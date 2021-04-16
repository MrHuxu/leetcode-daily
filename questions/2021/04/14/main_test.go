package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partition(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		partition(
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 2,
								},
							},
						},
					},
				},
			}, 3,
		))
}
