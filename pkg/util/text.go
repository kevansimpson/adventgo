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

var numberRegex = regexp.MustCompile(`[-]?\d*`)

// Extracts all integers and returns them in a slice.
func ExtractInts(str string) []int {
	m := numberRegex.FindAllString(str, -1)
	var num []int
	for _, s := range m {
		i, _ := strconv.ParseInt(s, 10, 32)
		num = append(num, int(i))
	}
	return num
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
