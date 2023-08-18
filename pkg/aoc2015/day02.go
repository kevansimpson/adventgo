package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/2">Day 2</a>
 */

import (
	"sort"

	"github.com/kevansimpson/util"
)

type Present []int

func howMuchPaperAndRibbon(input []string) (int, int) {
	paper, ribbon := 0, 0
	for _, str := range input {
		dim := util.ExtractInts(str)
		sort.Ints(dim)
		paper += wrap(dim)
		ribbon += tieBow(dim)
	}
	return paper, ribbon
}

func wrap(p Present) int {
	return (3 * p[0] * p[1]) + (2 * p[1] * p[2]) + (2 * p[2] * p[0])
}

func tieBow(p Present) int {
	return (2 * p[0]) + (2 * p[1]) + /* bow */ (p[0] * p[1] * p[2])
}
