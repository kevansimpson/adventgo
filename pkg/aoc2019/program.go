package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/2">Day 2</a>
 */

import "sync"

type Program struct {
	Codes  map[int]int
	Input  chan int
	Output chan int

	Index        int
	RelativeBase int
}

func (p Program) runProgramAsync(wg *sync.WaitGroup) {
	defer wg.Done()
	p.runProgram()
}

func (p *Program) runProgram() int {

Loop:
	for p.Codes[p.Index] != 99 {
		params := p.NewParameters()

		switch params.opCode {
		case 1:
			p.assign(params.c3(), params.a()+params.b())
		case 2:
			p.assign(params.c3(), params.a()*params.b())
		case 3:
			p.assign(params.c1(), <-p.Input)
		case 4:
			p.Output <- params.a()
		case 5:
			params.jump(params.a() != 0, params.b())
		case 6:
			params.jump(params.a() == 0, params.b())
		case 7:
			p.assign(params.c3(), params.compare(params.a() < params.b()))
		case 8:
			p.assign(params.c3(), params.compare(params.a() == params.b()))
		case 9:
			p.RelativeBase += params.a()
		case 99:
			break Loop
		}
		p.Index = params.next
	}

	return p.Codes[0]
}

func (p Program) assign(index int, value int) {
	p.Codes[index] = value
}

func (p Program) get(index int) int {
	code, has := p.Codes[index]
	if has {
		return code
	} else {
		return 0
	}
}

// creates new program with fresh copy of int codes
func CreateProgram(intCodes []int, input chan int, output chan int) Program {
	doppleganger := make(map[int]int, len(intCodes))
	for i, c := range intCodes {
		doppleganger[i] = c
	}
	return Program{Codes: doppleganger, Input: input, Output: output, Index: 0, RelativeBase: 0}
}

// creates an input channel with a single value loaded
func SingleInputChannel(input int) chan int {
	c := make(chan int, 1)
	c <- input
	return c
}
