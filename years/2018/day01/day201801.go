package day01

import (
	"github.com/mleone10/advent-of-code-go/internal/slice"
)

func CalcFinalFreq(fs []int) int {
	return slice.Sum(fs)
}

func FindFirstDuplicateFreq(fs []int) int {
	currentFreq := 0
	frequencies := map[int]bool{0: true}
	for {
		for _, f := range fs {
			currentFreq += f
			if _, ok := frequencies[currentFreq]; ok {
				return currentFreq
			} else {
				frequencies[currentFreq] = true
			}
		}
	}
}
