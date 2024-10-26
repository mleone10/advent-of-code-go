package day04_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2023/day04"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo int
}{
	{
		testInput,
		13,
		30,
	},
	{
		input,
		21138,
		7185540,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := day04.PilePoints(slice.TrimSplit(tc.input))
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		actual := day04.NumCards(slice.TrimSplit(tc.input))
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
