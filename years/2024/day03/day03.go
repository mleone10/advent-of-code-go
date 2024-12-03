package day03

import (
	"regexp"

	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/internal/slice"
)

// This regex matches instances of `mul(X,Y)` (with submatch on X and Y), `do()`, and `dont'()`
var mulReg = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)

func InterpretCorruptedMemory(l string, handleConditionals bool) int {
	// Extract all matches in the corrupted memory string
	ms := mulReg.FindAllStringSubmatch(l, -1)

	enabled := true
	// Iterate through all matches in order
	return slice.Reduce(ms, 0, func(match []string, ret int) int {
		if match[0] == "don't()" {
			enabled = false
		} else if match[0] == "do()" {
			enabled = true
		} else if !handleConditionals || enabled {
			// If we're not currently handling conditionals (as in part 1) or we are and we're currently enabled, execute the multiplication instruction
			return ret + (mth.Atoi(match[1]) * mth.Atoi(match[2]))
		}
		// Otherwise, do nothing
		return ret
	})
}
