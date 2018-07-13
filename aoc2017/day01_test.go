package aoc2017

import (
  "testing"
)

func TestSolve1(t *testing.T) {
  result := Solve1(Input())
  if 1251 != result {
    t.Errorf("Solve1() == %d, want %d", result, 1251)
	}
}

func TestSolve2(t *testing.T) {
  result := Solve2(Input())
  if 1244 != result {
    t.Errorf("Solve2() == %d, want %d", result, 1244)
	}
}
