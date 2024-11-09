package geo_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/geo/v2"
)

func TestLocSum(t *testing.T) {
	actual := geo.LocSum(geo.Location{A: 1, B: 2, C: 3, D: 4}, geo.Location{A: 1, B: 2, C: 3, D: 4})

	assert.Equals(t, actual, geo.Location{A: 2, B: 4, C: 6, D: 8})
}
