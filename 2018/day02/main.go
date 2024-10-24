package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type id struct {
	value string
}

func main() {
	ids := []id{}
	numIdsWithTwoLetters := 0
	numIdsWithThreeLetters := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		newId := id{value: input}
		ids = append(ids, newId)
		if newId.hasLetterXTimes(2) {
			numIdsWithTwoLetters++
		}
		if newId.hasLetterXTimes(3) {
			numIdsWithThreeLetters++
		}
	}

	fmt.Println("Checksum:", numIdsWithTwoLetters*numIdsWithThreeLetters)

	for i, id := range ids[:len(ids)-1] {
		if commonChars, found := id.findPartner(ids[i+1:]); found {
			fmt.Println("Common chars in partners:", commonChars)
			os.Exit(1)
		}
	}
}

func (i id) hasLetterXTimes(occurrences int) bool {
	for _, c := range i.value {
		if strings.Count(i.value, string(c)) == occurrences {
			return true
		}
	}
	return false
}

func (i id) findPartner(candidates []id) (string, bool) {
	for _, candidate := range candidates {
		if commonChars, found := i.isPartner(candidate); found {
			return commonChars, true
		}
	}

	return "", false
}

func (i id) isPartner(candidate id) (string, bool) {
	indexDiff := -1
	for i, c := range i.value {
		if candidate.value[i] != byte(c) {
			if indexDiff == -1 {
				indexDiff = i
			} else {
				return "", false
			}
		}
	}
	if indexDiff == -1 {
		return "", false
	}
	return i.value[:indexDiff] + i.value[indexDiff+1:], true
}
