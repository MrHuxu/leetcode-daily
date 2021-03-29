package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_advantageCount(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{
		2, 11, 7, 15,
	}, advantageCount([]int{
		2, 7, 11, 15,
	}, []int{
		1, 10, 4, 11,
	}))

	assert.Equal([]int{
		24, 32, 8, 12,
	}, advantageCount([]int{
		12, 24, 8, 32,
	}, []int{
		13, 25, 32, 11,
	}))

	assert.Equal([]int{
		341560426, 718967141, 189971378, 23521218, 339517772,
	}, advantageCount([]int{
		718967141, 189971378, 341560426, 23521218, 339517772,
	}, []int{
		967482459, 978798455, 744530040, 3454610, 940238504,
	}))
}
