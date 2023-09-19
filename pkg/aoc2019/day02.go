package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/2">Day 2</a>
 */

type Day02 struct{}

func (d Day02) runProgram(input []int) (int, int) {
	p := CreateProgram(input)
	p.gravityAssist(12, 2)
	return p.runProgram(), d.targetOutput(input)
}

func (d Day02) targetOutput(input []int) int {
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			p := CreateProgram(input)
			p.gravityAssist(n, v)
			if 19690720 == p.runProgram() {
				return 100*n + v
			}
		}
	}
	return -1
}
