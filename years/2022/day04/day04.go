package day04

import (
	"strconv"
	"strings"
)

type Day04 struct {
	R1 rng
	R2 rng
}

type rng struct {
	Start, End int
}

func New(line string) Day04 {
	return Day04{
		newRng(strings.Split(line, ",")[0]),
		newRng(strings.Split(line, ",")[1]),
	}
}

func newRng(span string) rng {
	start, _ := strconv.Atoi(strings.Split(span, "-")[0])
	end, _ := strconv.Atoi(strings.Split(span, "-")[1])
	return rng{start, end}
}
