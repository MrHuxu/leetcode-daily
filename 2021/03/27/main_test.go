package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordSubsets(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		[]string{"facebook", "google", "leetcode"},
		wordSubsets(
			[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"},
		),
	)

	assert.Equal(
		[]string{"apple", "google", "leetcode"},
		wordSubsets(
			[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"},
		),
	)

	assert.Equal(
		[]string{"facebook", "google"},
		wordSubsets(
			[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"},
		),
	)

	assert.Equal(
		[]string{"google", "leetcode"},
		wordSubsets(
			[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lo", "eo"},
		),
	)

	assert.Equal(
		[]string{"facebook", "leetcode"},
		wordSubsets(
			[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"ec", "oc", "ceo"},
		),
	)
}
