package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/19">Day 19</a>
 */

import (
	"strings"
	"unicode"

	"github.com/kevansimpson/util"
)

type Day19 struct{}

func (d Day19) createMolecules(input []string) (int, int) {
	last := len(input) - 1
	medicine := input[last]
	rmap := d.buildReplacementMap(input[0 : last-1])

	return d.applyAllReplacements(rmap, medicine), d.shortestPath(medicine)
}

func (d Day19) applyAllReplacements(rmap map[string][]string, medicine string) int {
	molecules := make(util.Set[string], 0)

	for key, reps := range rmap {
		unique := d.applyReplacement(medicine, key, reps)
		util.AddSet(molecules, unique)
	}

	return len(molecules)
}

func (d Day19) applyReplacement(chain string, key string, replacements []string) util.Set[string] {
	unique := make(util.Set[string], 0)
	klen := len(key)

	// fmt.Printf("%s - %v", key, replacements)
	for _, r := range replacements {
		index := strings.Index(chain, key)
		for index >= 0 {
			rest := chain[index+klen:]
			variation := chain[0:index] + r + rest
			util.Add(unique, variation)
			next := strings.Index(rest, key)
			if next >= 0 {
				index += next + klen
			} else {
				index = next
			}
		}
	}

	return unique
}

// credit to https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/cy4h7ji/
func (d Day19) shortestPath(medicine string) int {
	upper := 0
	for _, ch := range medicine {
		if unicode.IsUpper(ch) {
			upper += 1
		}
	}

	return upper - strings.Count(medicine, "Rn") - strings.Count(medicine, "Ar") - 2*strings.Count(medicine, "Y") - 1
}

func (d Day19) buildReplacementMap(input []string) map[string][]string {
	rmap := make(map[string][]string, 0)
	for _, str := range input {
		tokens := strings.Split(str, " => ")
		list, has := rmap[tokens[0]]
		if has {
			rmap[tokens[0]] = append(list, tokens[1])
		} else {
			rmap[tokens[0]] = []string{tokens[1]}
		}
	}

	return rmap
}
