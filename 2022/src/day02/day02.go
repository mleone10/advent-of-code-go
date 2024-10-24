package day02

import (
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/array"
)

type Day02 struct {
	rounds []round
}

type round struct {
	// oppMove is my opponent's move: A = rock, B = paper, C = scissors
	oppMove string
	// myMove is my move: X = rock, Y = paper, Z = scissors
	myMove string
}

// scoreShape represents the static point value for my moves: 1 if I choose rock, 2 if paper, 3 if scissors.
var scoreShape = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

// scoreResult represents the win/loss/draw point value for all nine possible outcomes.
var scoreResult = map[round]int{
	{"A", "X"}: 3,
	{"A", "Y"}: 6,
	{"A", "Z"}: 0,
	{"B", "X"}: 0,
	{"B", "Y"}: 3,
	{"B", "Z"}: 6,
	{"C", "X"}: 6,
	{"C", "Y"}: 0,
	{"C", "Z"}: 3,
}

// myMoveAi provides my ideal move given an opponent's move and a desired outcome (X means lose, Y draw, Z win)
var myMoveAi = map[string]map[string]string{
	"A": {
		"X": "Z",
		"Y": "X",
		"Z": "Y",
	},
	"B": {
		"X": "X",
		"Y": "Y",
		"Z": "Z",
	},
	"C": {
		"X": "Y",
		"Y": "Z",
		"Z": "X",
	},
}

func New(input string) Day02 {
	rds := []round{}

	for _, r := range strings.Split(input, "\n") {
		moves := strings.Split(r, " ")
		rds = append(rds, round{moves[0], moves[1]})
	}

	return Day02{rds}
}

func (d Day02) SolvePartOne() int {
	return array.Reduce(d.rounds, func(s int, r round) int {
		return s + simulateRound(r)
	}, 0)
}

func (d Day02) SolvePartTwo() int {
	return array.Reduce(d.rounds, func(s int, r round) int {
		return s + simulateRound(round{r.oppMove, myMoveAi[r.oppMove][r.myMove]})
	}, 0)
}

func simulateRound(r round) int {
	return scoreShape[r.myMove] + scoreResult[r]
}
