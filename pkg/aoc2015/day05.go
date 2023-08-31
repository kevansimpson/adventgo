package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/5">Day 5</a>
 */

import (
	"regexp"
	"strings"

	"github.com/kevansimpson/util"
)

type Day05 struct{}

func (d Day05) oldNiceStrings(input []string) int {
	return util.CountMatches(input, d.isOldNice)
}

func (d Day05) newNiceStrings(input []string) int {
	return util.CountMatches(input, d.isNewNice)
}

var vowels = regexp.MustCompile(`[aeiou]`)
var badStrings = regexp.MustCompile(`(ab|cd|pq|xy)`)

func (d Day05) isOldNice(str string) bool {
	return (len(vowels.FindAllString(str, -1)) >= 3 &&
		util.HasLetterPair(str) &&
		len(badStrings.FindAllString(str, -1)) <= 0)
}

func (d Day05) isNewNice(str string) bool {
	return d.hasAdjacentPairs(str) && util.HasLetterSandwich(str)
}

func (d Day05) hasAdjacentPairs(str string) bool {
	max := len(str) - 2
	for ix := 0; ix < max; ix += 1 {
		if strings.Contains(str[ix+2:max+2], str[ix:ix+2]) {
			return true
		}
	}
	return false
}
