package mathutil_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/mathutil"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, mathutil.Abs(1), 1)
	assert.Equal(t, mathutil.Abs(-1), 1)
}
