package day01

import (
	"sort"
	"strings"

	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/internal/slice"
)

type Day01 struct {
	Input string
	a, b  []int
}

func NewDay01(in string) Day01 {
	d := Day01{in, []int{}, []int{}}
	for _, l := range slice.TrimSplit(in) {
		ns := strings.Split(l, "   ")
		d.a = append(d.a, mth.Atoi(ns[0]))
		d.b = append(d.b, mth.Atoi(ns[1]))
	}

	sort.Ints(d.a)
	sort.Ints(d.b)

	return d
}

func (d Day01) Distance() int {
	var ret int

	for i, a := range d.a {
		ret += mth.Abs(a - d.b[i])
	}

	return ret
}

func (d Day01) SimilarityScore() int {
	var ret int

	// Reusing the FrequencyList function from 2022!
	freq := slice.FrequencyList(d.b)

	for _, a := range d.a {
		ret += a * freq[a]
	}

	return ret
}
