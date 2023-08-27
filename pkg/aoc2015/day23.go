package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/23">Day 23</a>
 */

import (
	"strconv"
	"strings"
	"sync"
)

type Day23 struct{}

func (d Day23) followComputerInstructions(input []string) (int, int) {
	var wg sync.WaitGroup
	channel0 := make(chan int)
	channel1 := make(chan int)
	wg.Add(2)

	go d.solve(input, 0, channel0, &wg)
	go d.solve(input, 1, channel1, &wg)

	register0, register1 := <-channel0, <-channel1

	close(channel0)
	close(channel1)
	wg.Wait()

	return register0, register1
}

func (d Day23) solve(instructions []string, a int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var registers = map[string]int{
		"a": a,
		"b": 0}
	d.updateRegister(registers, instructions, 0)
	ch <- registers["b"]
}

func (d Day23) updateRegister(registers map[string]int, instructions []string, index int) {
	if index < 0 || index >= len(instructions) {
		return
	}

	tokens := strings.Split(strings.ReplaceAll(instructions[index], ",", ""), " ")
	switch tokens[0] {
	case "hlf":
		registers[tokens[1]] = registers[tokens[1]] / 2
		d.updateRegister(registers, instructions, index+1)
	case "tpl":
		registers[tokens[1]] = registers[tokens[1]] * 3
		d.updateRegister(registers, instructions, index+1)
	case "inc":
		registers[tokens[1]] = registers[tokens[1]] + 1
		d.updateRegister(registers, instructions, index+1)
	case "jmp":
		dist, _ := strconv.Atoi(tokens[1])
		d.updateRegister(registers, instructions, index+dist)
	case "jie":
		if (registers[tokens[1]] % 2) == 0 {
			even, _ := strconv.Atoi(tokens[2])
			d.updateRegister(registers, instructions, index+even)
		} else {
			d.updateRegister(registers, instructions, index+1)
		}
	case "jio":
		if registers[tokens[1]] == 1 {
			odd, _ := strconv.Atoi(tokens[2])
			d.updateRegister(registers, instructions, index+odd)
		} else {
			d.updateRegister(registers, instructions, index+1)
		}
	}
}
