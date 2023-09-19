package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/2">Day 2</a>
 */

type Program struct {
	Codes []int
}

func (p Program) runProgram() int {
	index := 0

Loop:
	for index < len(p.Codes) {
		opCode := p.Codes[index]
		// fmt.Println(opCode)

		switch opCode {
		case 1:
			// fmt.Printf("add  %v\n", p.Codes[index:index+4])
			index = p.add(index)
		case 2:
			// fmt.Printf("mult %v\n", p.Codes[index:index+4])
			index = p.multiply(index)
			break
		case 99:
			// fmt.Printf("done @ %d\n", index)
			break Loop
		}
		// fmt.Printf("%v\n", p.Codes)
	}

	return p.Codes[0]
}

func (p Program) add(index int) int {
	a, b, c := p.Codes[index+1], p.Codes[index+2], p.Codes[index+3] // raw
	if true {                                                       // indexes
		a = p.Codes[a]
		b = p.Codes[b]
		// c = p.Codes[c]
	}
	// fmt.Printf("add  %d %d %d\n", a, b, c)
	p.Codes[c] = a + b
	return index + 4
}

func (p Program) multiply(index int) int {
	a, b, c := p.Codes[index+1], p.Codes[index+2], p.Codes[index+3] // raw
	if true {                                                       // indexes
		a = p.Codes[a]
		b = p.Codes[b]
		// c = p.Codes[c]
	}
	// fmt.Printf("mult %d %d %d\n", a, b, c)
	p.Codes[c] = a * b
	return index + 4
}

func (p Program) gravityAssist(position1 int, position2 int) {
	p.Codes[1] = position1
	p.Codes[2] = position2
}

// creates new program with fresh copy of int codes
func CreateProgram(intCodes []int) Program {
	doppleganger := make([]int, len(intCodes))
	copy(doppleganger, intCodes)
	return Program{Codes: doppleganger}
}
