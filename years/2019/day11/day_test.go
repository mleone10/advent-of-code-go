package day11_test

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo string
}{
	{
		input,
		1709,
		"PGUEHCJH",
	},
}

func TestSolvePartOne(t *testing.T) {}

func TestSolvePartTwo(t *testing.T) {}
