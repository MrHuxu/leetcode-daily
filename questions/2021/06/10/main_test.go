package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCalendar(t *testing.T) {
	assert := assert.New(t)

	mc := Constructor()
	assert.True(mc.Book(20, 29))
	assert.False(mc.Book(13, 22))
	assert.True(mc.Book(44, 50))
	assert.True(mc.Book(1, 7))
	assert.False(mc.Book(2, 10))
	assert.True(mc.Book(14, 20))
	assert.False(mc.Book(19, 25))
	assert.True(mc.Book(36, 42))
	assert.False(mc.Book(45, 50))
	assert.False(mc.Book(47, 50))
}

/*

["MyCalendar","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book"]
[[],[20,29],[13,22],[44,50],[1,7],[2,10],[14,20],[19,25],[36,42],[45,50],[47,50],[39,45],[44,50],[16,25],[45,50],[45,50],[12,20],[21,29],[11,20],[12,17],[34,40],[10,18],[38,44],[23,32],[38,44],[15,20],[27,33],[34,42],[44,50],[35,40],[24,31]]

*/
