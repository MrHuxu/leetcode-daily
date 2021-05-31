package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_suggestedProducts(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]string{
		{"mobile", "moneypot", "monitor"},
		{"mobile", "moneypot", "monitor"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
	}, suggestedProducts([]string{
		"mobile", "mouse", "moneypot", "monitor", "mousepad",
	}, "mouse"))

	assert.Equal([][]string{
		{"havana"},
		{"havana"},
		{"havana"},
		{"havana"},
		{"havana"},
		{"havana"},
	}, suggestedProducts([]string{"havana"}, "havana"))

	assert.Equal([][]string{
		{"baggage", "bags", "banner"},
		{"baggage", "bags", "banner"},
		{"baggage", "bags"},
		{"bags"},
	}, suggestedProducts([]string{
		"bags", "baggage", "banner", "box", "cloths",
	}, "bags"))

	assert.Equal([][]string{
		{}, {}, {}, {}, {}, {}, {},
	}, suggestedProducts([]string{"havana"}, "tatiana"))
}
