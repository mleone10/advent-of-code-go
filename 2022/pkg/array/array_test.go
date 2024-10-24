package array_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/array"
	"github.com/mleone10/advent-of-code-2022/pkg/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, array.Max([]int{5, 10, 15}), 15)
	assert.Equal(t, array.Max([]float32{1.1, 2.2, 3.3}), 3.3)
}

func TestMin(t *testing.T) {
	assert.Equal(t, array.Min([]int{5, 10, 15}), 5)
	assert.Equal(t, array.Min([]float32{1.1, 2.2, 3.3}), 1.1)
}

func TestSum(t *testing.T) {
	assert.Equal(t, array.Sum([]int{5, 10, 15}), 30)
	assert.Equal(t, array.Sum([]float64{1.1, 2.2, 3.3}), 6.6)
}

func TestTake(t *testing.T) {
	assert.ArraysEqual(t, array.Take([]int{1, 2, 3, 4, 5}, 3), []int{1, 2, 3})
	assert.ArraysEqual(t, array.Take([]int{1, 2, 3}, 3), []int{1, 2, 3})
	assert.ArraysEqual(t, array.Take([]int{1, 2}, 3), []int{1, 2})
	assert.ArraysEqual(t, array.Take([]string{"foo", "bar", "fizz", "buzz"}, 3), []string{"foo", "bar", "fizz"})
}

func TestFrequencyList(t *testing.T) {
	testIntFreqs := array.FrequencyList([]int{5, 5, 10, 10, 10, 15})
	assert.Equal(t, testIntFreqs[5], 2)
	assert.Equal(t, testIntFreqs[10], 3)
	assert.Equal(t, testIntFreqs[15], 1)

	testFloatFreqs := array.FrequencyList([]float32{1.1, 1.1, 2.2, 2.2, 2.2, 3.3})
	assert.Equal(t, testFloatFreqs[1.1], 2)
	assert.Equal(t, testFloatFreqs[2.2], 3)
	assert.Equal(t, testFloatFreqs[3.3], 1)
}

func TestSlidingSum(t *testing.T) {
	assert.ArraysEqual(t, array.SlidingSum(2, []int{1, 2, 3, 4, 5, 6}), []int{3, 5, 7, 9, 11})
	assert.ArraysEqual(t, array.SlidingSum(4, []int{1, 2, 3, 4, 5, 6}), []int{10, 14, 18})
	assert.ArraysEqual(t, array.SlidingSum(3, []int{1, 2}), []int{})
}

func TestMap(t *testing.T) {
	assert.ArraysEqual(t, array.Map([]int{1, 2, 3, 4, 5}, func(i int) int { return i * i }), []int{1, 4, 9, 16, 25})
	assert.ArraysEqual(t, array.Map([]string{"foo", "bar", "fizz", "buzz"}, func(i string) int { return len(i) }), []int{3, 3, 4, 4})
}

func TestReduce(t *testing.T) {
	assert.Equal(t, array.Reduce([]int{1, 2, 3, 4, 5}, func(acc, v int) int { return acc + v }, 0), 15)
	assert.Equal(t, array.Reduce([]int{1, 2, 3, 4, 5}, func(acc, v int) int {
		// Recreating `max(int, int)`
		if acc > v {
			return acc
		}
		return v
	}, 0), 5)
}

func TestFilter(t *testing.T) {
	assert.ArraysEqual(t, array.Filter([]int{1, 2, 3, 4, 5, 6}, func(i int) bool { return i%2 == 0 }), []int{2, 4, 6})
	assert.ArraysEqual(t, array.Filter([]string{"foo", "bar", "fizz", "buzz"}, func(s string) bool { return s[0] == 'f' }), []string{"foo", "fizz"})
}

func TestContains(t *testing.T) {
	assert.Equal(t, array.Contains([]int{1, 2, 3, 4, 5}, 1), true)
	assert.Equal(t, array.Contains([]int{1, 2, 3, 4, 5}, 6), false)
	assert.Equal(t, array.Contains([]string{"foo", "bar", "fizz", "buzz"}, "foo"), true)
	assert.Equal(t, array.Contains([]string{"foo", "bar", "fizz", "buzz"}, "quix"), false)
}

func TestContainsSubarray(t *testing.T) {
	assert.Equal(t, array.ContainsSubarray([]int{1, 2, 3, 4}, []int{2, 3}), true)
	assert.Equal(t, array.ContainsSubarray([]int{1, 2, 3}, []int{3}), true)
	assert.Equal(t, array.ContainsSubarray([]int{1, 2, 3}, []int{4, 5, 6}), false)
}

func TestEqual(t *testing.T) {
	assert.Equal(t, array.Equal([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}), true)
	assert.Equal(t, array.Equal([]int{1, 2, 3, 4}, []int{1, 2, 3}), false)
	assert.Equal(t, array.Equal([]int{1, 2, 3, 4}, []int{4, 3, 2, 1}), false)
	assert.Equal(t, array.Equal([]string{"foo", "bar"}, []string{"foo", "bar"}), true)
}

func TestReverse(t *testing.T) {
	assert.ArraysEqual(t, array.Reverse([]int{1, 2, 3, 4, 5, 6}), []int{6, 5, 4, 3, 2, 1})
	assert.ArraysEqual(t, array.Reverse([]string{"foo", "bar", "fizz", "buzz"}), []string{"buzz", "fizz", "bar", "foo"})
}
