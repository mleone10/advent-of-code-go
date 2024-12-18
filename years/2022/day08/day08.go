package day08

import (
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/internal/slice"
	"github.com/mleone10/advent-of-code-go/years/2022/pkg/grid"
)

type Day08 struct {
	Input string
}

func NewGrid(input string) grid.Plane[int] {
	g := grid.Plane[int]{}
	for i, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for j, col := range strings.Split(row, "") {
			height, _ := strconv.Atoi(col)
			g.Set(j, i, height)
		}
	}
	return g
}

func northTreeHeights(g grid.Plane[int], x, y int) []int {
	return g.Col(x)[:y]
}

func southTreeHeights(g grid.Plane[int], x, y int) []int {
	return g.Col(x)[y+1:]
}

func westTreeHeights(g grid.Plane[int], x, y int) []int {
	return g.Row(y)[:x]
}

func eastTreeHeights(g grid.Plane[int], x, y int) []int {
	return g.Row(y)[x+1:]
}

func IsVisible(g grid.Plane[int], x, y int) bool {
	if x == 0 || y == 0 || x == g.Width()-1 || y == g.Height()-1 {
		// Point is the edge of the map, thus is visible
		return true
	}
	return isVisibleThroughTrees(slice.Reverse(northTreeHeights(g, x, y)))(g, x, y) ||
		isVisibleThroughTrees(southTreeHeights(g, x, y))(g, x, y) ||
		isVisibleThroughTrees(slice.Reverse(westTreeHeights(g, x, y)))(g, x, y) ||
		isVisibleThroughTrees(eastTreeHeights(g, x, y))(g, x, y)
}

func isVisibleThroughTrees(trees []int) func(g grid.Plane[int], x, y int) bool {
	return func(g grid.Plane[int], x, y int) bool {
		return mth.Max(trees...) < g.Get(x, y)
	}
}

func ScenicScore(g grid.Plane[int], x, y int) int {
	return directionalScenicScore(slice.Reverse(northTreeHeights(g, x, y)))(g, x, y) *
		directionalScenicScore(southTreeHeights(g, x, y))(g, x, y) *
		directionalScenicScore(slice.Reverse(westTreeHeights(g, x, y)))(g, x, y) *
		directionalScenicScore(eastTreeHeights(g, x, y))(g, x, y)
}

func directionalScenicScore(trees []int) func(g grid.Plane[int], x, y int) int {
	return func(g grid.Plane[int], x, y int) int {
		sightLine := 0
		for i := 0; i < len(trees); i++ {
			sightLine += 1
			if trees[i] >= g.Get(x, y) {
				break
			}
		}
		return sightLine
	}
}
