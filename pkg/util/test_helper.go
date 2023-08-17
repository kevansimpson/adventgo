package util

import (
	"testing"
)

func AssertSolutions(t *testing.T, s PuzzleSolver, input any, expected1 any, expected2 any) {
	AssertEquals(t, "solve1()", s.solve1(input), expected1)
	AssertEquals(t, "solve2()", s.solve2(input), expected2)
}

func AssertEquals(t *testing.T, name string, result any, expected any) {
	if expected != result {
		t.Errorf("%s == %d, want %d", name, result, expected)
	}
}
