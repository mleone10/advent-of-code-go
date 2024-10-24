package day10_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/src/day10"
)

//go:embed test_input.txt
var testInput string

//go:embed test_output.txt
var testOutput string

//go:embed input.txt
var input string

//go:embed output.txt
var output string

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo string
}{
	{
		input:           testInput,
		expectedPartOne: 13140,
		expectedPartTwo: testOutput,
	},
	{
		input:           input,
		expectedPartOne: 13740,
		expectedPartTwo: output,
	},
}

func TestLoadProgram(t *testing.T) {
	cs := day10.LoadProgram([]string{"noop", "addx 3", "addx -5"})
	assert.ArraysEqual(t, cs.Cycles(), []int{1, 1, 1, 4, 4})
	assert.Equal(t, cs.RegisterX(), -1)
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		cs := day10.LoadProgram(strings.Split(strings.TrimSpace(tc.input), "\n"))
		assert.Equal(t, partOneProductSignalStrength(cs), tc.expectedPartOne)
	}
}

func partOneProductSignalStrength(cs day10.CommSystem) int {
	return cs.SignalStrengthCycleN(20) +
		cs.SignalStrengthCycleN(60) +
		cs.SignalStrengthCycleN(100) +
		cs.SignalStrengthCycleN(140) +
		cs.SignalStrengthCycleN(180) +
		cs.SignalStrengthCycleN(220)
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		cs := day10.LoadProgram(strings.Split(strings.TrimSpace(tc.input), "\n"))
		assert.Equal(t, cs.Render(), strings.TrimSpace(tc.expectedPartTwo))
	}
}
