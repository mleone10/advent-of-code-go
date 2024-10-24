package day02

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
)

func TestParseInput(t *testing.T) {
	d := New(
		`A Y
			B X
			C Z`)
	assert.Equal(t, len(d.rounds), 3)
}

func TestSimulateRound(t *testing.T) {
	tcs := []struct {
		round         round
		expectedScore int
	}{
		{round{"A", "Y"}, 8},
		{round{"B", "X"}, 1},
		{round{"C", "Z"}, 6},
	}

	for _, tc := range tcs {
		assert.Equal(t, simulateRound(tc.round), tc.expectedScore)
	}
}
