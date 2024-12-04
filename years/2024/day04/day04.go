package day04

import (
	"github.com/mleone10/advent-of-code-go/internal/geo/v2"
	"github.com/mleone10/advent-of-code-go/internal/slice"
)

type WordSearch struct {
	grid geo.Space2D[rune]
}

func NewWordSearch(in string) WordSearch {
	w := geo.Space2D[rune]{}

	for i, r := range slice.TrimSplit(in) {
		for j, c := range r {
			w.Set(geo.Location{A: j, B: i}, c)
		}
	}

	return WordSearch{w}
}

func (w WordSearch) NumInstancesXMas() int {
	// Get a list of candidate "XMAS" starting points by finding all 'X's in the grid
	xs := findCandidateStarts(w.grid, 'X')
	// For each candidate, aggregate the number of "XMAS"s found
	return slice.Reduce(xs, 0, func(x geo.Location, ret int) int {
		return ret + evaluateXLoc(x, w.grid)
	})
}

func (w WordSearch) NumInstancesCrossMas() int {
	// Get a list of candidate "X-MAS" starting points by finding all 'A's in the grid
	as := findCandidateStarts(w.grid, 'A')
	// For each candidate, aggregate the number of "X-MAS"s found
	return slice.Reduce(as, 0, func(a geo.Location, ret int) int {
		if evaluateALoc(a, w.grid) {
			ret++
		}
		return ret
	})
}

func findCandidateStarts(grid geo.Space[rune], c rune) []geo.Location {
	return geo.Reduce(grid, []geo.Location{}, func(p geo.Point[rune], ret []geo.Location) []geo.Location {
		if p.Val == c {
			ret = append(ret, p.Loc)
		}
		return ret
	})
}

// From an initial starting 'X', these are the possible coordinates of subsequent 'M', 'A', and 'S' runes
var xMasDeltaMasks = [][]geo.Location{
	{{A: 1}, {A: 2}, {A: 3}},
	{{A: -1}, {A: -2}, {A: -3}},
	{{B: 1}, {B: 2}, {B: 3}},
	{{B: -1}, {B: -2}, {B: -3}},
	{{A: 1, B: 1}, {A: 2, B: 2}, {A: 3, B: 3}},
	{{A: 1, B: -1}, {A: 2, B: -2}, {A: 3, B: -3}},
	{{A: -1, B: 1}, {A: -2, B: 2}, {A: -3, B: 3}},
	{{A: -1, B: -1}, {A: -2, B: -2}, {A: -3, B: -3}},
}

func evaluateXLoc(x geo.Location, g geo.Space[rune]) int {
	// For each possible "XMAS" match delta mask, check the current candidate 'X'
	return slice.Reduce(xMasDeltaMasks, 0, func(d []geo.Location, ret int) int {
		if g.Get(geo.LocSum(x, d[0])) == 'M' &&
			g.Get(geo.LocSum(x, d[1])) == 'A' &&
			g.Get(geo.LocSum(x, d[2])) == 'S' {
			ret += 1
		}
		return ret
	})
}

// From an initial starting 'A', these are the possible coordinates of adjacent 'M' and 'S' runes
var crossMasDeltaMasks = []map[rune][]geo.Location{
	{'M': {{A: 1, B: 1}, {A: 1, B: -1}}, 'S': {{A: -1, B: 1}, {A: -1, B: -1}}},
	{'M': {{A: 1, B: 1}, {A: -1, B: 1}}, 'S': {{A: 1, B: -1}, {A: -1, B: -1}}},
	{'M': {{A: -1, B: 1}, {A: -1, B: -1}}, 'S': {{A: 1, B: 1}, {A: 1, B: -1}}},
	{'M': {{A: -1, B: -1}, {A: 1, B: -1}}, 'S': {{A: 1, B: 1}, {A: -1, B: 1}}},
}

func evaluateALoc(a geo.Location, g geo.Space[rune]) bool {
	// Whereas an 'X' can be the start of multiple "XMAS"s, an 'A' can only be the center of one "X-MAS", so we can loop until we find a match and return
	for _, d := range crossMasDeltaMasks {
		if g.Get(geo.LocSum(a, d['M'][0])) == 'M' &&
			g.Get(geo.LocSum(a, d['M'][1])) == 'M' &&
			g.Get(geo.LocSum(a, d['S'][0])) == 'S' &&
			g.Get(geo.LocSum(a, d['S'][1])) == 'S' {
			return true
		}
	}
	return false
}
