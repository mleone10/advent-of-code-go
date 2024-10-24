package day09

import (
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/grid"
	"github.com/mleone10/advent-of-code-2022/pkg/linkedlist"
	"github.com/mleone10/advent-of-code-2022/pkg/maputil"
	"github.com/mleone10/advent-of-code-2022/pkg/mathutil"
)

type Dir grid.Point

var (
	DirUp    = Dir{X: 0, Y: -1}
	DirDown  = Dir{X: 0, Y: 1}
	DirLeft  = Dir{X: -1, Y: 0}
	DirRight = Dir{X: 1, Y: 0}
)

var Cmds = map[string]Dir{
	"U": DirUp,
	"D": DirDown,
	"L": DirLeft,
	"R": DirRight,
}

type Knot struct {
	Pos     grid.Point
	visited map[grid.Point]bool
}

func NewRope(len int) *linkedlist.Node[*Knot] {
	head := linkedlist.NewNode(newKnot())
	for i := 0; i < len-1; i++ {
		head.Tail().LinkNext(linkedlist.NewNode(newKnot()))
	}
	return head
}

func newKnot() *Knot {
	return &Knot{visited: map[grid.Point]bool{{X: 0, Y: 0}: true}}
}

func SimulateMoves(r *linkedlist.Node[*Knot], mvs []string) {
	for _, mv := range mvs {
		mParts := strings.Split(mv, " ")
		dist, _ := strconv.Atoi(mParts[1])
		MoveN(r, Cmds[mParts[0]], dist)
	}
}

func MoveN(r *linkedlist.Node[*Knot], d Dir, dist int) {
	for i := 0; i < dist; i++ {
		moveHead(r, d)
	}
}

func moveHead(r *linkedlist.Node[*Knot], d Dir) {
	r.Value().Pos.X += d.X
	r.Value().Pos.Y += d.Y
	r.Value().Visit()

	if r.Next() != nil {
		updateKnot(r.Next())
	}
}

func updateKnot(k *linkedlist.Node[*Knot]) {
	var moved bool
	dx := k.Prev().Value().Pos.X - k.Value().Pos.X
	dy := k.Prev().Value().Pos.Y - k.Value().Pos.Y
	adx := mathutil.Abs(dx)
	ady := mathutil.Abs(dy)

	if adx > 1 && ady > 1 {
		k.Value().Pos.X += (dx / adx)
		k.Value().Pos.Y += (dy / ady)
		moved = true
	} else if adx >= 2 {
		k.Value().Pos.X += (dx / adx)
		if ady != 0 {
			k.Value().Pos.Y += (dy / ady)
		}
		moved = true
	} else if ady >= 2 {
		k.Value().Pos.Y += (dy / ady)
		if adx != 0 {
			k.Value().Pos.X += (dx / adx)
		}
		moved = true
	}

	// If we moved at all, remember the new position.
	if moved {
		k.Value().Visit()
	}

	// If there's another knot, update it's position as well.
	if k.Next() != nil {
		updateKnot(k.Next())
	}
}

func (k *Knot) Visit() {
	k.visited[grid.Point{X: k.Pos.X, Y: k.Pos.Y}] = true
}

func (k Knot) Visited() []grid.Point {
	return maputil.Keys(k.visited)
}
