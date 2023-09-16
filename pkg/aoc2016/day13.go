package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/13">Day 13</a>
 */

import (
	"strconv"

	"github.com/kevansimpson/util"
)

type Day13 struct{}

func (d Day13) cubicleMaze(input int) (int, int) {
	depth := -1
	depthMap := make(map[util.Point]int)
	search := []util.Point{{X: 1, Y: 1}}

maze:
	for len(search) > 0 {
		depth += 1
		current := make([]util.Point, len(search))
		copy(current, search)
		search = make([]util.Point, 0)

		for _, pt := range current {
			_, seen := depthMap[pt]
			if seen {
				continue
			} else {
				if d.isWall(pt, input) {
					depthMap[pt] = -1
				} else {
					depthMap[pt] = depth
				}
			}

			if pt.X == 31 && pt.Y == 39 {
				break maze
			}

			for _, dir := range "UDLR" {
				next := pt.RuneStep(dir)
				if d.isOpenSpace(next, input) {
					search = append(search, next)
				}
			}
		}
	}

	count := 0
	for _, v := range depthMap {
		if v <= 50 {
			count += 1
		}
	}
	return depth, count
}

func (d Day13) isOpenSpace(pt util.Point, input int) bool {
	return pt.X >= 0 && pt.Y >= 0 && !d.isWall(pt, input)
}

func (d Day13) isWall(pt util.Point, input int) bool {
	ones := 0
	for _, bit := range d.toBinary(pt, input) {
		if bit == '1' {
			ones += 1
		}
	}

	return (ones % 2) != 0
}

func (d Day13) toBinary(pt util.Point, input int) string {
	return strconv.FormatInt(int64(pt.X*pt.X+3*pt.X+2*pt.X*pt.Y+pt.Y+pt.Y*pt.Y+input), 2)
}
