package grid_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/grid"
)

func TestSetGet(t *testing.T) {
	var g grid.Plane[int]
	g.Set(0, 0, 1)
	assert.Equal(t, g.Get(0, 0), 1)
	g.Set(100, 100, 1)
	assert.Equal(t, g.Get(100, 100), 1)
}

func TestHas(t *testing.T) {
	var g grid.Plane[int]
	assert.Equal(t, g.Has(0, 0), false)
	g.Set(0, 0, 1)
	assert.Equal(t, g.Has(0, 0), true)
}

func TestWidth(t *testing.T) {
	var g grid.Plane[int]
	g.Set(0, 0, 1)
	g.Set(5, 1, 1)
	g.Set(10, 2, 1)
	assert.Equal(t, g.Width(), 11)
}

func TestHeight(t *testing.T) {
	var g grid.Plane[int]
	g.Set(0, 0, 1)
	g.Set(5, 1, 1)
	g.Set(10, 2, 1)
	assert.Equal(t, g.Height(), 3)
}

func TestRow(t *testing.T) {
	var g grid.Plane[int]
	g.Set(0, 0, 1)
	g.Set(5, 1, 1)
	g.Set(10, 2, 1)
	assert.ArraysEqual(t, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, g.Row(0))
	assert.ArraysEqual(t, []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, g.Row(1))
	assert.ArraysEqual(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, g.Row(2))
}

func TestCol(t *testing.T) {
	var g grid.Plane[int]
	g.Set(0, 0, 1)
	g.Set(5, 1, 1)
	g.Set(10, 2, 1)
	assert.ArraysEqual(t, []int{1, 0, 0}, g.Col(0))
	assert.ArraysEqual(t, []int{0, 1, 0}, g.Col(5))
	assert.ArraysEqual(t, []int{0, 0, 1}, g.Col(10))
}

func TestAll(t *testing.T) {
	var g grid.Plane[int]
	g.Set(0, 0, 1)
	g.Set(5, 1, 1)
	g.Set(10, 2, 1)
	grid := g.All()
	assert.ArraysEqual(t, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, grid[0])
	assert.ArraysEqual(t, []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, grid[1])
	assert.ArraysEqual(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, grid[2])
}
