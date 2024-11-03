package day16_test

import (
	_ "embed"
	"regexp"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2020/day16"
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
		71,
		1,
	},
	{
		input,
		27802,
		279139880759,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		d := parseInput(tc.input)
		assert.Equals(t, tc.expectedPartOne, d.CalcErrorRate())
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		d := parseInput(tc.input)
		assert.Equals(t, tc.expectedPartTwo, d.CalcDepartureProduct())
	}
}

func parseInput(input string) day16.Day16 {
	d := day16.Day16{}

	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	fieldsRegExp := regexp.MustCompile(`([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)`)
	for _, l := range slice.TrimSplit(parts[0]) {
		m := fieldsRegExp.FindStringSubmatch(l)
		d.AddField(m[1],
			mth.Atoi(m[2]),
			mth.Atoi(m[3]),
			mth.Atoi(m[4]),
			mth.Atoi(m[5]),
		)
	}

	fs := []int{}
	for _, f := range strings.Split(slice.TrimSplit(parts[1])[1], ",") {
		fs = append(fs, mth.Atoi(f))
	}
	d.AddMyTicket(fs...)

	for _, t := range slice.TrimSplit(parts[2])[1:] {
		fs = []int{}
		for _, f := range strings.Split(t, ",") {
			fs = append(fs, mth.Atoi(f))
		}
		d.AddNearbyTicket(fs...)
	}

	return d
}
