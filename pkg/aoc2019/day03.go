package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/3">Day 3</a>
 */

import (
	"math"
	"strconv"
	"strings"

	"github.com/kevansimpson/util"
)

type Day03 struct{}
type WirePath struct {
	Path  []util.Point
	Steps map[util.Point]int
}

func (d Day03) installCircuits(input []string) (int, int) {
	wp1, wp2 := d.runWires(input[0]), d.runWires(input[1])
	set := wp1.intersection(wp2)
	centralPort, fewestToIntersection := math.MaxInt, math.MaxInt
	for pt := range set {
		dist := pt.ManhattanDistanceOrigin()
		if dist < centralPort {
			centralPort = dist
		}

		steps := wp1.Steps[pt] + wp2.Steps[pt]
		if steps < fewestToIntersection {
			fewestToIntersection = steps
		}
	}

	return centralPort, fewestToIntersection
}

func (d Day03) runWires(wire string) WirePath {
	var path []util.Point
	steps := make(map[util.Point]int)
	pt := util.ORIGIN
	count := 0
	for _, w := range strings.Split(wire, ",") {
		dir := rune(w[0])
		dist, _ := strconv.Atoi(w[1:])
		for x := 1; x <= dist; x++ {
			pt = pt.RuneStep(dir)
			path = append(path, pt)
			count += 1
			_, hasCount := steps[pt]
			if !hasCount {
				steps[pt] = count
			}
		}
	}

	return WirePath{Path: path, Steps: steps}
}

func (wp WirePath) intersection(that WirePath) util.Set[util.Point] {
	return util.Intersection(
		util.SliceToSet(wp.Path), util.SliceToSet(that.Path))
}
