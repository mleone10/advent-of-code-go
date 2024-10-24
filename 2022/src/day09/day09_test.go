package day09_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/src/day09"
)

//go:embed test_input.txt
var testInput string

//go:embed test_input_2.txt
var testInput2 string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo int
}{
	{
		input:           testInput,
		expectedPartOne: 13,
		expectedPartTwo: 1,
	},
	{
		input:           input,
		expectedPartOne: 6236,
		expectedPartTwo: 2449,
	},
}

func TestNewRope(t *testing.T) {
	r := day09.NewRope(3)
	assert.Equal(t, r.Length(), 3)
}

func TestMoveN(t *testing.T) {
	r := day09.NewRope(2)
	day09.MoveN(r, day09.Cmds["D"], 2)
	assert.Equal(t, r.Head().Value().Pos.X, 0)
	assert.Equal(t, r.Head().Value().Pos.Y, 2)
	assert.Equal(t, r.Tail().Value().Pos.X, 0)
	assert.Equal(t, r.Tail().Value().Pos.Y, 1)

	day09.MoveN(r, day09.Cmds["R"], 2)
	assert.Equal(t, r.Head().Value().Pos.X, 2)
	assert.Equal(t, r.Head().Value().Pos.Y, 2)
	assert.Equal(t, r.Tail().Value().Pos.X, 1)
	assert.Equal(t, r.Tail().Value().Pos.Y, 2)

	day09.MoveN(r, day09.Cmds["U"], 3)
	assert.Equal(t, r.Head().Value().Pos.X, 2)
	assert.Equal(t, r.Head().Value().Pos.Y, -1)
	assert.Equal(t, r.Tail().Value().Pos.X, 2)
	assert.Equal(t, r.Tail().Value().Pos.Y, 0)

	day09.MoveN(r, day09.Cmds["L"], 3)
	assert.Equal(t, r.Head().Value().Pos.X, -1)
	assert.Equal(t, r.Head().Value().Pos.Y, -1)
	assert.Equal(t, r.Tail().Value().Pos.X, 0)
	assert.Equal(t, r.Tail().Value().Pos.Y, -1)

	day09.MoveN(r, day09.Cmds["L"], 1)
	assert.Equal(t, r.Head().Value().Pos.X, -2)
	assert.Equal(t, r.Head().Value().Pos.Y, -1)
	assert.Equal(t, r.Tail().Value().Pos.X, -1)
	assert.Equal(t, r.Tail().Value().Pos.Y, -1)
}

func TestHeadMovesDiagonal(t *testing.T) {
	r := day09.NewRope(2)
	day09.MoveN(r, day09.Cmds["R"], 1)
	day09.MoveN(r, day09.Cmds["D"], 1)
	day09.MoveN(r, day09.Cmds["R"], 1)
	assert.Equal(t, r.Tail().Value().Pos.X, 1)
	assert.Equal(t, r.Tail().Value().Pos.Y, 1)
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		r := day09.NewRope(2)
		day09.SimulateMoves(r, strings.Split(strings.TrimSpace(tc.input), "\n"))
		assert.Equal(t, len(r.Tail().Value().Visited()), tc.expectedPartOne)
	}
}

func TestSolvePartTwoLongerTest(t *testing.T) {
	r := day09.NewRope(10)
	day09.SimulateMoves(r, strings.Split(strings.TrimSpace(testInput2), "\n"))
	assert.Equal(t, len(r.Tail().Value().Visited()), 36)
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		r := day09.NewRope(10)
		day09.SimulateMoves(r, strings.Split(strings.TrimSpace(tc.input), "\n"))
		assert.Equal(t, len(r.Tail().Value().Visited()), tc.expectedPartTwo)
	}
}
