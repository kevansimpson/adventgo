package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/1">Day 1</a>
 */

import (
	"strconv"
	"strings"

	"github.com/kevansimpson/util"
)

type Day01 struct{}

func (d Day01) visitEasterBunnyHQ(input string) (int, int) {
	directions := strings.Split(input, ", ")
	nesw := []string{"^", ">", "v", "<"}
	facing := 0 // up/north
	visited := make(util.Set[util.Point])
	util.Add(visited, util.ORIGIN)
	twice := util.ORIGIN

	next := util.ORIGIN
	for _, dir := range directions {
		dist, _ := strconv.Atoi(dir[1:])
		switch dir[0] {
		case 'L':
			facing = (4 + facing - 1) % 4
		case 'R':
			facing = (4 + facing + 1) % 4
		}
		for i := 0; i < dist; i++ {
			next = next.Advance(nesw[facing], 1)
			_, alreadyVisited := visited[next]
			if alreadyVisited && twice == util.ORIGIN {
				twice = next
			} else {
				util.Add(visited, next)
			}
		}
	}

	return next.ManhattanDistanceOrigin(), twice.ManhattanDistanceOrigin()
}
