package day10

import (
	"math"
	"sort"

	"github.com/mleone10/advent-of-code-go/internal/slice"
)

type Field struct {
	optLoc  location
	optView view
}

type location struct {
	x, y int
	r    float64
}

type view map[float64][]location

func NewField(grid string) Field {
	f := Field{}
	ls := readInput(grid)

	for _, l := range ls {
		v := computeStationView(l, ls)
		if len(v) > len(f.optView) {
			f.optView = v
			f.optLoc = l
		}
	}

	return f
}

func (f Field) AsteroidsInView() int {
	return len(f.optView)
}

// func main() {
// 	l, v := locateStation(readInput())
// 	log.Printf("Best location for station is at (%d,%d) with view to %d asteroids", l.x, l.y, len(v))

// 	a := v.findNthDestroyed(200)
// 	log.Printf("200th asteroid to be destroyed is at (%d,%d), whose hash is %d", a.x, a.y, a.x*100+a.y)
// }

func LocateIdealStation(grid string) (location, view) {
	ls := readInput(grid)
	var optLoc location
	var optView view
	for _, l := range ls {
		v := computeStationView(l, ls)
		if len(v) > len(optView) {
			optView = v
			optLoc = l
		}
	}
	return optLoc, optView
}

func readInput(grid string) []location {
	ls := []location{}

	for y, l := range slice.TrimSplit(grid) {
		for x, c := range l {
			if c == '#' {
				ls = append(ls, location{x: x, y: y})
			}
		}
	}

	return ls
}

func computeStationView(s location, ls []location) view {
	v := view{}
	for _, l := range ls {
		x, y := l.x-s.x, l.y-s.y
		a := bearing(x, y)
		l.r = math.Hypot(float64(x), float64(y))
		if l.r > 0 {
			v[a] = append(v[a], l)
		}
	}

	for _, a := range v {
		sort.Slice(a, func(i, j int) bool {
			return a[i].r < a[j].r
		})
	}

	return v
}

func (f Field) NthDestroyedProduct(n int) int {
	var l location
	for _, a := range angles(f.optView) {
		l, f.optView[a] = f.optView[a][0], f.optView[a][1:]
		if len(f.optView[a]) == 0 {
			delete(f.optView, a)
		}

		n--
		if n <= 0 {
			break
		}
	}

	return l.x*100 + l.y
}

func angles(v view) []float64 {
	as := []float64{}
	for a := range v {
		as = append(as, a)
	}
	sort.Float64s(as)
	return as
}

func bearing(x, y int) float64 {
	return math.Mod(360-math.Atan2(float64(-x), float64(-y))*(180/math.Pi), 360)
}
