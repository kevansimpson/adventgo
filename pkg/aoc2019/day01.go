package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/1">Day 1</a>
 */

type Day01 struct{}

func (d Day01) sumFuelRequirements(input []int) (int, int) {
	requiredFuel, accumulatedFuel := 0, 0
	for _, num := range input {
		requiredFuel += d.calculate(num)
		accumulatedFuel += d.accumulate(num)
	}
	return requiredFuel, accumulatedFuel
}

func (d Day01) accumulate(mass int) int {
	fuel, total := d.calculate(mass), 0
	for fuel > 0 {
		total += fuel
		fuel = d.calculate(fuel)
	}

	return total
}

func (d Day01) calculate(mass int) int {
	return (mass / 3) - 2
}
