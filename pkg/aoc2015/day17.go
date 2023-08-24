package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/17">Day 17</a>
 */

import (
	"math"
)

type Day17 struct{}

type CanPermutations struct {
	totalPermutations, totalWith150Litres int
	fewestCans, canCount                  int
}

func (d Day17) transferEggnog(cans []int) (int, int) {
	cp := CanPermutations{0, 0, math.MaxInt, len(cans)}
	perm := make([]bool, cp.canCount)

	d.sumCans(&cp, perm, 0, cans)
	return cp.totalPermutations, cp.totalWith150Litres
}

func (d Day17) sumCans(cp *CanPermutations, perm []bool, index int, cans []int) {
	if index >= cp.canCount {
		if sum(perm, cans) == 150 {
			cp.totalPermutations += 1

			used := used(perm, cp.canCount)
			if used < cp.fewestCans {
				cp.fewestCans = used
				cp.totalWith150Litres = 1
			} else if used == cp.fewestCans {
				cp.totalWith150Litres += 1
			}
		}
		return
	}

	off, on := make([]bool, cp.canCount), make([]bool, cp.canCount)
	for i, b := range perm {
		off[i] = b
		on[i] = b
	}
	d.sumCans(cp, off, 1+index, cans)
	on[index] = true
	d.sumCans(cp, on, 1+index, cans)
}

func sum(perm []bool, cans []int) int {
	sum := 0
	for i, c := range cans {
		if perm[i] {
			sum += c
		}
	}

	return sum
}

func used(perm []bool, cans int) int {
	used := 0
	for _, b := range perm {
		if b {
			used += 1
		}
	}
	return used
}
