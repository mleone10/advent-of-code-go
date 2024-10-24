package day12_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/grid"
	"github.com/mleone10/advent-of-code-2022/src/day12"
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
	terrain := day12.NewTerrain(testInput)
	assert.Equal(t, len(terrain.All()), 5)
	assert.Equal(t, len(terrain.All()[0]), 8)
	assert.Equal(t, terrain.Get(0, 0), 0)
	assert.Equal(t, terrain.Get(5, 2), 25)
	assert.Equal(t, terrain.Get(7, 4), 8)
	assert.Equal(t, terrain.Get(2, 3), 2)
	assert.Equal(t, terrain.Start, grid.Point{X: 0, Y: 0})
	assert.Equal(t, terrain.End, grid.Point{X: 5, Y: 2})
}

func TestValidNeighbors(t *testing.T) {
	terrain := day12.NewTerrain(`abc
	def
	ghi`)
	assert.ArraysEqual(t, terrain.ValidNeighbors(grid.Point{X: 1, Y: 1}, []grid.Point{{X: 0, Y: 1}}), []grid.Point{{X: 1, Y: 0}, {X: 2, Y: 1}})
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		terrain := day12.NewTerrain(tc.input)
		assert.Equal(t, terrain.DistanceToEnd(), tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {}
