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

func calculateSantaRoutes(input []string) (int, int) {
	sleigh := buildSleigh(input)
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

type jumpKey struct {
	departs string
	arrives string
}

type santaSleigh struct {
	locations util.Set[string]
	jumps     map[jumpKey]int
	distances map[string]int
}

func buildSleigh(input []string) santaSleigh {
	regex := regexp.MustCompile(`(.+)\sto\s(.+)\s=\s(\d+)`)
	locations := make(util.Set[string])
	jumps := make(map[jumpKey]int)
	distances := make(map[string]int)

	for _, str := range input {
		m := regex.FindStringSubmatch(str)
		city1, city2 := m[1], m[2]
		dist, _ := strconv.Atoi(m[3])
		util.Add(locations, city1)
		util.Add(locations, city2)
		jumps[jumpKey{city1, city2}] = dist
		jumps[jumpKey{city2, city1}] = dist
	}

	routes := util.Permutations(util.SetToSlice(locations))
	for _, r := range routes {
		distances[fmt.Sprintf("%s", r)] = calculateDistance(r, jumps)
	}

	return santaSleigh{locations, jumps, distances}
}

func calculateDistance(path []string, jumps map[jumpKey]int) int {
	dist, max := 0, len(path)-1
	for ix := 0; ix < max; ix += 1 {
		d, ok := jumps[jumpKey{path[ix], path[ix+1]}]
		if ok {
			dist += d
		}
	}

	return dist
}
