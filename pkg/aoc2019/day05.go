package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/5">Day 5</a>
 */

type Day05 struct{}

func (d Day05) runProgram(input []int) (int, int) {
	out1, out5 := make(chan int, 12), make(chan int, 2)
	p1 := CreateProgram(input, SingleInputChannel(1), out1)
	p5 := CreateProgram(input, SingleInputChannel(5), out5)

	p1.runProgram()
	p5.runProgram()
	diagCode1, diagCode5 := -1, -1

	for out1 != nil || out5 != nil {
		select {
		case v, ok := <-out1:
			if !ok || v != 0 {
				diagCode1 = v
				close(out1)
				out1 = nil
			}

		case v, ok := <-out5:
			if !ok || v != 0 {
				diagCode5 = v
				close(out5)
				out5 = nil
			}
		}
	}

	return diagCode1, diagCode5
}
