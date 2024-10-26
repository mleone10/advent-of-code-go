package queue_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/queue"
)

func TestNewQueue(t *testing.T) {
	var q queue.Queue[int]
	assert.Equals(t, q.Length(), 0)
}

func TestPush(t *testing.T) {
	var q queue.Queue[int]
	q.Push(1)
	q.Push(2)
	assert.Equals(t, q.Length(), 2)
}

func TestPop(t *testing.T) {
	var q queue.Queue[int]
	q.Push(1)
	q.Push(2)
	assert.Equals(t, q.Pop(), 1)
	assert.Equals(t, q.Pop(), 2)
}
