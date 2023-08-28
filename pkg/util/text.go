package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

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

// Generates MD5 hash in hexidecimal from given string
func GenerateMD5inHex(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Finds the next MD5 hash in hexidecimal with the given prefix
func NextHash(input string, prefix string, start int) (string, int) {
	for i := start; start <= math.MaxInt; i += 1 {
		next := GenerateMD5inHex(input + fmt.Sprintf("%d", i))
		if strings.HasPrefix(next, prefix) {
			return next, i
		}
	}
	return input, -1
}

func ReverseString(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
