package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/2">Day 2</a>
 * <a href="https://adventofcode.com/2019/day/5">Day 5</a>
 * <a href="https://adventofcode.com/2019/day/7">Day 7</a>
 * <a href="https://adventofcode.com/2019/day/9">Day 9</a>
 */

import (
	"fmt"
	"strconv"
)

type Parameters struct {
	program    Program
	fullOpCode string
	opCode     int
	next       int
}

func (p *Parameters) compare(condition bool) int {
	if condition {
		return 1
	} else {
		return 0
	}
}

func (p *Parameters) jump(condition bool, distance int) {
	if condition {
		p.next = distance
	}
}

func (p Parameters) a() int {
	return p.read(1, 2)
}

func (p Parameters) b() int {
	return p.read(2, 1)
}

func (p Parameters) c1() int {
	return p.write(1, 2)
}

func (p Parameters) c3() int {
	return p.write(3, 0)
}

func (p Parameters) read(offset int, opCodeIndex int) int {
	switch p.fullOpCode[opCodeIndex] {
	case '0':
		return p.program.get(p.get(offset))
	case '1':
		return p.get(offset)
	case '2':
		return p.program.get(p.get(offset) + p.program.RelativeBase)
	}
	panic("Parameters.read")
}

func (p Parameters) write(offset int, opCodeIndex int) int {
	switch p.fullOpCode[opCodeIndex] {
	case '0', '1':
		return p.get(offset)
	case '2':
		return p.get(offset) + p.program.RelativeBase
	}
	panic("Parameters.write")
}

func (p Parameters) get(offset int) int {
	return p.program.get(p.program.Index + offset)
}

func (p Program) NewParameters() Parameters {
	baseOpCode := p.Codes[p.Index]
	fullOpCode := fmt.Sprintf("%05d", baseOpCode)
	opCode, _ := strconv.Atoi(fullOpCode[3:])
	next := p.Index
	switch opCode {
	case 1, 2, 7, 8:
		next += 4
	case 3, 4, 9:
		next += 2
	case 5, 6:
		next += 3
	}

	return Parameters{program: p, fullOpCode: fullOpCode, opCode: opCode, next: next}
}
