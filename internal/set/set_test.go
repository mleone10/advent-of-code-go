package set_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/set"
)

func TestAddAndSize(t *testing.T) {
	s := set.Set[int]{}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	assert.Equals(t, 3, s.Size())
}

func TestAddMultiple(t *testing.T) {
	s := set.Set[int]{}

	s.Add([]int{1, 2, 3}...)

	assert.Equals(t, 3, s.Size())
}

func TestRemove(t *testing.T) {
	subj := set.From(1, 2, 3, 4, 5)

	subj.Remove(1)
	assert.Equals(t, 4, len(subj))
	assert.Equals(t, false, set.Contains(subj, 1))
}

func TestSlice(t *testing.T) {
	s := set.Set[int]{}

	s.Add([]int{1, 2, 3}...)

	assert.ArrayEquals(t, []int{1, 2, 3}, s.Slice())
}

func TestFrom(t *testing.T) {
	subj := []int{1, 2, 3, 4, 5, 6, 1, 1, 4, 10}

	actual := set.From(subj...)

	assert.Equals(t, 7, len(actual))
	assert.ArrayEquals(t, []int{1, 2, 3, 4, 5, 6, 10}, actual.Slice())
}

func TestContains(t *testing.T) {
	subj := set.From(1, 2, 3, 4, 5)

	assert.Equals(t, true, set.Contains(subj, 3))
	assert.Equals(t, false, set.Contains(subj, 6))
}

func TestIntersection(t *testing.T) {
	a := set.From(1, 2, 3, 4, 5)
	b := set.From(3, 4, 5, 6)
	c := set.From(4, 5, 6, 7)

	actual := set.Intersection(a, b, c).Slice()
	assert.Equals(t, 2, len(actual))
	assert.Contains(t, actual, 4)
	assert.Contains(t, actual, 5)
}
