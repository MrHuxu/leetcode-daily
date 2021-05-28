package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDuplicate(t *testing.T) {
	assert := assert.New(t)

	// assert.Equal([][]string{
	// 	{"root/a/1.txt", "root/c/3.txt"},
	// 	{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"},
	// }, findDuplicate([]string{
	// 	"root/a 1.txt(abcd) 2.txt(efgh)",
	// 	"root/c 3.txt(abcd)",
	// 	"root/c/d 4.txt(efgh)",
	// 	"root 4.txt(efgh)",
	// }))
	assert.Nil(findDuplicate([]string{
		"root/a 1.txt(abcd) 2.txt(efsfgh)", "root/c 3.txt(abdfcd)", "root/c/d 4.txt(efggdfh)",
	}))
}
