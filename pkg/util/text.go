package util

import (
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile(`[-]?\d*`)

// Extracts all integers and returns them in a slice.
func ExtractInts(str string) []int {
	m := numberRegex.FindAllString(str, -1)
	num := make([]int, 0)
	for _, s := range m {
		i, _ := strconv.ParseInt(s, 10, 32)
		num = append(num, int(i))
	}
	return num
}
