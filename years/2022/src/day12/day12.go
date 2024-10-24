package day12

import (
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/array"
	"github.com/mleone10/advent-of-code-2022/pkg/grid"
	"github.com/mleone10/advent-of-code-2022/pkg/queue"
)

type Terrain struct {
	grid.Plane[int]
	Start, End  grid.Point
	distanceMap grid.Plane[int]
}

func NewTerrain(input string) *Terrain {
	t := Terrain{}

	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, h := range row {
			if h == 'S' {
				t.Start = grid.Point{X: x, Y: y}
				t.Set(x, y, heightOf('a'))
			} else if h == 'E' {
				t.End = grid.Point{X: x, Y: y}
				t.Set(x, y, heightOf('z'))
			} else {
				t.Set(x, y, int(h)-97)
			}
		}
	}

	t.populateDistanceMap()

	return &t
}

func heightOf(r rune) int {
	return int(r) - 97
}

func (t Terrain) populateDistanceMap() {
	// TODO: implement breadth-first search to construct map of every point's shortest distance from t.Start.
	var q queue.Queue[grid.Point]
	distanceMap := grid.Plane[int]{}
	distanceMap.Set(t.Start.X, t.Start.Y, 0)
	q.Push(t.Start)

	for q.Length() != 0 {

	}

}

func (t Terrain) ValidNeighbors(p grid.Point, path []grid.Point) []grid.Point {
	var neighbors []grid.Point

	for _, delta := range []grid.Point{{X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0}} {
		candidate := grid.Point{X: p.X + delta.X, Y: p.Y + delta.Y}
		if t.Has(candidate.X, candidate.Y) && !array.Contains(path, candidate) {
			candidateHeight := t.Get(candidate.X, candidate.Y)
			currentHeight := t.Get(p.X, p.Y)
			if candidateHeight <= currentHeight+1 {
				neighbors = append(neighbors, candidate)
			}
		}
	}

	return neighbors
}

func (t Terrain) DistanceToEnd() int {
	return t.distanceMap.Get(t.End.X, t.End.Y)
}
