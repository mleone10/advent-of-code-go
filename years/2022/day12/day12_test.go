package day12_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/years/2022/day12"
	"github.com/mleone10/advent-of-code-go/years/2022/pkg/grid"
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
		expectedPartOne: 31,
	},
	{
		input:           input,
		expectedPartOne: 0,
	},
}

func TestNewTerrain(t *testing.T) {
	t.SkipNow()
	terrain := day12.NewTerrain(testInput)
	assert.Equals(t, len(terrain.All()), 5)
	assert.Equals(t, len(terrain.All()[0]), 8)
	assert.Equals(t, terrain.Get(0, 0), 0)
	assert.Equals(t, terrain.Get(5, 2), 25)
	assert.Equals(t, terrain.Get(7, 4), 8)
	assert.Equals(t, terrain.Get(2, 3), 2)
	assert.Equals(t, terrain.Start, grid.Point{X: 0, Y: 0})
	assert.Equals(t, terrain.End, grid.Point{X: 5, Y: 2})
}

func TestSolvePartOne(t *testing.T) {
	t.SkipNow()
	for _, tc := range tcs {
		terrain := day12.NewTerrain(tc.input)
		assert.Equals(t, terrain.DistanceToEnd(), tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {}
