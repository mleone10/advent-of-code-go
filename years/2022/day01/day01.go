package day01

import (
	"sort"
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/internal/slice"
)

type Day1 struct {
	elves [][]int
}

func New(input string) Day1 {
	d := Day1{}

	for _, elf := range strings.Split(input, "\n\n") {
		calorieCounts := []int{}
		for _, snack := range strings.Split(elf, "\n") {
			snackCalories, _ := strconv.Atoi(snack)
			calorieCounts = append(calorieCounts, snackCalories)
		}
		d.elves = append(d.elves, calorieCounts)
	}

	return d
}

func (d Day1) MaxCaloriesSingleElf() int {
	return mth.Max(d.reduceElves()...)
}

func (d Day1) CaloriesTopThreeElves() int {
	elfCalories := d.reduceElves()

	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))

	return slice.Sum(slice.Take(elfCalories, 3))
}

func (d Day1) reduceElves() []int {
	elfCalories := []int{}
	for _, elf := range d.elves {
		elfCalories = append(elfCalories, slice.Sum(elf))
	}
	return elfCalories
}
