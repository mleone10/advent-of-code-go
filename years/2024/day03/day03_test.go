package day03_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/years/2024/day03"
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
		161,
		48,
	},
	{
		input,
		179834255,
		80570939,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		in := strings.ReplaceAll(tc.input, "\n", "")
		actual := day03.InterpretCorruptedMemory(in, false)
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		in := strings.ReplaceAll(tc.input, "\n", "")
		actual := day03.InterpretCorruptedMemory(in, true)
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
