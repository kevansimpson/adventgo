package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/4">Day 4</a>
 */

import (
	"fmt"
	"strconv"

	"github.com/kevansimpson/util"
)

type Day04 struct{}

func (d Day04) countValidPasswords(start int, end int) (int, int) {
	pswdCount1, pswdCount2 := 0, 0
	for i := start; i <= end; i++ {
		pswd := fmt.Sprintf("%d", i)
		if d.isOrdered(i, pswd) {
			if util.HasLetterPair(pswd) {
				pswdCount1 += 1
			}

			if d.hasStandalonePair(pswd) {
				pswdCount2 += 1
			}
		}
	}

	return pswdCount1, pswdCount2
}

func (d Day04) isOrdered(num int, pswd string) bool {
	ordered, _ := strconv.Atoi(util.SortString(pswd))
	return ordered == num
}

func (d Day04) hasStandalonePair(pswd string) bool {
	chars := []rune(pswd)
	for _, ch := range chars {
		if 2 == util.CountMatches(chars, func(r rune) bool { return r == ch }) {
			return true
		}
	}
	return false
}
