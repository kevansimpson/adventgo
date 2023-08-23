package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/0">Day 0</a>
 */

import (
	"regexp"

	"github.com/kevansimpson/util"
)

const SANTA_CURRENT_PASSWORD_11_2015 = "vzbxkghb"

type Day11 struct{}

func (d Day11) nextTwoPasswords(input string) (string, string) {
	first := d.nextValidPswd(input)
	return first, d.nextValidPswd(first)
}

func (d Day11) isValidPswd(pswd string) bool {
	return !d.isPswdDisallowed(pswd) && d.hasDistinctPairs(pswd) && d.hasLetterSequence(pswd)
}

func (d Day11) nextValidPswd(pswd string) string {
	next := d.nextPswd(pswd)
	for !d.isValidPswd(next) {
		next = d.nextPswd(next)
	}
	return next
}

func (d Day11) nextPswd(pswd string) string {
	next := []rune(pswd)
	for ix := len(pswd) - 1; ix >= 0; ix-- {
		nl := d.nextPswdLetter(next[ix])
		next[ix] = nl
		if nl != 'a' {
			break
		}
	}

	return string(next)
}

func (d Day11) nextPswdLetter(ch rune) rune {
	switch ch {
	case 'z':
		return 'a'
	case 'h', 'k', 'n':
		return rune(2 + int(ch))
	default:
		return rune(1 + int(ch))
	}
}

// Passwords must include one increasing straight of at least three letters,
// like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
func (d Day11) hasLetterSequence(pswd string) bool {
	count, current := 1, 0
	for _, ch := range []rune(pswd) {
		if (current + 1) == int(ch) {
			count += 1
			if count >= 3 {
				return true
			}
		} else {
			count = 1
		}

		current = int(ch)
	}
	return false
}

// Passwords may not contain the letters i, o, or l, as these letters
// can be mistaken for other characters and are therefore confusing.
var disallowedPswdRegex = regexp.MustCompile(`[^iol]+`)

func (d Day11) isPswdDisallowed(pswd string) bool {
	return !disallowedPswdRegex.MatchString(pswd)
}

// Passwords must contain at least two different, non-overlapping pairs of letters,like aa, bb, or zz.
func (d Day11) hasDistinctPairs(pswd string) bool {
	firstIndex := util.FindLetterPairIndex(pswd)
	if firstIndex < 0 {
		return false
	}

	return util.HasLetterPair(pswd[firstIndex+2:])
}
