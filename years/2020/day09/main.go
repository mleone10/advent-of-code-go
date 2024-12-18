package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/mleone10/advent-of-code-go/internal/mth"
)

const preambleLen = 25

func main() {
	is := []int{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		is = append(is, i)
	}

	invalidInt := findFirstInvalidNumber(is)
	log.Printf("First invalid XMAS number: %d", invalidInt)
	log.Printf("Encryption weakness: %d", findEncryptionWeakness(invalidInt, is))
}

func findFirstInvalidNumber(ints []int) int {
	for i := preambleLen; i < len(ints); i++ {
		testInts := make([]int, preambleLen)
		copy(testInts, ints[i-preambleLen:i])
		if !isValidTwoSum(ints[i], testInts) {
			return ints[i]
		}
	}
	log.Fatalln("No valid two sum found")
	return 0
}

func findEncryptionWeakness(target int, ints []int) int {
	subInts := findSubInts(target, ints)
	return intSliceMin(subInts) + mth.Max(subInts...)
}

func isValidTwoSum(target int, ints []int) bool {
	sort.Ints(ints)
	i, j := 0, len(ints)-1
	for i < j {
		sum := ints[i] + ints[j]
		if sum == target {
			return true
		} else if sum > target {
			j--
		} else if sum < target {
			i++
		}
	}
	return false
}

func findSubInts(target int, ints []int) []int {
	for i := range ints {
		for j := range ints[i:] {
			if intSliceSum(ints[i:i+j]) == target {
				return ints[i : i+j]
			}
		}
	}
	log.Fatal("No valid sub-slice found")
	return []int{}
}

func intSliceSum(ints []int) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return sum
}

func intSliceMin(ints []int) int {
	min := math.MaxInt64
	for _, i := range ints {
		min = mth.Min(i, min)
	}
	return min
}
