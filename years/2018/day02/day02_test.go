package day02_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2018/day02"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo string
}{
	{
		testInput,
		12,
		"abcde",
	},
	{
		input,
		4693,
		"pebjqsalrdnckzfihvtxysomg",
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := day02.ComputeChecksum(slice.TrimSplit(tc.input))
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		actual := day02.FindCommonCharsBetweenTargetBoxIds(slice.TrimSplit(tc.input))
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
