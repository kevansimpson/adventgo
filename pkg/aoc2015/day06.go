package aoc2015

import (
	"math"
	"regexp"

	"github.com/kevansimpson/util"
)

/**
 * <a href="https://adventofcode.com/2015/day/6">Day 6</a>
 */

func followLightCommands(input []string) (int, int) {
	cmds := parseCmds(input)
	return flipLights(cmds, flip), flipLights(cmds, brighten)
}

func flipLights(cmds []LightCmd, fxn func(string, int) int) int {
	var grid [1000][1000]int
	for i := range grid {
		var row [1000]int
		grid[i] = row
	}

	for _, cmd := range cmds {
		for x := cmd.start.X; x <= cmd.end.X; x += 1 {
			for y := cmd.start.Y; y <= cmd.end.Y; y += 1 {
				grid[y][x] = fxn(cmd.cmd, grid[y][x])
			}
		}
	}

	sum := 0
	for x := 0; x < 1000; x += 1 {
		for y := 0; y < 1000; y += 1 {
			sum += grid[y][x]
		}
	}

	return sum
}

func flip(cmd string, light int) int {
	switch cmd {
	case "toggle":
		if light == 0 {
			return 1
		} else {
			return 0
		}
	case "turn on":
		return 1
	case "turn off":
		return 0
	default:
		return math.MaxInt
	}
}

func brighten(cmd string, light int) int {
	switch cmd {
	case "toggle":
		return light + 2
	case "turn on":
		return light + 1
	case "turn off":
		dim := light - 1
		if dim > 0 {
			return dim
		} else {
			return 0
		}
	default:
		return math.MaxInt
	}
}

type LightCmd struct {
	cmd   string
	start util.Point
	end   util.Point
}

func parseCmds(input []string) []LightCmd {
	regex := regexp.MustCompile(`(toggle|turn on|turn off) ([\d,]+) through ([\d,]+)`)
	var cmds []LightCmd
	for _, str := range input {
		m := regex.FindStringSubmatch(str)
		cmds = append(cmds, LightCmd{m[1], util.MakePoint(m[2]), util.MakePoint(m[3])})
	}

	return cmds
}
