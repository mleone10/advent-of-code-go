package day02

import (
	"strings"
)

func ComputeChecksum(ids []string) int {
	hasLetter2Times := hasLetterNTimes(2)
	hasLetter3Times := hasLetterNTimes(3)
	numIdsWithTwoLetters := 0
	numIdsWithThreeLetters := 0

	for _, id := range ids {
		if hasLetter2Times(id) {
			numIdsWithTwoLetters++
		}
		if hasLetter3Times(id) {
			numIdsWithThreeLetters++
		}
	}

	return numIdsWithTwoLetters * numIdsWithThreeLetters
}

func FindCommonCharsBetweenTargetBoxIds(ids []string) string {
	for i, id := range ids {
		for _, candidate := range ids[i+1:] {
			diffIndices := findDiffIndices(id, candidate)
			if len(diffIndices) == 1 {
				return id[:diffIndices[0]] + id[diffIndices[0]+1:]
			}
		}
	}

	return ""
}

func hasLetterNTimes(n int) func(id string) bool {
	return func(id string) bool {
		for _, c := range id {
			if strings.Count(id, string(c)) == n {
				return true
			}
		}
		return false
	}
}

func findDiffIndices(a, b string) []int {
	ret := []int{}
	for i, c := range a {
		if b[i] != byte(c) {
			ret = append(ret, i)
		}
	}
	return ret
}
