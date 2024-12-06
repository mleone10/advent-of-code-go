package day05

import (
	"sort"
	"strings"

	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/internal/set"
	"github.com/mleone10/advent-of-code-go/internal/slice"
)

type Rules struct {
	rm map[int]set.Set[int]
}

type Page struct {
	ns []int
	rs Rules
}

func NewRules(in string) Rules {
	rm := map[int]set.Set[int]{}
	for _, r := range slice.TrimSplit(in) {
		ps := strings.Split(r, "|")
		a := mth.Atoi(ps[0])
		if _, ok := rm[a]; !ok {
			rm[a] = set.Set[int]{}
		}
		rm[a].Add(mth.Atoi(ps[1]))
	}
	return Rules{rm}
}

func NewPages(in string, rs Rules) []Page {
	ps := []Page{}
	for _, p := range slice.TrimSplit(in) {
		ps = append(ps, Page{
			ns: slice.Map(strings.Split(p, ","), func(n string) int {
				return mth.Atoi(n)
			}),
			rs: rs})
	}
	return ps
}

func (p Page) GetMiddlePage() int {
	return p.ns[len(p.ns)/2]
}

func (p Page) Repair() {
	sort.Sort(p)
}

func (p Page) Len() int {
	return len(p.ns)
}

func (p Page) Swap(i, j int) {
	p.ns[i], p.ns[j] = p.ns[j], p.ns[i]
}

func (p Page) Less(i, j int) bool {
	return set.Contains(p.rs.rm[p.ns[i]], p.ns[j])
}
