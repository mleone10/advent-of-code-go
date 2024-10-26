package day10_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/years/2019/day10"
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
		210,
		802,
	},
	{
		input,
		314,
		1513,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		f := day10.NewField(tc.input)
		assert.Equals(t, tc.expectedPartOne, f.AsteroidsInView())
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		f := day10.NewField(tc.input)
		assert.Equals(t, tc.expectedPartTwo, f.NthDestroyedProduct(200))
	}
}
