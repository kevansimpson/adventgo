package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/1">Day 1</a>
 */

import "strings"

func whatFloorSanta(input string) int {
	return len(strings.ReplaceAll(input, ")", "")) - len(strings.ReplaceAll(input, "(", ""))
}

func santaEntersBasement(input string) int {
	var floor = 0
	var position = 0
	for _, ch := range []rune(input) {
		position += 1
		switch ch {
		case ')':
			floor -= 1
		case '(':
			floor += 1
		}
		if floor < 0 {
			return position
		}
	}
	return -1
}
