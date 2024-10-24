package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	finalFreq := 0
	changes := []int{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		changes = append(changes, i)
		finalFreq += i
	}

	fmt.Println("Final Frequency:", finalFreq)

	currentFreq := 0
	frequencies := map[int]bool{0: true}
	for {
		for _, change := range changes {
			currentFreq += change
			if _, ok := frequencies[currentFreq]; ok {
				fmt.Println("First visited twice:", currentFreq)
				os.Exit(0)
			} else {
				frequencies[currentFreq] = true
			}
		}
	}

}
