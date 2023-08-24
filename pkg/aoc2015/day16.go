package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/16">Day 16</a>
 */

import (
	"regexp"
	"strconv"
)

type Day16 struct{}

type AuntSue map[string]int

const INDEX = "INDEX"

var tickerTape = map[string]int{
	"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0,
	"vizslas": 0, "goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1}

func (d Day16) identifyAuntSue(input []string) (int, int) {
	var gift map[string]int
	var real map[string]int
	for _, auntSue := range input {
		attr := d.parseSue(auntSue)
		if d.hasSameAttr(attr) {
			gift = attr
		}
		if d.hasSameRange(attr) {
			real = attr
		}
	}

	return gift[INDEX], real[INDEX]
}

func (d Day16) hasSameAttr(sue AuntSue) bool {
	return d.hasSameValues(tickerTape, sue) && d.hasSameValues(sue, tickerTape)
}

func (d Day16) hasSameValues(sue1 AuntSue, sue2 AuntSue) bool {
	for k, v := range sue1 {
		if INDEX == k {
			continue
		}

		two, has2 := sue2[k]
		if has2 && v != two {
			return false
		}
	}
	return true
}

func (d Day16) hasSameRange(sue AuntSue) bool {
	return d.satisfiesTicker(sue) && d.reverseTicker(sue)
}

func (d Day16) satisfiesTicker(sue AuntSue) bool {
	for k, v := range tickerTape {
		count, has := sue[k]
		switch k {
		case INDEX:
			continue
		case "tree", "cats":
			if has && v >= count {
				return false
			}
		case "pomeranians", "goldfish":
			if has && v <= count {
				return false
			}
		default:
			if has && v != count {
				return false
			}
		}
	}
	return true
}

func (d Day16) reverseTicker(sue AuntSue) bool {
	for k, v := range sue {
		switch k {
		case INDEX, "trees", "cats", "pomeranians", "goldfish":
			continue
		default:
			tt, has := tickerTape[k]
			if has && v != tt {
				return false
			}
		}
	}
	return true
}

var regex = regexp.MustCompile(`Sue (\d+): (.+): (\d+), (.+): (\d+), (.+): (\d+)`)

func (d Day16) parseSue(input string) AuntSue {
	m := regex.FindStringSubmatch(input)
	index, _ := strconv.Atoi(m[1])
	pet1, _ := strconv.Atoi(m[3])
	pet2, _ := strconv.Atoi(m[5])
	pet3, _ := strconv.Atoi(m[7])

	return map[string]int{INDEX: index, m[2]: pet1, m[4]: pet2, m[6]: pet3}

}
