package linkedlist_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/linkedlist"
)

func TestNewNode(t *testing.T) {
	n := linkedlist.NewNode(1)
	assert.Equals(t, n.Next(), nil)
	assert.Equals(t, n.Prev(), nil)
	assert.Equals(t, n.Value(), 1)
}

func TestLinkNext(t *testing.T) {
	n := linkedlist.NewNode(1)
	m := linkedlist.NewNode(2)
	n.LinkNext(m)
	assert.Equals(t, n.Next(), m)
	assert.Equals(t, n.Prev(), nil)
	assert.Equals(t, m.Prev(), n)
	assert.Equals(t, m.Next(), nil)
}

func TestLinkPrev(t *testing.T) {
	n := linkedlist.NewNode(1)
	m := linkedlist.NewNode(2)
	n.LinkPrev(m)
	assert.Equals(t, n.Prev(), m)
	assert.Equals(t, n.Next(), nil)
	assert.Equals(t, m.Next(), n)
	assert.Equals(t, m.Prev(), nil)
}

func TestLength(t *testing.T) {
	head := linkedlist.NewNode(0)
	n := head
	for i := 0; i < 10; i++ {
		m := linkedlist.NewNode(i)
		n.LinkNext(m)
		n = m
	}
	assert.Equals(t, head.Length(), 11)
}

func TestValueModification(t *testing.T) {
	type testType struct {
		innerVal int
	}
	tt := testType{1}
	n := linkedlist.NewNode(&tt)
	n.Value().innerVal += 1
	assert.Equals(t, tt.innerVal, 2)
}

func TestHead(t *testing.T) {
	n := linkedlist.NewNode(1)
	m := linkedlist.NewNode(2)
	o := linkedlist.NewNode(3)
	n.LinkNext(m)
	m.LinkNext(o)
	assert.Equals(t, o.Head(), n)
}

func TestTail(t *testing.T) {
	n := linkedlist.NewNode(1)
	m := linkedlist.NewNode(2)
	o := linkedlist.NewNode(3)
	n.LinkNext(m)
	m.LinkNext(o)
	assert.Equals(t, n.Tail(), o)
}
