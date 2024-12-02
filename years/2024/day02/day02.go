package day02

import (
	"strings"

	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/internal/slice"
)

func IsSafe(l string, withProblemDampener bool) bool {
	// Convert the report to a slice of ints
	ns := slice.Map(strings.Split(l, " "), func(n string) int {
		return mth.Atoi(n)
	})

	// If the report is safe or the Problem Dampener is off, return whether the full report is safe
	s := isSafe(ns)
	if s || !withProblemDampener {
		return s
	}

	// Otherwise, iterate over variants of the report with each element removed; if a variant is safe, return
	for i := 0; i < len(ns); i++ {
		tmp := make([]int, i)
		copy(tmp, ns[:i])
		if isSafe(append(tmp, ns[i+1:]...)) {
			return true
		}
	}

	// No variants were safe
	return false
}

func isSafe(ns []int) bool {
	// Assume we're increasing, then iterate over the report
	inc := true
	for i, n := range ns {
		// If this is the first level in the report, do nothing
		if i == 0 {
			continue
		}
		// If this is the second level, determine if we're actually decreasing
		if i == 1 && n < ns[0] {
			inc = false
		}
		// If the current level matches that of the previous one, the report is unsafe
		if n == ns[i-1] {
			return false
		}
		// If we're increasing, check if the current level is greater than the previous by no more than 3
		if inc && (n < ns[i-1] || n > ns[i-1]+3) {
			return false
		}
		// If we're decreasing, check if the current level is less than the previous by no more than 3
		if !inc && (n < ns[i-1]-3 || n > ns[i-1]) {
			return false
		}
	}
	return true
}
