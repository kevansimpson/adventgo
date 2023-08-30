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
func FindLetterPairIndex(str string) int {
	ltrs := []rune(str)
	max := len(str) - 1
	for ix := 0; ix < max; ix += 1 {
		if ltrs[ix] == ltrs[ix+1] {
			return ix
		}
	}
	return -1
}

// Returns true if given string contains at least one letter that appears twice in a row.
func HasLetterPair(str string) bool {
	return FindLetterPairIndex(str) >= 0
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
