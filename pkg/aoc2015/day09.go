package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/9">Day 9</a>
 */

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/kevansimpson/util"
)

type Day09 struct{}

func (d Day09) calculateSantaRoutes(input []string) (int, int) {
	sleigh := d.buildSleigh(input)
	min, max := math.MaxInt, math.MinInt
	for _, v := range sleigh.distances {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}

type JumpKey struct {
	departs string
	arrives string
}

type SantaSleigh struct {
	locations util.Set[string]
	jumps     map[JumpKey]int
	distances map[string]int
}

func (d Day09) buildSleigh(input []string) SantaSleigh {
	regex := regexp.MustCompile(`(.+)\sto\s(.+)\s=\s(\d+)`)
	locations := make(util.Set[string])
	jumps := make(map[JumpKey]int)
	distances := make(map[string]int)

	for _, str := range input {
		m := regex.FindStringSubmatch(str)
		city1, city2 := m[1], m[2]
		dist, _ := strconv.Atoi(m[3])
		util.Add(locations, city1)
		util.Add(locations, city2)
		jumps[JumpKey{city1, city2}] = dist
		jumps[JumpKey{city2, city1}] = dist
	}

	routes := util.Permutations(util.SetToSlice(locations))
	for _, r := range routes {
		distances[fmt.Sprintf("%s", r)] = d.calculateDistance(r, jumps)
	}

	return SantaSleigh{locations, jumps, distances}
}

func (d Day09) calculateDistance(path []string, jumps map[JumpKey]int) int {
	dist, max := 0, len(path)-1
	for ix := 0; ix < max; ix += 1 {
		d, ok := jumps[JumpKey{path[ix], path[ix+1]}]
		if ok {
			dist += d
		}
	}

	return dist
}
