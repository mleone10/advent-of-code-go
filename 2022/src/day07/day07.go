package day07

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/array"
	"github.com/mleone10/advent-of-code-2022/pkg/maputil"
)

type FileSystem struct {
	fs  map[string]int
	dir string
	ds  []string
}

const diskTotal = 70000000

func New(input []string) *FileSystem {
	d := FileSystem{
		fs: map[string]int{},
		ds: []string{},
	}

	for _, cmd := range input {
		switch cmd[:4] {
		case "$ cd":
			args := strings.Split(cmd, " ")
			d.ChangeDir(args[2])
		case "$ ls": // Do nothing
		default:
			d.Discover(cmd)
		}
	}

	return &d
}

func (d *FileSystem) ChangeDir(path string) {
	if strings.HasPrefix(path, "/") {
		d.dir = path
	} else {
		d.dir = filepath.Join(d.dir, path)
	}
}

func (d FileSystem) Pwd() string {
	return d.dir
}

func (d *FileSystem) Discover(file string) {
	fParts := strings.Split(file, " ")
	if strings.HasPrefix(file, "dir") {
		// File is a directory, just remember it exists
		d.ds = append(d.ds, filepath.Join(d.dir, fParts[1]))
	} else {
		size, _ := strconv.Atoi(fParts[0])
		d.fs[filepath.Join(d.dir, fParts[1])] = size
	}
}

func (d FileSystem) List() []string {
	fs := []string{}
	for dir := range d.fs {
		if filepath.Dir(dir) == d.dir {
			fs = append(fs, filepath.Base(dir))
		}
	}
	return fs
}

func (d FileSystem) DirectorySizes() map[string]int {
	dirs := map[string]int{}
	for _, dir := range d.ds {
		dirSum := 0
		for file, size := range d.fs {
			if filepath.HasPrefix(filepath.Dir(file), dir) {
				dirSum += size
			}
		}
		dirs[dir] = dirSum
	}
	return dirs
}

func (d FileSystem) UsedSpace() int {
	return array.Reduce(maputil.Values(d.fs), func(sum int, filesize int) int {
		return sum + filesize
	}, 0)
}

func (d FileSystem) FreeSpace() int {
	return diskTotal - d.UsedSpace()
}
