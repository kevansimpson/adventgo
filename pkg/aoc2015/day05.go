package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/5">Day 5</a>
 */

import (
	"regexp"
	"strings"

	"github.com/kevansimpson/util"
)

func oldNiceStrings(input []string) int {
	return util.CountMatches(input, isOldNice)
}

func newNiceStrings(input []string) int {
	return util.CountMatches(input, isNewNice)
}

var vowels = regexp.MustCompile(`[aeiou]`)
var badStrings = regexp.MustCompile(`(ab|cd|pq|xy)`)

func isOldNice(str string) bool {
	return (len(vowels.FindAllString(str, -1)) >= 3 &&
		hasDoubleLetters(str) &&
		len(badStrings.FindAllString(str, -1)) <= 0)
}

func isNewNice(str string) bool {
	return hasAdjacentPairs(str) && hasLetterSandwich(str)
}

func hasDoubleLetters(str string) bool {
	ltrs := []rune(str)
	max := len(str) - 1
	for ix := 0; ix < max; ix += 1 {
		if ltrs[ix] == ltrs[ix+1] {
			return true
		}
	}
	return false
}

func hasAdjacentPairs(str string) bool {
	max := len(str) - 2
	for ix := 0; ix < max; ix += 1 {
		if strings.Contains(str[ix+2:max+2], str[ix:ix+2]) {
			return true
		}
	}
	return false
}

func hasLetterSandwich(str string) bool {
	ltrs := []rune(str)
	max := len(str) - 2
	for ix := 0; ix < max; ix += 1 {
		if ltrs[ix] == ltrs[ix+2] {
			return true
		}
	}
	return false
}
