package day07_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/internal/mp"
	"github.com/mleone10/advent-of-code-go/internal/mth"
	"github.com/mleone10/advent-of-code-go/years/2022/day07"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

const DiskNeeded = 30000000

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo int
}{
	{
		input:           testInput,
		expectedPartOne: 95437,
		expectedPartTwo: 24933642,
	},
	{
		input:           input,
		expectedPartOne: 1307902,
		expectedPartTwo: 7068748,
	},
}

func TestChangeDir(t *testing.T) {
	d := day07.FileSystem{}
	d.ChangeDir("/")
	assert.Equals(t, d.Pwd(), "/")
	d.ChangeDir("a")
	assert.Equals(t, d.Pwd(), "/a")
	d.ChangeDir("..")
	assert.Equals(t, d.Pwd(), "/")
}

func TestDiscover(t *testing.T) {
	d := day07.New([]string{})
	d.ChangeDir("/")
	d.Discover("dir a")
	d.Discover("123 b.txt")
	d.Discover("456 c.txt")
	assert.Contains(t, d.List(), "b.txt")
	assert.Contains(t, d.List(), "c.txt")
}

func TestTotalUsedSpace(t *testing.T) {
	d := day07.New(strings.Split(strings.TrimSpace(testInput), "\n"))
	assert.Equals(t, d.UsedSpace(), 48381165)
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		d := day07.New(strings.Split(strings.TrimSpace(tc.input), "\n"))
		sum := 0
		for _, size := range d.DirectorySizes() {
			if size <= 100000 {
				sum += size
			}
		}
		assert.Equals(t, sum, tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		d := day07.New(strings.Split(strings.TrimSpace(tc.input), "\n"))
		candidates := map[string]int{}
		spaceNeeded := DiskNeeded - d.FreeSpace()
		for dir, size := range d.DirectorySizes() {
			if size >= spaceNeeded {
				candidates[dir] = size
			}
		}
		assert.Equals(t, mth.Min(mp.Values(candidates)...), tc.expectedPartTwo)
	}
}
