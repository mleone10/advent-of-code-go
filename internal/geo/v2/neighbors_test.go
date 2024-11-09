package geo_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/geo/v2"
)

func TestNeighbors1DCardinal(t *testing.T) {
	testNeighbors1D(t, geo.NeighborModeCardinal)
}

func TestNeighbors1DFull(t *testing.T) {
	testNeighbors1D(t, geo.NeighborModeFull)
}

// Cardinal and Full neighbor modes are identical for 1D Spaces
func testNeighbors1D(t *testing.T, n geo.NeighborMode) {
	subj := geo.Location{A: 5}

	actual := geo.Neighbors1D(subj, n)

	assert.Equals(t, 2, len(actual))
	assert.Contains(t, actual, geo.Location{A: 4})
	assert.Contains(t, actual, geo.Location{A: 6})
}

func TestNeighbors2DCardinal(t *testing.T) {
	subj := geo.Location{A: 5, B: 3}

	actual := geo.Neighbors2D(subj, geo.NeighborModeCardinal)

	assert.Equals(t, 4, len(actual))
	assert.Contains(t, actual, geo.Location{A: 4, B: 3})
	assert.Contains(t, actual, geo.Location{A: 6, B: 3})
	assert.Contains(t, actual, geo.Location{A: 5, B: 2})
	assert.Contains(t, actual, geo.Location{A: 5, B: 4})
}

func TestNeighbors2DFull(t *testing.T) {
	subj := geo.Location{A: 5, B: 3}

	actual := geo.Neighbors2D(subj, geo.NeighborModeFull)

	assert.Equals(t, 8, len(actual))
	assert.Contains(t, actual, geo.Location{A: 4, B: 3})
	assert.Contains(t, actual, geo.Location{A: 6, B: 3})
	assert.Contains(t, actual, geo.Location{A: 5, B: 2})
	assert.Contains(t, actual, geo.Location{A: 5, B: 4})
	assert.Contains(t, actual, geo.Location{A: 4, B: 2})
	assert.Contains(t, actual, geo.Location{A: 6, B: 2})
	assert.Contains(t, actual, geo.Location{A: 4, B: 4})
	assert.Contains(t, actual, geo.Location{A: 6, B: 4})
}

func TestNeighbors3DCardinal(t *testing.T) {
	subj := geo.Location{A: 5, B: 3, C: 1}

	actual := geo.Neighbors3D(subj, geo.NeighborModeCardinal)

	assert.Equals(t, 6, len(actual))
	assert.Contains(t, actual, geo.Location{A: 4, B: 3, C: 1})
	assert.Contains(t, actual, geo.Location{A: 6, B: 3, C: 1})
	assert.Contains(t, actual, geo.Location{A: 5, B: 2, C: 1})
	assert.Contains(t, actual, geo.Location{A: 5, B: 4, C: 1})
	assert.Contains(t, actual, geo.Location{A: 5, B: 3, C: 0})
	assert.Contains(t, actual, geo.Location{A: 5, B: 3, C: 2})
}

func TestNeighbors3DFull(t *testing.T) {
	subj := geo.Location{A: 5, B: 3, C: 1}

	actual := geo.Neighbors3D(subj, geo.NeighborModeFull)

	assert.Equals(t, 26, len(actual))
	assert.Contains(t, actual, geo.Location{A: 4, B: 3, C: 1})
	assert.Contains(t, actual, geo.Location{A: 4, B: 2, C: 1})
	assert.Contains(t, actual, geo.Location{A: 5, B: 2, C: 1})
	assert.Contains(t, actual, geo.Location{A: 6, B: 2, C: 1})
	assert.Contains(t, actual, geo.Location{A: 6, B: 3, C: 1})
	assert.Contains(t, actual, geo.Location{A: 6, B: 4, C: 1})
	assert.Contains(t, actual, geo.Location{A: 5, B: 4, C: 1})
	assert.Contains(t, actual, geo.Location{A: 4, B: 4, C: 1})

	assert.Contains(t, actual, geo.Location{A: 4, B: 3, C: 0})
	assert.Contains(t, actual, geo.Location{A: 4, B: 2, C: 0})
	assert.Contains(t, actual, geo.Location{A: 5, B: 2, C: 0})
	assert.Contains(t, actual, geo.Location{A: 6, B: 2, C: 0})
	assert.Contains(t, actual, geo.Location{A: 6, B: 3, C: 0})
	assert.Contains(t, actual, geo.Location{A: 6, B: 4, C: 0})
	assert.Contains(t, actual, geo.Location{A: 5, B: 4, C: 0})
	assert.Contains(t, actual, geo.Location{A: 4, B: 4, C: 0})
	assert.Contains(t, actual, geo.Location{A: 5, B: 4, C: 0})

	assert.Contains(t, actual, geo.Location{A: 4, B: 3, C: 2})
	assert.Contains(t, actual, geo.Location{A: 4, B: 2, C: 2})
	assert.Contains(t, actual, geo.Location{A: 5, B: 2, C: 2})
	assert.Contains(t, actual, geo.Location{A: 6, B: 2, C: 2})
	assert.Contains(t, actual, geo.Location{A: 6, B: 3, C: 2})
	assert.Contains(t, actual, geo.Location{A: 6, B: 4, C: 2})
	assert.Contains(t, actual, geo.Location{A: 5, B: 4, C: 2})
	assert.Contains(t, actual, geo.Location{A: 4, B: 4, C: 2})
	assert.Contains(t, actual, geo.Location{A: 5, B: 4, C: 2})
}
