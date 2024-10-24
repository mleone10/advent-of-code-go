// Array contains utility functions to simplify common operations on slices and arrays.
package array

import (
	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Float | constraints.Integer
}

// Max returns the largest value among a slice of values via the `>` operator.
func Max[T number](args []T) T {
	return Reduce(args, func(max, v T) T {
		if max > v {
			return max
		}
		return v
	}, 0)
}

// Min returns the smallest value among a slice of values via the `<` operator.
func Min[T number](args []T) T {
	return Reduce(args[1:], func(min, v T) T {
		if min < v {
			return min
		}
		return v
	}, args[0])
}

// Sum returns the aggregate total of a slice of values via the `+=` operator.
func Sum[T number](args []T) T {
	return Reduce(args, func(acc, v T) T {
		return acc + v
	}, 0)
}

// Take returns the first n elements of the given slice.
func Take[T any](arr []T, n int) []T {
	if len(arr) <= n {
		return arr
	}

	return arr[:n]
}

// FrequencyList accepts a slice of comparable-type values and returns a map[T]int representing how many times each key appears in the input slice.
func FrequencyList[T comparable](args []T) map[T]int {
	freqs := map[T]int{}

	for _, arg := range args {
		freqs[arg] += 1
	}

	return freqs
}

// SlidingSum iterates over a given slice and sums sub-slices of w-width elements and returns an ordered slice of those sums.
func SlidingSum[T number](w int, args []T) []T {
	if len(args) < w {
		return []T{}
	}

	var slidingSum []T

	for i := range args[:len(args)-w+1] {
		slidingSum = append(slidingSum, Sum(args[i:i+w]))
	}

	return slidingSum
}

// Map applies a function to each element in a given slice and returns a new slice of the results.
func Map[T, R any](arr []T, f func(T) R) []R {
	var res []R
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}

// Reduce applies a function to each element in a slice, remembering the result of that function over subsequent iterations
func Reduce[T, R any](arr []T, f func(R, T) R, init R) R {
	res := init
	for _, v := range arr {
		res = f(res, v)
	}
	return res
}

// Filter applies a boolean function `f` to each element in a slice and returns a new slice containing only those elements `e` for which `f(e)` evaluates to true.
func Filter[T any](arr []T, f func(T) bool) []T {
	var res []T
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// Contains determines if a value `v` is found within a given slice `arr`, returning true if so and false otherwise.
func Contains[T comparable](arr []T, v T) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}

// ContainsSubarray returns true if `arr` contains all elements of `sub` in the same order.
func ContainsSubarray[T comparable](arr, sub []T) bool {
	if len(arr) < len(sub) {
		return false
	}
	if len(arr) == len(sub) {
		return Equal(arr, sub)
	}

	for i, a := range arr[:len(arr)-len(sub)+1] {
		if a == sub[0] && Equal(arr[i:i+len(sub)], sub) {
			return true
		}
	}

	return false
}

// Equal returns true if two slices are of the same length and contents.
func Equal[T comparable](x, y []T) bool {
	if len(x) != len(y) {
		return false
	}

	for i, v := range x {
		if v != y[i] {
			return false
		}
	}

	return true
}

// Reverse returns a new slice which is a reversed copy of the original one.
func Reverse[T any](arr []T) []T {
	var rev []T
	for i := range arr {
		rev = append(rev, arr[len(arr)-1-i])
	}
	return rev
}
