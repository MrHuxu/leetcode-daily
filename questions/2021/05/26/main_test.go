package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minPartitions(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, minPartitions("32"))
	assert.Equal(8, minPartitions("82734"))
	assert.Equal(9, minPartitions("27346209830709182346"))
}
