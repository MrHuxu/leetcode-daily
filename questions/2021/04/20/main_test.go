package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_preorder(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{
		1, 3, 5, 6, 2, 4,
	}, preorder(&Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}))

	assert.Equal([]int{
		1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10,
	}, preorder(&Node{
		Val: 1,
		Children: []*Node{
			{Val: 2},
			{
				Val: 3,
				Children: []*Node{
					{Val: 6},
					{
						Val: 7,
						Children: []*Node{
							{
								Val: 11,
								Children: []*Node{
									{Val: 14},
								},
							},
						},
					},
				},
			},
			{
				Val: 4,
				Children: []*Node{
					{
						Val: 8,
						Children: []*Node{
							{Val: 12},
						},
					},
				},
			},
			{
				Val: 5,
				Children: []*Node{
					{
						Val: 9,
						Children: []*Node{
							{Val: 13},
						},
					},
					{Val: 10},
				},
			},
		},
	}))
}
