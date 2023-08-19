package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/3">Day 3</a>
 */

import "github.com/kevansimpson/util"

type void struct{}

var member void

type Set map[util.Point]void

func addPointToSet(set Set, elem util.Point) {
	set[elem] = member
}

func santaRoute(input string) int {
	set := make(map[util.Point]void)
	return follow(set, input, 0, 1)
}

func roboSantaRoute(input string) int {
	set := make(map[util.Point]void)
	follow(set, input, 0, 2)
	return follow(set, input, 1, 2)
}

func follow(set Set, directions string, start int, increment int) int {
	location := util.ORIGIN
	addPointToSet(set, util.ORIGIN)
	dirs := []rune(directions)
	for i := start; i < len(dirs); i += increment {
		location = location.RuneStep(dirs[i])
		addPointToSet(set, location)
	}
	return len(set)
}
