package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedIterator(t *testing.T) {
	assert := assert.New(t)

	ni1 := Constructor([]*NestedInteger{
		{
			isInteger: false,
			list: []*NestedInteger{
				{isInteger: true, integer: 1},
				{isInteger: true, integer: 1},
			},
		},
		{
			isInteger: true,
			integer:   2,
		},
		{
			isInteger: false,
			list: []*NestedInteger{
				{isInteger: true, integer: 1},
				{isInteger: true, integer: 1},
			},
		},
	})
	assert.True(ni1.HasNext())
	assert.Equal(1, ni1.Next())
	assert.Equal(1, ni1.Next())
	assert.Equal(2, ni1.Next())
	assert.Equal(1, ni1.Next())
	assert.Equal(1, ni1.Next())
	assert.False(ni1.HasNext())

	ni2 := Constructor([]*NestedInteger{
		{
			isInteger: true,
			integer:   1,
		},
		{
			isInteger: false,
			list: []*NestedInteger{
				{
					isInteger: true,
					integer:   4,
				},
				{
					isInteger: false,
					list: []*NestedInteger{
						{isInteger: true, integer: 6},
					},
				},
			},
		},
	})
	assert.True(ni2.HasNext())
	assert.Equal(1, ni2.Next())
	assert.Equal(4, ni2.Next())
	assert.Equal(6, ni2.Next())
	assert.False(ni2.HasNext())
}
