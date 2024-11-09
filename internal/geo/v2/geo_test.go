package geo_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/geo/v2"
)

func funcMapSquareInt(p geo.Point[int]) geo.Point[int] {
	return geo.Point[int]{p.Loc, p.Val * p.Val}
}

func funcFilterOdd(p geo.Point[int]) bool {
	return p.Val%2 == 0
}

func funcReduceSum(p geo.Point[int], acc int) int {
	return acc + p.Val
}

func TestMap1D(t *testing.T) {
	subj := make(geo.Space1D[int])
	subj.Set(geo.Location{A: 1}, 2)
	subj.Set(geo.Location{A: 2}, 4)
	subj.Set(geo.Location{A: 3}, 6)

	actual := geo.Map1D[int, int](subj, funcMapSquareInt)

	assert.Equals(t, 4, actual.Get(geo.Location{A: 1}))
	assert.Equals(t, 16, actual.Get(geo.Location{A: 2}))
	assert.Equals(t, 36, actual.Get(geo.Location{A: 3}))
}

func TestMap2D(t *testing.T) {
	subj := make(geo.Space2D[int])
	subj.Set(geo.Location{A: 1, B: 2}, 2)
	subj.Set(geo.Location{A: 2, B: 4}, 4)
	subj.Set(geo.Location{A: 3, B: 8}, 6)

	actual := geo.Map2D[int, int](subj, funcMapSquareInt)

	assert.Equals(t, 4, actual.Get(geo.Location{A: 1, B: 2}))
	assert.Equals(t, 16, actual.Get(geo.Location{A: 2, B: 4}))
	assert.Equals(t, 36, actual.Get(geo.Location{A: 3, B: 8}))
}

func TestMap3D(t *testing.T) {
	subj := make(geo.Space3D[int])
	subj.Set(geo.Location{A: 1, B: 2, C: 3}, 2)
	subj.Set(geo.Location{A: 2, B: 4, C: 6}, 4)
	subj.Set(geo.Location{A: 3, B: 8, C: 9}, 6)

	actual := geo.Map3D[int, int](subj, funcMapSquareInt)

	assert.Equals(t, 4, actual.Get(geo.Location{A: 1, B: 2, C: 3}))
	assert.Equals(t, 16, actual.Get(geo.Location{A: 2, B: 4, C: 6}))
	assert.Equals(t, 36, actual.Get(geo.Location{A: 3, B: 8, C: 9}))
}

func TestMap4D(t *testing.T) {
	subj := make(geo.Space4D[int])
	subj.Set(geo.Location{A: 1, B: 2, C: 3, D: 4}, 2)
	subj.Set(geo.Location{A: 2, B: 4, C: 6, D: 8}, 4)
	subj.Set(geo.Location{A: 3, B: 8, C: 9, D: 16}, 6)

	actual := geo.Map4D[int, int](subj, funcMapSquareInt)

	assert.Equals(t, 4, actual.Get(geo.Location{A: 1, B: 2, C: 3, D: 4}))
	assert.Equals(t, 16, actual.Get(geo.Location{A: 2, B: 4, C: 6, D: 8}))
	assert.Equals(t, 36, actual.Get(geo.Location{A: 3, B: 8, C: 9, D: 16}))
}

func TestFilter1D(t *testing.T) {
	subj := make(geo.Space1D[int])
	subj.Set(geo.Location{A: 1}, 1)
	subj.Set(geo.Location{A: 2}, 2)
	subj.Set(geo.Location{A: 3}, 3)
	subj.Set(geo.Location{A: 4}, 4)

	actual := geo.Filter1D[int](subj, funcFilterOdd)

	assert.Equals(t, actual.Size(), 2)
	assert.Equals(t, 2, actual.Get(geo.Location{A: 2}))
	assert.Equals(t, 4, actual.Get(geo.Location{A: 4}))
}

func TestFilter2D(t *testing.T) {
	subj := make(geo.Space2D[int])
	subj.Set(geo.Location{A: 1, B: 2}, 1)
	subj.Set(geo.Location{A: 2, B: 4}, 2)
	subj.Set(geo.Location{A: 3, B: 8}, 3)
	subj.Set(geo.Location{A: 4, B: 16}, 4)

	actual := geo.Filter2D[int](subj, funcFilterOdd)

	assert.Equals(t, actual.Size(), 2)
	assert.Equals(t, 2, actual.Get(geo.Location{A: 2, B: 4}))
	assert.Equals(t, 4, actual.Get(geo.Location{A: 4, B: 16}))
}

func TestFilter3D(t *testing.T) {
	subj := make(geo.Space3D[int])
	subj.Set(geo.Location{A: 1, B: 2, C: 3}, 1)
	subj.Set(geo.Location{A: 2, B: 4, C: 6}, 2)
	subj.Set(geo.Location{A: 3, B: 8, C: 9}, 3)
	subj.Set(geo.Location{A: 4, B: 16, C: 81}, 4)

	actual := geo.Filter3D[int](subj, funcFilterOdd)

	assert.Equals(t, actual.Size(), 2)
	assert.Equals(t, 2, actual.Get(geo.Location{A: 2, B: 4, C: 6}))
	assert.Equals(t, 4, actual.Get(geo.Location{A: 4, B: 16, C: 81}))
}

func TestFilter4D(t *testing.T) {
	subj := make(geo.Space4D[int])
	subj.Set(geo.Location{A: 1, B: 2, C: 3, D: 4}, 1)
	subj.Set(geo.Location{A: 2, B: 4, C: 6, D: 8}, 2)
	subj.Set(geo.Location{A: 3, B: 8, C: 9, D: 16}, 3)
	subj.Set(geo.Location{A: 4, B: 16, C: 81, D: 256}, 4)

	actual := geo.Filter4D[int](subj, funcFilterOdd)

	assert.Equals(t, actual.Size(), 2)
	assert.Equals(t, 2, actual.Get(geo.Location{A: 2, B: 4, C: 6, D: 8}))
	assert.Equals(t, 4, actual.Get(geo.Location{A: 4, B: 16, C: 81, D: 256}))
}

func TestReduce1D(t *testing.T) {
	subj := make(geo.Space1D[int])
	subj.Set(geo.Location{A: 1}, 1)
	subj.Set(geo.Location{A: 2}, 2)
	subj.Set(geo.Location{A: 3}, 3)

	actual := geo.Reduce[int](subj, 0, funcReduceSum)

	assert.Equals(t, 6, actual)
}

func TestReduce2D(t *testing.T) {
	subj := make(geo.Space2D[int])
	subj.Set(geo.Location{A: 1, B: 2}, 2)
	subj.Set(geo.Location{A: 2, B: 4}, 4)
	subj.Set(geo.Location{A: 3, B: 8}, 6)

	actual := geo.Reduce[int](subj, 0, funcReduceSum)

	assert.Equals(t, 12, actual)
}

func TestReduce3D(t *testing.T) {
	subj := make(geo.Space3D[int])
	subj.Set(geo.Location{A: 1, B: 2, C: 3}, 2)
	subj.Set(geo.Location{A: 2, B: 4, C: 6}, 4)
	subj.Set(geo.Location{A: 3, B: 8, C: 9}, 6)

	actual := geo.Reduce[int](subj, 0, funcReduceSum)

	assert.Equals(t, 12, actual)
}

func TestReduce4D(t *testing.T) {
	subj := make(geo.Space4D[int])
	subj.Set(geo.Location{A: 1, B: 2, C: 3, D: 4}, 2)
	subj.Set(geo.Location{A: 2, B: 4, C: 6, D: 8}, 4)
	subj.Set(geo.Location{A: 3, B: 8, C: 9, D: 16}, 6)

	actual := geo.Reduce[int](subj, 0, funcReduceSum)

	assert.Equals(t, 12, actual)
}
