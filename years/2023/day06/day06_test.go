package day06_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2023/day06"
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
		288,
		71503,
	},
	{
		input,
		252000,
		36992486,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := day06.MarginOfError(slice.TrimSplit(tc.input))
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		actual := day06.SingleRaceWinningSolutions(slice.TrimSplit(tc.input))
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
