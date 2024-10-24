package grid_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/grid"
)

func TestPointNeighboring(t *testing.T) {
	p := grid.Point{X: 0, Y: 0}
	q := grid.Point{X: 1, Y: 0}
	assert.Equal(t, p.Neighboring(q), true)

	q.Y = 1
	assert.Equal(t, p.Neighboring(q), true)

	q.Y = 2
	assert.Equal(t, p.Neighboring(q), false)
}
