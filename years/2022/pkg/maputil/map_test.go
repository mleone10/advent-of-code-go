package maputil_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/maputil"
)

func TestKeys(t *testing.T) {
	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	keys := maputil.Keys(testMap)

	assert.Contains(t, keys, "one")
	assert.Contains(t, keys, "two")
	assert.Contains(t, keys, "three")
	assert.Equal(t, len(keys), 3)
}

func TestValues(t *testing.T) {
	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	values := maputil.Values(testMap)

	assert.Contains(t, values, 1)
	assert.Contains(t, values, 2)
	assert.Contains(t, values, 3)
	assert.Equal(t, len(values), 3)
}
