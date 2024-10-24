package day01_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2018/day01"
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
		4,
		10,
	},
	{
		input,
		592,
		241,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := day01.CalcFinalFreq(slice.Map(slice.TrimSplit(tc.input), func(f string) int { return mth.Atoi(f) }))
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		actual := day01.FindFirstDuplicateFreq(slice.Map(slice.TrimSplit(tc.input), func(f string) int { return mth.Atoi(f) }))
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
