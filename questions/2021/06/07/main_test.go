package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCostClimbingStairs(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(15, minCostClimbingStairs([]int{10, 15, 20}))
	assert.Equal(6, minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
