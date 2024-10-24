package stack_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/stack"
)

func TestNew(t *testing.T) {
	s := stack.New(1, 2, 3, 4, 5)

	assert.Equal(t, s.Len(), 5)
	assert.Equal(t, s.Pop(), 5)
}

func TestPush(t *testing.T) {
	var s stack.Stack[int]

	s.Push(1)
	assert.Equal(t, s.Len(), 1)

	s.Push(2)
	assert.Equal(t, s.Len(), 2)
}

func TestPop(t *testing.T) {
	var s stack.Stack[int]

	s.Push(1)
	s.Push(2)

	assert.Equal(t, s.Pop(), 2)
	assert.Equal(t, s.Pop(), 1)
}

func TestLen(t *testing.T) {
	var s stack.Stack[int]

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	assert.Equal(t, s.Len(), 10)
}

func TestFill(t *testing.T) {
	var s stack.Stack[int]

	s.Fill(1, 2, 3, 4, 5)
	assert.Equal(t, s.Len(), 5)
	assert.Equal(t, s.Pop(), 5)
}

func TestPeek(t *testing.T) {
	var s stack.Stack[int]

	s.Push(1)
	s.Push(2)

	assert.Equal(t, s.Peek(), 2)
	assert.Equal(t, s.Len(), 2)
}

func TestCopy(t *testing.T) {
	var s stack.Stack[int]

	s.Push(1)
	s.Push(2)

	r := s.Copy()

	assert.NotEqual(t, &s, &r)
	assert.Equal(t, s.Len(), r.Len())
	assert.Equal(t, s.Peek(), r.Peek())
}
