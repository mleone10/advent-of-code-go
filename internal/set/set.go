package set

import "github.com/mleone10/advent-of-code-go/internal/mp"

type setVal comparable

// A Set[T] is a structure which holds unique values of type T.
type Set[T setVal] map[T]struct{}

// Add stores an element `v` in the Set.
func (s Set[T]) Add(vals ...T) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

// Remove eliminates the given element from the set if it exists.
func (s Set[T]) Remove(e T) {
	for k := range s {
		if k == e {
			delete(s, e)
		}
	}
}

// Size returns the number of unique elements in the Set.
func (s Set[T]) Size() int {
	return len(s)
}

// Slice returns a slice of unique values in the Set.
func (s Set[T]) Slice() []T {
	return mp.Keys(s)
}

// From creates a set using the values of the input arguments `vals`.
func From[T setVal](vals ...T) Set[T] {
	s := Set[T]{}
	s.Add(vals...)
	return s
}

// Contains returns true if the element t is within the set s.
func Contains[T setVal](s Set[T], t T) bool {
	f := false
	for e := range s {
		if e == t {
			f = true
		}
	}
	return f
}

// Intersection returns a set containing only those elements which appear in all provided sets.
func Intersection[T setVal](ss ...Set[T]) Set[T] {
	ret := ss[0]

	for _, r := range ss[1:] {
		ret = intersection(ret, r)
	}

	return ret
}

func intersection[T setVal](r, s Set[T]) Set[T] {
	// The loop below assumes that r is the larger set.  If that's not the case, though, call this method again with the arguments flipped.
	if len(s) > len(r) {
		return intersection(s, r)
	}

	ret := Set[T]{}
	for e := range r {
		if Contains(s, e) {
			ret.Add(e)
		}
	}
	return ret
}
