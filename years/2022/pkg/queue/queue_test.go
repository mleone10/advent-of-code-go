package queue_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/queue"
)

func TestNewQueue(t *testing.T) {
	var q queue.Queue[int]
	assert.Equal(t, q.Length(), 0)
}

func TestPush(t *testing.T) {
	var q queue.Queue[int]
	q.Push(1)
	q.Push(2)
	assert.Equal(t, q.Length(), 2)
}

func TestPop(t *testing.T) {
	var q queue.Queue[int]
	q.Push(1)
	q.Push(2)
	assert.Equal(t, q.Pop(), 1)
	assert.Equal(t, q.Pop(), 2)
}
