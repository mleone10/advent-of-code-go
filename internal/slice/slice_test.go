package slice_test

import (
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/slice"
)

var (
	strLength = func(x string) int { return len(x) }
	intDouble = func(x int) int { return 2 * x }
	isEven    = func(x int) bool { return x%2 == 0 }
	isUpper   = func(x string) bool { return x == strings.ToUpper(x) }
)

func TestMap(t *testing.T) {
	assert.ArrayEquals(t, slice.Map([]string{"a", "b", "c"}, strings.ToUpper), []string{"A", "B", "C"})
	assert.ArrayEquals(t, slice.Map([]string{"fizz", "buzz", "hello"}, strLength), []int{4, 4, 5})
	assert.ArrayEquals(t, slice.Map([]int{2, 3, 4}, intDouble), []int{4, 6, 8})
}

func TestFilter(t *testing.T) {
	assert.ArrayEquals(t, slice.Filter([]int{1, 2, 3, 4, 5}, isEven), []int{2, 4})
	assert.ArrayEquals(t, slice.Filter([]string{"FIZZ", "bUZz", "foO", "BAR"}, isUpper), []string{"FIZZ", "BAR"})
}

func TestTrimSplit(t *testing.T) {
	input := `
	foo
	bar
	fizz
	buzz
	`
	assert.ArrayEquals(t, []string{"foo", "bar", "fizz", "buzz"}, slice.TrimSplit(input))
}

func TestSum(t *testing.T) {
	assert.Equals(t, 15, slice.Sum([]int{1, 2, 3, 4, 5}))
}

func TestContains(t *testing.T) {
	tcs := []struct {
		subj     []int
		v        int
		expected bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2, true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			6, false,
		},
	}

	for _, tc := range tcs {
		assert.Equals(t, tc.expected, slice.Contains(tc.subj, tc.v))
	}
}

func TestReverse(t *testing.T) {
	subj := []int{1, 2, 3, 4, 5}

	actual := slice.Reverse(subj)

	assert.ArrayEquals(t, []int{5, 4, 3, 2, 1}, actual)
}

func TestTake(t *testing.T) {
	assert.ArrayEquals(t, slice.Take([]int{1, 2, 3, 4, 5}, 3), []int{1, 2, 3})
	assert.ArrayEquals(t, slice.Take([]int{1, 2, 3}, 3), []int{1, 2, 3})
	assert.ArrayEquals(t, slice.Take([]int{1, 2}, 3), []int{1, 2})
	assert.ArrayEquals(t, slice.Take([]string{"foo", "bar", "fizz", "buzz"}, 3), []string{"foo", "bar", "fizz"})
}

func TestFrequencyList(t *testing.T) {
	testIntFreqs := slice.FrequencyList([]int{5, 5, 10, 10, 10, 15})
	assert.Equals(t, testIntFreqs[5], 2)
	assert.Equals(t, testIntFreqs[10], 3)
	assert.Equals(t, testIntFreqs[15], 1)

	testFloatFreqs := slice.FrequencyList([]float32{1.1, 1.1, 2.2, 2.2, 2.2, 3.3})
	assert.Equals(t, testFloatFreqs[1.1], 2)
	assert.Equals(t, testFloatFreqs[2.2], 3)
	assert.Equals(t, testFloatFreqs[3.3], 1)
}
