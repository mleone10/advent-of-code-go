package day06_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/src/day06"
)

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo int
}{
	{
		input:           "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		expectedPartOne: 7,
		expectedPartTwo: 19,
	},
	{
		input:           "bvwbjplbgvbhsrlpgdmjqwftvncz",
		expectedPartOne: 5,
		expectedPartTwo: 23,
	},
	{
		input:           "nppdvjthqldpwncqszvftbrmjlhg",
		expectedPartOne: 6,
		expectedPartTwo: 23,
	},
	{
		input:           "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		expectedPartOne: 10,
		expectedPartTwo: 29,
	},
	{
		input:           "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		expectedPartOne: 11,
		expectedPartTwo: 26,
	},
	{
		input:           input,
		expectedPartOne: 1794,
		expectedPartTwo: 2851,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, day06.FindStartOfPacket(tc.input), tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, day06.FindStartOfMessage(tc.input), tc.expectedPartTwo)
	}
}
