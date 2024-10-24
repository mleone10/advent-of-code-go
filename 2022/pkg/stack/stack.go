package stack

type Stack[T any] struct {
	stack []T
}

func New[T comparable](items ...T) Stack[T] {
	var s Stack[T]
	s.Fill(items...)
	return s
}

func (s *Stack[T]) Push(item T) {
	s.stack = append(s.stack, item)
}

func (s *Stack[T]) Pop() T {
	ret := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return ret
}

func (s *Stack[T]) Len() int {
	return len(s.stack)
}

func (s *Stack[T]) Fill(items ...T) {
	for _, i := range items {
		s.Push(i)
	}
}

func (s *Stack[T]) Peek() T {
	return s.stack[len(s.stack)-1]
}

func (s *Stack[T]) Copy() Stack[T] {
	var cpy Stack[T]
	cpy.stack = make([]T, len(s.stack))
	copy(cpy.stack, s.stack)
	return cpy
}
