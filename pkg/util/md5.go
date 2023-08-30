package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"strings"
)

// Generates MD5 hash in hexidecimal from given string
func GenerateMD5inHex(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Finds the next MD5 hash in hexidecimal with the given prefix
func NextHash(input string, prefix string, start int) (string, int) {
	h := md5.New()
	for i := start; start <= math.MaxInt; i += 1 {
		// next := GenerateMD5inHex(input + fmt.Sprintf("%d", i))
		io.WriteString(h, fmt.Sprintf("%s%d", input, i))
		next := fmt.Sprintf("%x", h.Sum(nil))
		if strings.HasPrefix(next, prefix) {
			return next, i
		} else {
			h.Reset()
		}
	}
	return input, -1
}
