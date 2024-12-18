package day04_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2022/day04"
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
		input:           testInput,
		expectedPartOne: 2,
		expectedPartTwo: 4,
	},
	{
		input:           "1-3,3-3",
		expectedPartOne: 1,
		expectedPartTwo: 1,
	},
	{
		input:           "9-82,2-3",
		expectedPartOne: 0,
		expectedPartTwo: 0,
	},
	{
		input:           input,
		expectedPartOne: 459,
		expectedPartTwo: 779,
	},
}

func TestParseInput(t *testing.T) {
	d := day04.New("1-5,6-10")
	assert.Equals(t, d.R1.Start, 1)
	assert.Equals(t, d.R1.End, 5)
	assert.Equals(t, d.R2.Start, 6)
	assert.Equals(t, d.R2.End, 10)
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		numOverlapping := slice.Reduce(strings.Split(strings.TrimSpace(tc.input), "\n"), 0, func(line string, sum int) int {
			rgs := day04.New(line)
			// If R2 is within the span of R1 or R1 is within the span of R2, count this range pair.
			if (rgs.R2.Start >= rgs.R1.Start && rgs.R2.End <= rgs.R1.End) || (rgs.R1.Start >= rgs.R2.Start && rgs.R1.End <= rgs.R2.End) {
				return sum + 1
			}
			return sum
		})
		assert.Equals(t, numOverlapping, tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		numOverlapping := slice.Reduce(strings.Split(strings.TrimSpace(tc.input), "\n"), 0, func(line string, sum int) int {
			rgs := day04.New(line)
			for i := rgs.R1.Start; i <= rgs.R1.End; i++ {
				if i >= rgs.R2.Start && i <= rgs.R2.End {
					return sum + 1
				}
			}
			return sum
		})
		assert.Equals(t, numOverlapping, tc.expectedPartTwo)
	}
}
