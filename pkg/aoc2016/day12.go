package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/12">Day 12</a>
 */

import (
	"strconv"
	"strings"
	"sync"
)

type Day12 struct{}

func (d Day12) computerInstructions(input []string) (int, int) {
	var wg sync.WaitGroup
	channel1, channel2 := make(chan int), make(chan int)
	wg.Add(2)

	go d.operateAssembunnyCode(input, 0, channel1, &wg)
	go d.operateAssembunnyCode(input, 1, channel2, &wg)

	code0, code1 := <-channel1, <-channel2
	close(channel1)
	close(channel2)
	wg.Wait()

	return code0, code1
}

func (d Day12) operateAssembunnyCode(input []string, c int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	register := map[string]int{"a": 0, "b": 0, "c": c, "d": 0}
	ix := 0
	for ix < len(input) {
		jump := 1
		bits := strings.Split(input[ix], " ")
		switch bits[0] {
		case "cpy":
			register[bits[2]] = d.resolve(register, bits[1])
			break
		case "jnz":
			x := d.resolve(register, bits[1])
			if x != 0 {
				j, _ := strconv.Atoi(bits[2])
				jump = j
			}
			break
		case "inc":
			d.bump(register, bits[1], 1)
			break
		case "dec":
			d.bump(register, bits[1], -1)
			break
		}

		ix += jump
	}

	ch <- register["a"]
}

func (d Day12) bump(register map[string]int, key string, bump int) {
	v, hasV := register[key]
	if hasV {
		register[key] = v + bump
	} else {
		register[key] = bump
	}
}

func (d Day12) resolve(register map[string]int, key string) int {
	x, hasX := register[key]
	if !hasX {
		xx, _ := strconv.Atoi(key)
		x = xx
	}

	return x
}
