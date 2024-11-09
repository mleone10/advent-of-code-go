package geo

import "github.com/mleone10/advent-of-code-go/internal/slice"

type NeighborMode int

const (
	// Cardinal mode only targets direct North/South/East/West adjacent Locations
	NeighborModeCardinal NeighborMode = iota
	// Full mode includes Cardinal neighbors as well as diagonal ones
	NeighborModeFull
)

func Neighbors1D(l Location, m NeighborMode) []Location {
	return neighbors(l, neighborVectors1D)
}

func Neighbors2D(l Location, m NeighborMode) []Location {
	if m == NeighborModeCardinal {
		return neighbors(l, neighborVectors2DCardinal)
	}
	return neighbors(l, neighborVectors2DFull)
}

func Neighbors3D(l Location, m NeighborMode) []Location {
	if m == NeighborModeCardinal {
		return neighbors(l, neighborVectors3DCardinal)
	}
	return neighbors(l, neighborVectors3DFull)
}

func Neighbors4D(l Location, m NeighborMode) []Location {
	if m == NeighborModeCardinal {
		return neighbors(l, neighborVectors4DCardinal)
	}
	return neighbors(l, neighborVectors4DFull)
}

func neighbors(l Location, v []Location) []Location {
	return slice.Map(v, func(vec Location) Location {
		return LocSum(vec, l)
	})
}

var neighborVectors1D = []Location{
	{A: 1},
	{A: -1},
}

var neighborVectors2DCardinal = append(neighborVectors1D, []Location{
	{B: 1},
	{B: -1},
}...)

var neighborVectors2DFull = append(neighborVectors2DCardinal, []Location{
	{A: 1, B: 1},
	{A: 1, B: -1},
	{A: -1, B: 1},
	{A: -1, B: -1},
}...)

var neighborVectors3DCardinal = append(neighborVectors2DCardinal, []Location{
	{C: -1},
	{C: 1},
}...)

var neighborVectors3DFull = append(neighborVectors2DFull, []Location{
	{A: 1, B: 1, C: 1},
	{A: 1, B: 1, C: -1},
	{A: 1, B: -1, C: 1},
	{A: 1, B: -1, C: -1},
	{A: -1, B: 1, C: 1},
	{A: -1, B: 1, C: -1},
	{A: -1, B: -1, C: 1},
	{A: -1, B: -1, C: -1},

	{B: 1, C: 1},
	{B: 1, C: -1},
	{B: -1, C: 1},
	{B: -1, C: -1},

	{A: 1, C: 1},
	{A: 1, C: -1},
	{A: -1, C: 1},
	{A: -1, C: -1},

	{C: 1},
	{C: -1},
}...)

// TODO: populate 4D cardinal neighbor vectors slice
var neighborVectors4DCardinal = append(neighborVectors3DCardinal, []Location{}...)

// TODO: populate 4D full neighbor vectors slice
var neighborVectors4DFull = append(neighborVectors3DFull, []Location{}...)
