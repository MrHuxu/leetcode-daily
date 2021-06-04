package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_openLock(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, openLock([]string{
		"0201", "0101", "0102", "1212", "2002",
	}, "0202"))
	assert.Equal(1, openLock([]string{
		"8888",
	}, "0009"))
	assert.Equal(-1, openLock([]string{
		"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888",
	}, "8888"))
	assert.Equal(-1, openLock([]string{"0000"}, "8888"))
}
