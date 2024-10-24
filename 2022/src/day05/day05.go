package day05

import (
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/stack"
)

type Day05 []stack.Stack[string]

type Move struct {
	Num, From, To int
}

func NewMove(move string) Move {
	// move 1 from 2 to 3 => Move{1, 2, 3}
	parts := strings.Split(move, " ")
	num, _ := strconv.Atoi(parts[1])
	from, _ := strconv.Atoi(parts[3])
	to, _ := strconv.Atoi(parts[5])
	// Stacks are 1-indexed in the input, we'll switch to 0-indexing.
	return Move{num, from - 1, to - 1}
}

func (d Day05) ApplyMove(m Move, maintainOrder bool) {
	switch maintainOrder {
	case true:
		var temp stack.Stack[string]
		for i := 0; i < m.Num; i++ {
			temp.Push(d[m.From].Pop())
		}
		for i := 0; i < m.Num; i++ {
			d[m.To].Push(temp.Pop())
		}
	case false:
		for i := 0; i < m.Num; i++ {
			d[m.To].Push(d[m.From].Pop())
		}
	}
}

func (d Day05) GetTopString() string {
	var top string
	for _, s := range d {
		top += s.Peek()
	}
	return top
}

func (d Day05) Copy() Day05 {
	s := make(Day05, len(d))
	for i, stk := range d {
		s[i] = stk.Copy()
	}
	return s
}
