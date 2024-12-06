package day05_test

import (
	_ "embed"
	"sort"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2024/day05"
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
		143,
		123,
	},
	{
		input,
		4135,
		5285,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		r := day05.NewRules(strings.Split(tc.input, "\n\n")[0])
		ps := day05.NewPages(strings.Split(tc.input, "\n\n")[1], r)

		actual := slice.Reduce(ps, 0, func(p day05.Page, ret int) int {
			if sort.IsSorted(p) {
				ret += p.GetMiddlePage()
			}
			return ret
		})

		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		r := day05.NewRules(strings.Split(tc.input, "\n\n")[0])
		ps := day05.NewPages(strings.Split(tc.input, "\n\n")[1], r)

		actual := slice.Reduce(
			slice.Filter(ps, func(p day05.Page) bool {
				return !sort.IsSorted(p)
			}), 0, func(p day05.Page, ret int) int {
				sort.Sort(p)
				return ret + p.GetMiddlePage()
			})

		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
