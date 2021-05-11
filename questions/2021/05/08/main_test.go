package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_superpalindromesInRange(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, superpalindromesInRange("4", "1000"))
	assert.Equal(1, superpalindromesInRange("1", "2"))
	assert.Equal(2, superpalindromesInRange("40000000000000000", "50000000000000000"))
	// 200000002, 200010002
	assert.Equal(45, superpalindromesInRange("38455498359", "999999999999999999"))
}
