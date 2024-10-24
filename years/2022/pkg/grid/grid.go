package grid

import (
	"github.com/mleone10/advent-of-code-2022/pkg/array"
	"github.com/mleone10/advent-of-code-2022/pkg/maputil"
)

type item interface {
	any
}

type Plane[T item] struct {
	grid          map[int]map[int]T
	rows, cols    map[int][]T
	height, width int
}

func (p *Plane[T]) Set(x, y int, item T) {
	if p.grid == nil {
		p.grid = map[int]map[int]T{}
	}
	if p.grid[y] == nil {
		p.grid[y] = map[int]T{}
	}
	p.grid[y][x] = item
}

func (p Plane[T]) Get(x, y int) T {
	return p.grid[y][x]
}

func (p Plane[T]) Has(x, y int) bool {
	if row, rowExists := p.grid[y]; rowExists {
		if _, colExists := row[x]; colExists {
			return true
		}
	}
	return false
}

func (p Plane[T]) All() [][]T {
	fullGrid := [][]T{}
	for i := 0; i < p.Height(); i++ {
		gridRow := []T{}
		for j := 0; j < p.Width(); j++ {
			if item, ok := p.grid[i][j]; ok {
				gridRow = append(gridRow, item)
			} else {
				var item T
				gridRow = append(gridRow, item)
			}
		}
		fullGrid = append(fullGrid, gridRow)
	}
	return fullGrid
}

func (p *Plane[T]) Width() int {
	if p.width == 0 {
		p.width = array.Max(array.Map(maputil.Values(p.grid), func(row map[int]T) int {
			return array.Max(maputil.Keys(row))
		})) + 1
	}
	return p.width
}

func (p *Plane[T]) Height() int {
	if p.height == 0 {
		p.height = array.Max(maputil.Keys(p.grid)) + 1
	}
	return p.height
}

func (p *Plane[T]) Row(y int) []T {
	if p.rows == nil {
		p.rows = map[int][]T{}
	}
	if p.rows[y] != nil {
		return p.rows[y]
	}

	var row []T
	for i := 0; i < p.Width(); i++ {
		if item, ok := p.grid[y][i]; ok {
			row = append(row, item)
		} else {
			var item T
			row = append(row, item)
		}
	}

	p.rows[y] = row
	return row
}

func (p *Plane[T]) Col(x int) []T {
	if p.cols == nil {
		p.cols = map[int][]T{}
	}
	if p.cols[x] != nil {
		return p.cols[x]
	}

	var col []T
	for i := 0; i < p.Height(); i++ {
		if item, ok := p.grid[i][x]; ok {
			col = append(col, item)
		} else {
			var item T
			col = append(col, item)
		}
	}

	p.cols[x] = col
	return col
}

func (p Plane[T]) Sparse() map[int]map[int]T {
	return p.grid
}
