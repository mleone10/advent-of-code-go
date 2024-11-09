package day17_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/geo/v2"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2020/day17"
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
		112,
		848,
	},
	{
		input,
		382,
		2552,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		d := parseInput(tc.input, false)
		assert.Equals(t, tc.expectedPartOne, d.NumActiveAfterN(6))
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		d := parseInput(tc.input, true)
		assert.Equals(t, tc.expectedPartTwo, d.NumActiveAfterN(6))
	}
}

func parseInput(in string, p2 bool) day17.Day17 {
	d := day17.Day17{geo.Space4D[bool]{}, 3}
	if p2 {
		d.Dimensions = 4
	}

	for i, r := range slice.TrimSplit(in) {
		for j, c := range r {
			if c == '#' {
				d.Space.Set(geo.Location{A: j, B: i}, true)
			}
		}
	}

	return d
}
