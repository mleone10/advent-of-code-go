package geo

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/mleone10/advent-of-code-go/internal/mp"
	"github.com/mleone10/advent-of-code-go/internal/mth"
)

const (
	runePopulated = '#'
	runeEmpty     = '.'
)

type (
	mapFunc[T, V comparable]        func(p Point[T]) Point[V]
	filterFunc[T comparable]        func(p Point[T]) bool
	reduceFunc[T comparable, V any] func(p Point[T], r V) V
)

type Space[T comparable] interface {
	fmt.Stringer
	Set(l Location, v T)
	Get(l Location) T
	Size() int
	Iter() iter.Seq[Point[T]]
}

type Point[T any] struct {
	Loc Location
	Val T
}

func mapSpace[T, V comparable, S Space[T], R Space[V]](s S, r R, f mapFunc[T, V]) R {
	for p := range s.Iter() {
		v := f(p)
		r.Set(v.Loc, v.Val)
	}
	return r
}

func Map1D[T, V comparable](s Space1D[T], f mapFunc[T, V]) Space1D[V] {
	ret := Space1D[V]{}
	return mapSpace(s, ret, f)
}

func Map2D[T, V comparable](s Space2D[T], f mapFunc[T, V]) Space2D[V] {
	ret := Space2D[V]{}
	return mapSpace(s, ret, f)
}

func Map3D[T, V comparable](s Space3D[T], f mapFunc[T, V]) Space3D[V] {
	ret := Space3D[V]{}
	return mapSpace(s, ret, f)
}

func Map4D[T, V comparable](s Space4D[T], f mapFunc[T, V]) Space4D[V] {
	ret := Space4D[V]{}
	return mapSpace(s, ret, f)
}

func filterSpace[T comparable, S Space[T]](s S, r S, f filterFunc[T]) S {
	for p := range s.Iter() {
		if f(p) {
			r.Set(p.Loc, p.Val)
		}
	}
	return r
}

func Filter1D[T comparable](s Space1D[T], f filterFunc[T]) Space1D[T] {
	ret := Space1D[T]{}
	return filterSpace(s, ret, f)
}

func Filter2D[T comparable](s Space2D[T], f filterFunc[T]) Space2D[T] {
	ret := Space2D[T]{}
	return filterSpace(s, ret, f)
}

func Filter3D[T comparable](s Space3D[T], f filterFunc[T]) Space3D[T] {
	ret := Space3D[T]{}
	return filterSpace(s, ret, f)
}

func Filter4D[T comparable](s Space4D[T], f filterFunc[T]) Space4D[T] {
	ret := Space4D[T]{}
	return filterSpace(s, ret, f)
}

func Reduce[T comparable, V any, S Space[T]](s S, acc V, f reduceFunc[T, V]) V {
	for p := range s.Iter() {
		acc = f(p, acc)
	}
	return acc
}

type Space1D[T comparable] map[int]Point[T]

func (s Space1D[T]) Set(l Location, v T) {
	s[l.A] = Point[T]{l, v}
}

func (s Space1D[T]) Get(l Location) T {
	return s[l.A].Val
}

func (s Space1D[T]) Size() int {
	return len(s)
}

func (s Space1D[T]) String() string {
	cs := mp.Keys(s)
	slices.Sort(cs)

	var tmpl string
	for range cs[len(cs)-1] + 1 {
		tmpl += string(runeEmpty)
	}

	rs := []rune(tmpl)
	for _, c := range cs {
		rs[c] = runePopulated
	}

	return string(rs)
}

func (s Space1D[T]) Iter() iter.Seq[Point[T]] {
	return func(yield func(Point[T]) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

type Space2D[T comparable] map[int]Space1D[T]

func (s Space2D[T]) Set(l Location, v T) {
	if _, ok := s[l.B]; !ok {
		s[l.B] = Space1D[T]{}
	}
	s[l.B].Set(l, v)
}

func (s Space2D[T]) Get(l Location) T {
	return s[l.B].Get(l)
}

func (s Space2D[T]) Size() int {
	var ret int
	for _, v := range s {
		ret += v.Size()
	}
	return ret
}

func (s Space2D[T]) String() string {
	var ret []string
	maxLength := 0
	for i := range maxWidth(mp.Keys(s)) + 1 {
		if r, ok := s[i]; ok {
			ns := r.String()
			ret = append(ret, ns)
			maxLength = mth.Max(len(ns), maxLength)
		} else {
			ret = append(ret, "")
		}
	}

	for i, s := range ret {
		if len(s) < maxLength {
			ret[i] = s + strings.Repeat(string(runeEmpty), maxLength-len(s))
		}
	}

	return strings.Join(ret, "\n")
}

func (s Space2D[T]) Iter() iter.Seq[Point[T]] {
	return func(yield func(Point[T]) bool) {
		for _, v := range s {
			for u := range v.Iter() {
				if !yield(u) {
					return
				}
			}
		}
	}
}

type Space3D[T comparable] map[int]Space2D[T]

func (s Space3D[T]) Set(l Location, v T) {
	if _, ok := s[l.C]; !ok {
		s[l.C] = Space2D[T]{}
	}
	s[l.C].Set(l, v)
}

func (s Space3D[T]) Get(l Location) T {
	return s[l.C].Get(l)
}

func (s Space3D[T]) Size() int {
	var ret int
	for _, v := range s {
		ret += v.Size()
	}
	return ret
}

// TODO: implement Space3D.String()
func (s Space3D[T]) String() string {
	return ""
}

func (s Space3D[T]) Iter() iter.Seq[Point[T]] {
	return func(yield func(Point[T]) bool) {
		for _, v := range s {
			for u := range v.Iter() {
				if !yield(u) {
					return
				}
			}
		}
	}
}

type Space4D[T comparable] map[int]Space3D[T]

func (s Space4D[T]) Set(l Location, v T) {
	if _, ok := s[l.D]; !ok {
		s[l.D] = Space3D[T]{}
	}
	s[l.D].Set(l, v)
}

func (s Space4D[T]) Get(l Location) T {
	return s[l.D].Get(l)
}

func (s Space4D[T]) Size() int {
	var ret int
	for _, v := range s {
		ret += v.Size()
	}
	return ret
}

// TODO: implement Space4D.String()
func (s Space4D[T]) String() string {
	return ""
}

func (s Space4D[T]) Iter() iter.Seq[Point[T]] {
	return func(yield func(Point[T]) bool) {
		for _, v := range s {
			for u := range v.Iter() {
				if !yield(u) {
					return
				}
			}
		}
	}
}

func maxWidth(s []int) int {
	slices.Sort(s)
	return s[len(s)-1] - s[0]
}
