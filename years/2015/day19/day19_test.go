package day19_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-go/internal/assert"
	"github.com/mleone10/advent-of-code-go/years/2015/day19"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	replacements    string
	initMolecule    string
	expectedPartOne int
	expectedPartTwo int
}{
	{
		testInput,
		"HOH",
		4,
		3,
	},
	{
		testInput,
		"HOHOHO",
		7,
		6,
	},
	{
		input,
		"CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr",
		509,
		195,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		r := day19.NewReplacements(tc.replacements)
		actual := day19.CalibrationSum(tc.initMolecule, r)
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		r := day19.NewReplacements(tc.replacements)
		actual := day19.FabricationLength(tc.initMolecule, r)
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
