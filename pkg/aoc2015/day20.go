package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/20">Day 20</a>
 */

const maxHouses = 1000000

type Day20 struct{}

func (d Day20) lowestHouseNumbers(target int) (int, int) {
	houses1, houses2 := make([]int, maxHouses), make([]int, maxHouses)
	for elf := 1; elf < maxHouses; elf++ {
		count := 0
		for v := elf; v < maxHouses; v += elf {
			houses1[v] += elf * 10
			count++
			if count < 50 {
				houses2[v] += elf * 11
			}
		}
	}
	return d.findTargetHouse(houses1, target), d.findTargetHouse(houses2, target)
}

func (d Day20) findTargetHouse(houses []int, target int) int {
	for i := 0; i < maxHouses; i++ {
		if houses[i] >= target {
			return i
		}
	}
	return -1
}
