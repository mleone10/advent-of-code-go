package day17

import (
	"github.com/mleone10/advent-of-code-go/internal/geo/v2"
)

type Day17 struct {
	Space      geo.Space[bool]
	Dimensions geo.Dimension
}

func (d Day17) NumActiveAfterN(n int) int {
	for range n {
		d.step()
	}

	return geo.Reduce(d.Space, 0, func(p geo.Point[bool], acc int) int {
		if p.Val {
			acc++
		}
		return acc
	})
}

func (d *Day17) step() {
	// Compile a frequency list of all neighbors of currently active points
	allNeighbors := geo.Reduce(d.Space, map[geo.Point[bool]]int{}, func(p geo.Point[bool], acc map[geo.Point[bool]]int) map[geo.Point[bool]]int {
		if p.Val {
			for _, n := range geo.Neighbors(p.Loc, geo.NeighborModeFull, d.Dimensions) {
				acc[geo.Point[bool]{Val: d.Space.Get(n), Loc: n}]++
			}
		}
		return acc
	})

	// Use the frequency list to determine the next iteration's active points
	next := geo.Space4D[bool]{}
	for loc, n := range allNeighbors {
		if (loc.Val && (n == 2 || n == 3)) || (!loc.Val && n == 3) {
			next.Set(loc.Loc, true)
		}
	}

	d.Space = next
}
