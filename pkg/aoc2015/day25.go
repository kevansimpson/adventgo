package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/25">Day 25</a>
 */

type Day25 struct{}

func (d Day25) foo(input string) int {
	return 1
}

func (d Day25) readMachineConsole(initialCode int, row int, column int) int {
	index := column - 1
	for i := 0; i < (row + column - 1); i++ {
		index += i
	}

	code := initialCode
	for i := 0; i < index; i++ {
		code = d.nextCode(code)
	}
	return code
}

func (d Day25) nextCode(num int) int {
	return (num * 252533) % 33554393
}
