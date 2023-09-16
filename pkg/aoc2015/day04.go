package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/0">Day 0</a>
 */

import "github.com/kevansimpson/util"

const SecretKey_04_2015 = "bgvyzdsv"

type Day04 struct{}

func (d Day04) fiveZeroHash(input string) int {
	return d.nextHash(input, "00000", util.FullOrFast(1, 254000))
}

func (d Day04) sixZeroHash(input string) int {
	return d.nextHash(input, "000000", util.FullOrFast(1, 1038000))
}

func (d Day04) nextHash(input string, prefix string, start int) int {
	_, ix := util.NextHashWithPrefix(input, prefix, start)
	return ix
}
