package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/9">Day 9</a>
 */

type Day09 struct{}

func (d Day09) runProgram(input []int) (int, int) {
	out1, out2 := make(chan int, 2000), make(chan int, 2000)
	p1 := CreateProgram(input, SingleInputChannel(1), out1)
	p2 := CreateProgram(input, SingleInputChannel(2), out2)

	p1.runProgram()
	p2.runProgram()

	diagCode1, diagCode2 := -1, -1
	for out1 != nil || out2 != nil {
		select {
		case v, _ := <-out1:
			diagCode1 = v
			close(out1)
			out1 = nil

		case v, _ := <-out2:
			diagCode2 = v
			close(out2)
			out2 = nil
		}
	}

	return diagCode1, diagCode2
}
