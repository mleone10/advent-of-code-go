package day20_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/years/2015/day20"
)

var tcs = []struct {
	input           int
	expectedPartOne int
	expectedPartTwo int
}{
	{
		10, 1, 1,
	},
	{
		70, 4, 4,
	},
	{
		150, 8, 8,
	},
	{
		36000000, 831600, 884520,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := day20.FindHouseWithMinPresents(tc.input, tc.input/10, 10)
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		actual := day20.FindHouseWithMinPresents(tc.input, 50, 11)
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
