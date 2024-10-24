package grid

type Point struct {
	X int
	Y int
}

type Vector struct {
	dir Point
	mag int
}

type UnitVector Vector

var (
	UnitVectorNorth = UnitVector{Point{0, -1}, 1}
	UnitVectorSouth = UnitVector{Point{0, 1}, 1}
	UnitVectorWest  = UnitVector{Point{-1, 0}, 1}
	UnitVectorEast  = UnitVector{Point{1, 0}, 1}
)

func (p Point) Neighboring(q Point) bool {
	return abs(q.X-p.X) <= 1 && abs(q.Y-p.Y) <= 1
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
