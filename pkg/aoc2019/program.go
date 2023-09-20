package aoc2019

import (
	"fmt"
	"strconv"
)

/**
 * <a href="https://adventofcode.com/2019/day/2">Day 2</a>
 */

type Program struct {
	Codes  []int
	Input  <-chan int
	Output chan<- int
}

func (p Program) runProgram() int {
	index := 0

Loop:
	for index < len(p.Codes) {
		baseOpCode := p.Codes[index]
		fullOpCode := fmt.Sprintf("%04d", baseOpCode)
		opCode, _ := strconv.Atoi(fullOpCode[2:])

		switch opCode {
		case 1:
			index = p.add(index, fullOpCode)
		case 2:
			index = p.multiply(index, fullOpCode)
		case 3:
			index = p.getInput(index, fullOpCode)
		case 4:
			index = p.setOutput(index, fullOpCode)
		case 5:
			index = p.jumpIfTrue(index, fullOpCode)
		case 6:
			index = p.jumpIfFalse(index, fullOpCode)
		case 7:
			index = p.lessThan(index, fullOpCode)
		case 8:
			index = p.equalTo(index, fullOpCode)
		case 99:
			break Loop
		}
	}

	return p.Codes[0]
}

func (p Program) paramMode(index int, fullOpCode string) (int, int, int) {
	a, b, c := p.Codes[index+1], p.Codes[index+2], p.Codes[index+3] // immediate mode
	// parameter modes
	if fullOpCode[1] == '0' {
		a = p.Codes[a]
	}
	if fullOpCode[0] == '0' {
		b = p.Codes[b]
	}

	return a, b, c // Parameters that an instruction writes to will never be in immediate mode.
}

func (p Program) singleParam(index int, opCodePlace string) int {
	a := p.Codes[index]
	if opCodePlace == "0" {
		a = p.Codes[a]
	}
	return a
}

func (p Program) add(index int, fullOpCode string) int {
	a, b, c := p.paramMode(index, fullOpCode)
	p.Codes[c] = a + b
	return index + 4
}

func (p Program) multiply(index int, fullOpCode string) int {
	a, b, c := p.paramMode(index, fullOpCode)
	p.Codes[c] = a * b
	return index + 4
}

func (p Program) getInput(index int, fullOpCode string) int {
	a := p.Codes[index+1]
	b := <-p.Input
	p.Codes[a] = b
	return index + 2
}

func (p Program) setOutput(index int, fullOpCode string) int {
	p.Output <- p.singleParam(index+1, fullOpCode[1:2])
	return index + 2
}

func (p Program) jumpIfTrue(index int, fullOpCode string) int {
	a, b, _ := p.paramMode(index, fullOpCode)
	if a != 0 {
		return b
	} else {
		return index + 3
	}
}

func (p Program) jumpIfFalse(index int, fullOpCode string) int {
	a, b, _ := p.paramMode(index, fullOpCode)
	if a == 0 {
		return b
	} else {
		return index + 3
	}
}

func (p Program) lessThan(index int, fullOpCode string) int {
	a, b, c := p.paramMode(index, fullOpCode)
	if a < b {
		p.Codes[c] = 1
	} else {
		p.Codes[c] = 0
	}
	return index + 4
}

func (p Program) equalTo(index int, fullOpCode string) int {
	a, b, c := p.paramMode(index, fullOpCode)
	if a == b {
		p.Codes[c] = 1
	} else {
		p.Codes[c] = 0
	}
	return index + 4
}

func (p Program) gravityAssist(position1 int, position2 int) {
	p.Codes[1] = position1
	p.Codes[2] = position2
}

// creates new program with fresh copy of int codes
func CreateProgram(intCodes []int, input <-chan int, output chan<- int) Program {
	doppleganger := make([]int, len(intCodes))
	copy(doppleganger, intCodes)
	return Program{Codes: doppleganger, Input: input, Output: output}
}

func SingleInputChannel(input int) <-chan int {
	c := make(chan int, 1)
	go func() {
		defer close(c)
		c <- input
	}()
	return c
}
