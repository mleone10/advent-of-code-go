package geo

import "fmt"

type NeighborMode int

const (
	// Cardinal mode only targets direct North/South/East/West adjacent Locations
	NeighborModeCardinal NeighborMode = iota
	// Full mode includes Cardinal neighbors as well as diagonal ones
	NeighborModeFull
)

func Neighbors(l Location, m NeighborMode, d Dimension) []Location {
	switch d {
	case D1:
		return neighbors1D(l)
	case D2:
		return neighbors2D(l, m)
	case D3:
		return neighbors3D(l, m)
	case D4:
		return neighbors4D(l, m)
	default:
		panic(fmt.Sprintf("invalid dimension provided to geo.Neighbors: %v", d))
	}
}

func neighbors1D(l Location) []Location {
	return []Location{{A: l.A - 1}, {A: l.A + 1}}
}

func neighbors2D(l Location, m NeighborMode) []Location {
	ret := []Location{}

	if m == NeighborModeCardinal {
		for i := -1; i <= 1; i++ {
			if i != 0 {
				ret = append(ret, Location{A: l.A + i, B: l.B})
			}
		}
		for i := -1; i <= 1; i++ {
			if i != 0 {
				ret = append(ret, Location{A: l.A, B: l.B + i})
			}
		}
		return ret
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				ret = append(ret, Location{A: l.A + i, B: l.B + j})
			}
		}
	}

	return ret
}

func neighbors3D(l Location, m NeighborMode) []Location {
	ret := []Location{}

	if m == NeighborModeCardinal {
		for i := -1; i <= 1; i++ {
			if i != 0 {
				ret = append(ret, Location{A: l.A + i, B: l.B, C: l.C})
			}
		}
		for i := -1; i <= 1; i++ {
			if i != 0 {
				ret = append(ret, Location{A: l.A, B: l.B + i, C: l.C})
			}
		}
		for i := -1; i <= 1; i++ {
			if i != 0 {
				ret = append(ret, Location{A: l.A, B: l.B, C: l.C + i})
			}
		}
		return ret
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if !(i == 0 && j == 0 && k == 0) {
					ret = append(ret, Location{A: l.A + i, B: l.B + j, C: l.C + k})
				}
			}
		}
	}

	return ret
}

// TODO: test geo.neighbors4D
func neighbors4D(l Location, m NeighborMode) []Location {
	ret := []Location{}

	if m == NeighborModeCardinal {
		for i := -1; i <= 1; i++ {
			if i != 0 {
				ret = append(ret, Location{A: l.A + i, B: l.B, C: l.C, D: l.D})
			}
		}
		for i := -1; i <= 1; i++ {
			if i != 0 {
				ret = append(ret, Location{A: l.A, B: l.B + i, C: l.C, D: l.D})
			}
		}
		for i := -1; i <= 1; i++ {
			if i != 0 {
				ret = append(ret, Location{A: l.A, B: l.B, C: l.C + i, D: l.D})
			}
		}
		for i := -1; i <= 1; i++ {
			if i != 0 {
				ret = append(ret, Location{A: l.A, B: l.B, C: l.C, D: l.D + i})
			}
		}
		return ret
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for h := -1; h <= 1; h++ {
					if !(i == 0 && j == 0 && k == 0 && h == 0) {
						ret = append(ret, Location{A: l.A + i, B: l.B + j, C: l.C + k, D: l.D + h})
					}
				}
			}
		}
	}

	return ret
}
