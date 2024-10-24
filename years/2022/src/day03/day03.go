package day03

import (
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/array"
)

type Day03 struct {
	Input string
}

// TODO: Move this elsewhere, perhaps as a HalveString.  Or refactor as SplitStringNParts.
func GroupCompartments(contents string) []string {
	return []string{contents[:len(contents)/2], contents[len(contents)/2:]}
}

// TODO: Move this elsewhere, refactor as a generalized "find common characters between N strings" method
func FindCommonContents(compartments []string) string {
	common := compartments[0]

	var newCommon []rune
	for _, comp := range compartments[1:] {
		for _, char := range comp {
			if array.Contains([]rune(common), char) {
				newCommon = append(newCommon, char)
			}
		}
		common = string(newCommon)
		newCommon = nil
	}

	return common
}

func CalculatePriority(b byte) int {
	if string(b) == strings.ToLower(string(b)) {
		return int(b) - 96
	}
	return int(b) - 38
}
