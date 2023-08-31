package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/7">Day 7</a>
 */

import (
	"fmt"
	"strings"

	"github.com/kevansimpson/util"
)

type Day07 struct{}

func (d Day07) tallySecureIps(input []string) (int, int) {
	tls, ssl := 0, 0
	for _, rawIPv7 := range input {
		ipv7List := strings.FieldsFunc(rawIPv7, d.brackets)
		if d.supportsTLS(ipv7List) {
			tls += 1
		}
		if d.supportsSSL(ipv7List) {
			ssl += 1
		}
	}

	return tls, ssl
}

func (d Day07) supportsTLS(ipv7List []string) bool {
	even := false
	for i, ipv7 := range ipv7List {
		if (i%2) == 0 && d.abba(ipv7) {
			even = true
		}
		if (i%2) > 0 && d.abba(ipv7) {
			return false
		}
	}
	return even
}

func (d Day07) supportsSSL(ipv7List []string) bool {
	abaSet := make(util.Set[string])
	oddList := make([]string, len(ipv7List)/2)

	for i, ipv7 := range ipv7List {
		if (i % 2) == 0 {
			loop := true
			str := ipv7
			for loop {
				a, b, hasLS := util.FindLetterSandwich(str)
				if hasLS {
					aba := fmt.Sprintf("%c%c%c", a, b, a)
					str = str[strings.Index(str, aba)+1:]
					if a != b {
						util.Add(abaSet, fmt.Sprintf("%c%c%c", b, a, b))
					}
				} else {
					loop = false
				}
			}

		} else {
			oddList[i/2] = ipv7
		}
	}

	if len(abaSet) > 0 {
		for _, odd := range oddList {
			for aba := range abaSet {
				if strings.Contains(odd, aba) {
					return true
				}
			}
		}
	}
	return false
}

func (d Day07) abba(str string) bool {
	b, hasLP := util.FindLetterPair(str)
	if hasLP {
		if b > 0 && len(str) > b+2 {
			a1, a2 := str[b-1], str[b+2]
			abba := a1 == a2 && a1 != str[b]
			if abba {
				return true
			}
		}
		return d.abba(str[b+1:])
	}
	return false
}

func (d Day07) brackets(r rune) bool {
	return r == '[' || r == ']'
}
