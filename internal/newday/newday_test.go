package newday_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/newday"
)

const (
	dayNum  = 99
	yearNum = 2099
)

var dayName = fmt.Sprintf("day%02d", dayNum)

func TestFilesCreated(t *testing.T) {
	tmp := t.TempDir()
	tcs := []string{
		filepath.Join(tmp, fmt.Sprintf("years/%[1]d/%[2]s/%[2]s.go", yearNum, dayName)),
		filepath.Join(tmp, fmt.Sprintf("years/%[1]d/%[2]s/%[2]s_test.go", yearNum, dayName)),
		filepath.Join(tmp, fmt.Sprintf("years/%d/%s/input.txt", yearNum, dayName)),
		filepath.Join(tmp, fmt.Sprintf("years/%d/%s/test_input.txt", yearNum, dayName)),
	}

	err := newday.Init(tmp, yearNum, dayNum)
	if err != nil {
		t.Fatalf("newday initialization failed: %v", err)
	}

	for _, tc := range tcs {
		if _, err := os.Stat(tc); errors.Is(err, os.ErrNotExist) {
			t.Errorf("file %s does not exist, but should", tc)
		}
	}
}

func TestSourceFilesPopulated(t *testing.T) {
	tmp := t.TempDir()
	tcs := []struct {
		file      string
		firstLine string
	}{
		{
			filepath.Join(tmp, fmt.Sprintf("years/%[1]d/%[2]s/%[2]s.go", yearNum, dayName)),
			fmt.Sprintf("package %s", dayName),
		}, {
			filepath.Join(tmp, fmt.Sprintf("years/%[1]d/%[2]s/%[2]s_test.go", yearNum, dayName)),
			fmt.Sprintf("package %s_test", dayName),
		},
	}

	err := newday.Init(tmp, yearNum, dayNum)
	if err != nil {
		t.Fatalf("newday initialization failed: %v", err)
	}

	for _, tc := range tcs {
		bs, err := os.ReadFile(tc.file)
		if err != nil {
			t.Errorf("failed to open %v to read: %v", tc.file, err)
			continue
		}
		actual := strings.Split(string(bs), "\n")[0]
		if actual != tc.firstLine {
			t.Errorf("first line [%s] did not match expected [%s]", actual, tc.firstLine)
		}
	}
}
