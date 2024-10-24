package main

import (
	_ "embed"
	"flag"
	"log"

	"github.com/mleone10/advent-of-code-2022/pkg/newday"
)

func main() {
	day := flag.Int("day", 0, "day number to initialize")
	flag.Parse()

	if *day < 1 || *day > 25 {
		log.Fatalf("invalid day number provided (must be between 1 and 25)")
	}

	err := newday.InitializeDay(*day)
	if err != nil {
		log.Fatalf("failed to initialize new day: %v", err)
	}
}
