package linkedlist

type nodeVal interface {
	any
}

type Node[T nodeVal] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

func NewNode[T nodeVal](val T) *Node[T] {
	return &Node[T]{value: val}
}

func (n *Node[T]) LinkPrev(m *Node[T]) {
	n.prev = m
	m.next = n
}

func (n *Node[T]) LinkNext(m *Node[T]) {
	n.next = m
	m.prev = n
}

func (n Node[T]) Value() T {
	return n.value
}

func (n Node[T]) Next() *Node[T] {
	return n.next
}

func (n Node[T]) Prev() *Node[T] {
	return n.prev
}

func (n *Node[T]) Head() *Node[T] {
	if n.prev == nil {
		return n
	}
	return n.prev.Head()
}

func (n *Node[T]) Tail() *Node[T] {
	if n.next == nil {
		return n
	}
	return n.next.Tail()
}

func (n Node[T]) Length() int {
	if n.next != nil {
		return 1 + n.next.Length()
	}
	return 1
}
