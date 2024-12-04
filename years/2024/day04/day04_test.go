package day04_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/years/2024/day04"
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
		18,
		9,
	},
	{
		input,
		2578,
		1972,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		d := day04.NewWordSearch(tc.input)
		actual := d.NumInstancesXMas()
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		d := day04.NewWordSearch(tc.input)
		actual := d.NumInstancesCrossMas()
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
