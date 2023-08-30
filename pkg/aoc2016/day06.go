package aoc2016

import (
	"strings"

	"github.com/kevansimpson/util"
)

/**
 * <a href="https://adventofcode.com/2016/day/6">Day 6</a>
 */

type Day06 struct{}

func (d Day06) readRepetitionCode(input []string) (string, string) {
	columns := util.Columns(input)
	counts := d.columnCounts(columns)
	correctedVersion, originalMessage := make([]rune, len(counts)), make([]rune, len(counts))
	for i, countMap := range counts {
		min, max := 50, 0
		for str, c := range countMap {
			if c < min {
				originalMessage[i] = str
				min = c
			}
			if c > max {
				correctedVersion[i] = str
				max = c
			}
		}
	}

	return string(correctedVersion), string(originalMessage)
}

type ColumnCounts []map[rune]int

func (d Day06) columnCounts(columns []string) ColumnCounts {
	counts := make(ColumnCounts, len(columns))
	for x, col := range columns {
		tally := make(map[rune]int, 0)
		for _, ch := range col {
			tally[ch] = strings.Count(col, string(ch))
		}
		counts[x] = tally
	}

	return counts
}
