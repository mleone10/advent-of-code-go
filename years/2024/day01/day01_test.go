package day01_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/years/2024/day01"
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
		11,
		31,
	},
	{
		input,
		2430334,
		28786472,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		d := day01.NewDay01(tc.input)

		actual := d.Distance()

		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		d := day01.NewDay01(tc.input)

		actual := d.SimilarityScore()

		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}