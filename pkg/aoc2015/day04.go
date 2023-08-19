package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/0">Day 0</a>
 */

import "github.com/kevansimpson/util"

const SecretKey_04_2015 = "bgvyzdsv"

func fiveZeroHash(input string) int {
	return nextHash(input, "00000", util.FullOrFast(1, 254000))
}

func sixZeroHash(input string) int {
	return nextHash(input, "000000", util.FullOrFast(1, 1038000))
}

func nextHash(input string, prefix string, start int) int {
	_, ix := util.NextHash(input, prefix, start)
	return ix
}