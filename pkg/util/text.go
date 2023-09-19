package util

import (
	"regexp"
	"strconv"
	"strings"
)

// Converts rows of strings to columns; assumes all rows are same length.
func Columns(rows []string) []string {
	cols := make([]string, len(rows[0]))
	for _, row := range rows {
		for x, r := range row {
			cols[x] = cols[x] + string(r)
		}
	}

	return cols
}

var numberRegex = regexp.MustCompile(`[-]?\d+`)

// Extracts all integers and returns them in a slice.
func ExtractInts(str string) []int {
	m := numberRegex.FindAllString(str, -1)
	var num []int
	for _, s := range m {
		i, _ := strconv.Atoi(s)
		num = append(num, int(i))
	}
	return num
}

// Returns the index if at least one letter that appears twice in a row in the given string, else returns -1.
func FindLetterPair(str string) (int, bool) {
	ltrs := []rune(str)
	max := len(str) - 1
	for ix := 0; ix < max; ix += 1 {
		if ltrs[ix] == ltrs[ix+1] {
			return ix, true
		}
	}
	return -1, false
}

// Returns the index if at least one letter that appears thrice in a row in the given string, else returns -1.
func FindLetterTriplet(str string) (int, bool) {
	ltrs := []rune(str)
	max := len(str) - 2
	for ix := 0; ix < max; ix += 1 {
		if ltrs[ix] == ltrs[ix+1] && ltrs[ix+1] == ltrs[ix+2] {
			return ix, true
		}
	}
	return -1, false
}

// Returns (A,B,true) if given string contains ABA, else (_,_,false).
func FindLetterSandwich(str string) (rune, rune, bool) {
	ltrs := []rune(str)
	max := len(str) - 2
	for ix := 0; ix < max; ix += 1 {
		if ltrs[ix] == ltrs[ix+2] {
			return ltrs[ix], ltrs[ix+1], true
		}
	}
	return NO_MATCH_RUNE, NO_MATCH_RUNE, false
}

// Returns true if given string contains at least one letter that appears twice in a row.
func HasLetterPair(str string) bool {
	_, hasLP := FindLetterPair(str)
	return hasLP
}

// Returns true if given string contains at least one letter that appears thrice in a row.
func HasLetterTriplet(str string) bool {
	_, hasLT := FindLetterTriplet(str)
	return hasLT
}

const NO_MATCH_RUNE = rune(0)

// Returns (A,B,true) if given string contains ABA, else (_,_,false).
func HasLetterSandwich(str string) bool {
	_, _, hasLS := FindLetterSandwich(str)
	return hasLS
}

func IndexAfter(str string, sub string, index int) int {
	found := strings.Index(str[index:], sub)
	if found >= 0 {
		found += index
	}
	return found
}

func ReverseString(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Rotates letters through the alphabet shift number of times
func ShiftText(str string, shift int) string {
	next := make([]rune, len(str))
	for i, ch := range str {
		ltr := strings.IndexRune(alphabet, ch)
		if ltr < 0 {
			next[i] = ch
		} else {
			ix := (ltr + shift) % 26
			next[i] = rune(alphabet[ix]) // ??
		}
	}
	return string(next)
}
