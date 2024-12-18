package day06

import (
	"strings"

	"github.com/mleone10/advent-of-code-go/internal/slice"
)

var FindStartOfPacket = findFirstNUnique(4)
var FindStartOfMessage = findFirstNUnique(14)

func findFirstNUnique(n int) func(string) int {
	return func(buf string) int {
		for i := 0; i <= len(buf)-n; i++ {
			if allUnique(buf[i : i+n]) {
				return i + n
			}
		}
		return 0
	}
}

func allUnique(arr string) bool {
	return len(slice.FrequencyList(strings.Split(arr, ""))) == len(arr)
}
