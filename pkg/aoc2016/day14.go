package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/14">Day 14</a>
 */

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/kevansimpson/util"
)

type Day14 struct{}

func (d Day14) find64thKeys(input string) (int, int) {
	key64 := util.FastOrFull(d.lookForCachedKey(input), func() int {
		return d.originalKeys(input)
	})
	stretchedKey64 := d.lookForCachedKey(fmt.Sprintf("stretched14_%s", input))
	return key64, stretchedKey64
}

func (d Day14) lookForCachedKey(input string) int {
	hashCache := util.ReadLines(fmt.Sprintf("data/%s.txt", input))
	keys := make([]int, 64)
	k := 0

	for index, hash := range hashCache {
		ltr, hasLT := util.FindLetterTriplet(hash)
		if hasLT {
			quintuple := strings.Repeat(string(hash[ltr]), 5)
			for i := index + 1; i < index+1000; i++ {
				if strings.Contains(hashCache[i], quintuple) {
					keys[k] = index
					k += 1
					if k >= 64 {
						return keys[63]
					}
					break
				}
			}
		}
	}

	return -1
}

func (d Day14) originalKeys(input string) int {
	keys := make([]int, 64)
	k := 0
	index := 0
	for k < 64 {
		nextHash, nextIndex := util.NextHash(input, index, util.HasLetterTriplet)
		if nextIndex >= 0 {
			if d.isKey(input, nextHash, nextIndex) {
				keys[k] = nextIndex
				k += 1
			}
			index = nextIndex + 1

		} else {
			break
		}
	}

	return keys[63]
}

func (d Day14) isKey(input string, hash string, index int) bool {
	trip, _ := util.FindLetterTriplet(hash)
	triple := hash[trip : trip+3]
	quintuple := triple + triple[0:2]
	_, matchAt := util.NextHashUntil(input, index+1, index+1000,
		func(s string) bool { return strings.Contains(s, quintuple) })
	return matchAt >= 0
}

func (d Day14) generateStretchedKeys(input string) {
	for i := 0; i < 25000; i++ {
		fmt.Println(d.stretch(fmt.Sprintf("%s%d", input, i)))
	}
}

func (d Day14) stretch(input string) string {
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%s", input))
	hex := fmt.Sprintf("%x", h.Sum(nil))
	h.Reset()
	for i := 0; i < 2016; i += 1 {
		io.WriteString(h, fmt.Sprintf("%s", hex))
		hex = fmt.Sprintf("%x", h.Sum(nil))
		h.Reset()
	}

	return hex
}
