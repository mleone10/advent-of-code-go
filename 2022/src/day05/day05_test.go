package day05_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/stack"
	"github.com/mleone10/advent-of-code-2022/src/day05"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	moves           string
	state           day05.Day05
	expectedPartOne string
	expectedPartTwo string
}{
	{
		moves: testInput,
		state: day05.Day05{
			stack.New("Z", "N"),
			stack.New("M", "C", "D"),
			stack.New("P"),
		},
		expectedPartOne: "CMZ",
		expectedPartTwo: "MCD",
	},
	{
		moves: input,
		state: day05.Day05{
			stack.New("N", "B", "D", "T", "V", "G", "Z", "J"),
			stack.New("S", "R", "M", "D", "W", "P", "F"),
			stack.New("V", "C", "R", "S", "Z"),
			stack.New("R", "T", "J", "Z", "P", "H", "G"),
			stack.New("T", "C", "J", "N", "D", "Z", "Q", "F"),
			stack.New("N", "V", "P", "W", "G", "S", "F", "M"),
			stack.New("G", "C", "V", "B", "P", "Q"),
			stack.New("Z", "B", "P", "N"),
			stack.New("W", "P", "J"),
		},
		expectedPartOne: "GFTNRBZPF",
		expectedPartTwo: "VRQWPDSGP",
	},
}

func TestApplyMove(t *testing.T) {
	d := day05.Day05{
		stack.New("1", "2", "3", "4"),
		stack.New[string](),
	}
	m := day05.Move{2, 0, 1}

	d.ApplyMove(m, false)

	assert.Equal(t, d[0].Len(), 2)
	assert.Equal(t, d[0].Peek(), "2")
	assert.Equal(t, d[1].Len(), 2)
	assert.Equal(t, d[1].Peek(), "3")

	d.ApplyMove(m, true)

	assert.Equal(t, d[0].Len(), 0)
	assert.Equal(t, d[1].Len(), 4)
	assert.Equal(t, d[1].Peek(), "2")
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		s := tc.state.Copy()
		for _, move := range strings.Split(strings.TrimSpace(tc.moves), "\n") {
			m := day05.NewMove(move)
			s.ApplyMove(m, false)
		}
		assert.Equal(t, s.GetTopString(), tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		s := tc.state.Copy()
		for _, move := range strings.Split(strings.TrimSpace(tc.moves), "\n") {
			m := day05.NewMove(move)
			s.ApplyMove(m, true)
		}
		assert.Equal(t, s.GetTopString(), tc.expectedPartTwo)
	}
}
