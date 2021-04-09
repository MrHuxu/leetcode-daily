package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularQueue(t *testing.T) {
	assert := assert.New(t)

	cq := Constructor(3)
	assert.True(cq.EnQueue(1))
	assert.True(cq.EnQueue(2))
	assert.True(cq.EnQueue(3))
	assert.False(cq.EnQueue(4))
	assert.Equal(3, cq.Rear())
	assert.True(cq.IsFull())
	assert.True(cq.DeQueue())
	assert.True(cq.EnQueue(4))
	assert.Equal(4, cq.Rear())

	// ["MyCircularQueue","enQueue","Rear","Rear","deQueue","enQueue","Rear","deQueue","Front","deQueue","deQueue","deQueue"]
	// [[6],[6],[],[],[],[5],[],[],[],[],[],[]]

	cq1 := Constructor(6)
	assert.True(cq1.EnQueue(6))
	assert.Equal(6, cq1.Rear())
	assert.Equal(6, cq1.Rear())
	assert.True(cq1.DeQueue())
	assert.True(cq1.EnQueue(5))
	assert.Equal(5, cq1.Rear())
	assert.True(cq1.DeQueue())
	assert.Equal(-1, cq1.Front())
	assert.False(cq1.DeQueue())
	assert.False(cq1.DeQueue())
	assert.False(cq1.DeQueue())
}
