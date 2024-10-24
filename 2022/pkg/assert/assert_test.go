package assert_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
)

func TestEqual(t *testing.T) {
	testT := &testing.T{}

	assert.Equal(testT, 10, 10)
	assert.Equal(testT, "hello", "hello")
	assert.Equal(testT, true, true)

	if testT.Failed() {
		t.Error()
	}

	assert.Equal(testT, 10, 12)

	if !testT.Failed() {
		t.Error()
	}
}

func TestNotEqual(t *testing.T) {
	testT := &testing.T{}

	assert.NotEqual(testT, 10, 12)
	assert.NotEqual(testT, "hello", "world")
	assert.NotEqual(testT, true, false)

	if testT.Failed() {
		t.Error()
	}

	assert.NotEqual(testT, 10, 10)
	if !testT.Failed() {
		t.Error()
	}
}

func TestArraysEqual(t *testing.T) {
	testT := &testing.T{}

	assert.ArraysEqual(testT, []int{1, 2, 3}, []int{1, 2, 3})
	assert.ArraysEqual(testT, []float32{1.1, 2.2, 3.3}, []float32{1.1, 2.2, 3.3})
	assert.ArraysEqual(testT, []string{"foo", "bar", "fizz", "buzz"}, []string{"foo", "bar", "fizz", "buzz"})

	if testT.Failed() {
		t.Error()
	}

	assert.ArraysEqual(testT, []int{1, 2, 3}, []int{5, 10, 15})
	if !testT.Failed() {
		t.Error()
	}
}

func TestIsNil(t *testing.T) {
	testT := &testing.T{}

	var err error

	assert.IsNil(testT, nil)
	assert.IsNil(testT, err)

	if testT.Failed() {
		t.Error()
	}

	assert.IsNil(testT, 0)
	if !testT.Failed() {
		t.Error()
	}
}

func TestContains(t *testing.T) {
	testT := &testing.T{}
	testArr := []int{0, 2, 4, 8}

	assert.Contains(testT, testArr, 0)
	assert.Contains(testT, testArr, 2)

	if testT.Failed() {
		t.Error()
	}

	assert.Contains(testT, testArr, 1)

	if !testT.Failed() {
		t.Error()
	}
}

func TestDoesNotContain(t *testing.T) {
	testT := &testing.T{}
	testArr := []int{0, 2, 4, 8}

	assert.DoesNotContain(testT, testArr, 1)
	assert.DoesNotContain(testT, testArr, 3)

	if testT.Failed() {
		t.Error()
	}

	assert.DoesNotContain(testT, testArr, 2)

	if !testT.Failed() {
		t.Error()
	}
}
