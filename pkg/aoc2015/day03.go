package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/3">Day 3</a>
 */

import "github.com/kevansimpson/util"

type Day03 struct{}

func (d Day03) santaRoute(input string) int {
	set := make(util.Set[util.Point])
	return d.follow(set, input, 0, 1)
}

func (d Day03) roboSantaRoute(input string) int {
	set := make(util.Set[util.Point])
	d.follow(set, input, 0, 2)
	return d.follow(set, input, 1, 2)
}

func (d Day03) follow(set util.Set[util.Point], directions string, start int, increment int) int {
	location := util.ORIGIN
	util.Add(set, util.ORIGIN)
	dirs := []rune(directions)
	for i := start; i < len(dirs); i += increment {
		location = location.RuneStep(dirs[i])
		util.Add(set, location)
	}
	return len(set)
}
