package day08_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/mp"
	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2022/day08"
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
		expectedPartOne: 21,
		expectedPartTwo: 8,
	},
	{
		input:           input,
		expectedPartOne: 1698,
		expectedPartTwo: 672280,
	},
}

func TestNewGrid(t *testing.T) {
	g := day08.NewGrid(testInput)
	assert.Equals(t, g.Get(0, 0), 3)
	assert.Equals(t, g.Get(1, 1), 5)
	assert.Equals(t, g.Get(2, 2), 3)
	assert.Equals(t, g.Get(4, 3), 9)
}

func TestHidden(t *testing.T) {
	g := day08.NewGrid(testInput)
	// Outer edge is always visible
	assert.Equals(t, day08.IsVisible(g, 0, 0), true)
	assert.Equals(t, day08.IsVisible(g, 3, 4), true)
	assert.Equals(t, day08.IsVisible(g, 4, 4), true)

	assert.Equals(t, day08.IsVisible(g, 1, 1), true)
	assert.Equals(t, day08.IsVisible(g, 2, 1), true)
	assert.Equals(t, day08.IsVisible(g, 3, 1), false)

	assert.Equals(t, day08.IsVisible(g, 1, 2), true)
	assert.Equals(t, day08.IsVisible(g, 2, 2), false)
	assert.Equals(t, day08.IsVisible(g, 3, 2), true)

	assert.Equals(t, day08.IsVisible(g, 2, 3), true)
	assert.Equals(t, day08.IsVisible(g, 3, 3), false)
	assert.Equals(t, day08.IsVisible(g, 1, 3), false)
}

func TestSolvePartOne(t *testing.T) {
	t.Parallel()
	for _, tc := range tcs {
		g := day08.NewGrid(tc.input)
		v := 0
		for i, row := range g.Sparse() {
			v += len(slice.Filter(mp.Keys(row), func(j int) bool { return day08.IsVisible(g, j, i) }))
		}
		assert.Equals(t, v, tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {
	t.Parallel()
	for _, tc := range tcs {
		g := day08.NewGrid(tc.input)
		v := 0
		for i, row := range g.Sparse() {
			for j := range row {
				v = mth.Max([]int{v, day08.ScenicScore(g, j, i)}...)
			}
		}
		assert.Equals(t, v, tc.expectedPartTwo)
	}
}
