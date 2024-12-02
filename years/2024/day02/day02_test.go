package day02_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2024/day02"
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
		2,
		4,
	},
	{
		input,
		213,
		0,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := slice.Reduce(slice.TrimSplit(tc.input), 0, func(l string, ret int) int {
			if day02.IsSafe(l, false) {
				return ret + 1
			}
			return ret
		})

		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		actual := slice.Reduce(slice.TrimSplit(tc.input), 0, func(l string, ret int) int {
			if day02.IsSafe(l, true) {
				return ret + 1
			}
			return ret
		})

		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
