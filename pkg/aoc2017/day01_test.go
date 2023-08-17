package aoc2017

import (
	"testing"

	"github.com/kevansimpson/util"
)

func TestDay01Solutions(t *testing.T) {
	var input = util.ReadSingleLine("input01.txt")
	util.AssertEquals(t, "solve1()", solve1(input), 1251)
	util.AssertEquals(t, "solve2()", solve2(input), 1244)
}
