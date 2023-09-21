package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/6">Day 6</a>
 */

import (
	"slices"
	"strings"

	"github.com/kevansimpson/util"
)

type Day06 struct{}

func (d Day06) mapOrbitalTransfers(input []string) (int, int) {
	orbits := d.mapOrbits(input)
	return d.totalOrbits(orbits), d.minimumTransfers(orbits)
}

func (d Day06) minimumTransfers(orbits map[string][]string) int {
	reversed := d.reverseOrbits(orbits)
	visited, planets := make(util.Set[string]), make(util.Set[string])
	util.Add(visited, "SAN")
	util.Add(planets, reversed["SAN"])

	count := 0
	for len(planets) > 0 && count < 360 {
		next := make(util.Set[string])
		for p := range planets {
			util.AddAll(next, orbits[p])
			util.Add(next, reversed[p])
			util.Add(visited, p)
		}

		for v := range visited {
			delete(next, v)
		}

		_, you := next["YOU"]
		if you {
			break
		} else {
			count += 1
		}
		planets = next
	}

	return count
}

func (d Day06) totalOrbits(orbits map[string][]string) int {
	sum := 0
	for p := range orbits {
		sum += d.countOrbits(orbits, p)
	}

	return sum
}

func (d Day06) countOrbits(orbits map[string][]string, planet string) int {
	planets, hasOrbits := orbits[planet]
	if hasOrbits {
		count := len(planets)
		for _, p := range planets {
			count += d.countOrbits(orbits, p)
		}
		return count
	} else {
		return 0
	}
}

func (d Day06) mapOrbits(input []string) map[string][]string {
	orbits := make(map[string][]string)
	slices.Sort(input)
	for _, o := range input {
		ab := strings.Split(o, ")")
		list, hasOrbits := orbits[ab[0]]
		if hasOrbits {
			orbits[ab[0]] = append(list, ab[1])
		} else {
			orbits[ab[0]] = []string{ab[1]}
		}
	}

	return orbits
}

func (d Day06) reverseOrbits(original map[string][]string) map[string]string {
	reversed := make(map[string]string)
	for p, orbits := range original {
		for _, o := range orbits {
			reversed[o] = p
		}
	}

	return reversed
}
