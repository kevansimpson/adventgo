package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/7">Day 7</a>
 */

import (
	"regexp"
	"strconv"
	"strings"
)

type Day07 struct{}

func (d Day07) assembleCircuits(input []string) (int, int) {
	circuits := d.wireCircuits(input)
	a := circuits.calculate("a", make(map[string]int))

	return a, circuits.calculate("a", map[string]int{"b": a})
}

type CircuitMap map[string][]string

func (circuits CircuitMap) calculate(wire string, allWires map[string]int) int {
	w, hasWire := allWires[wire]
	if hasWire {
		return w
	} else {
		logic, ok := circuits[wire]
		if ok {
			switch len(logic) {
			case 1:
				sig, nope := strconv.ParseInt(logic[0], 10, 32)
				if nope != nil {
					return circuits.calculate(logic[0], allWires)
				} else {
					return int(sig)
				}
			case 2:
				if "NOT" == logic[0] {
					allWires[wire] = 65535 - circuits.calculate(logic[1], allWires)
					// fmt.Printf("NOT logic @ %s = %d\n", wire, allWires[wire])
				}
			case 3:
				allWires[wire] = circuits.circuitOperation(logic, allWires)
			}
		} else {
			sig, _ := strconv.ParseInt(wire, 10, 32)
			// fmt.Printf("converting %s to %d", wire, sig)
			return int(sig)
		}
		result, _ := allWires[wire]
		return result
	}
}

func (circuits CircuitMap) circuitOperation(logic []string, allWires map[string]int) int {
	zero := circuits.calculate(logic[0], allWires)
	switch logic[1] {
	case "AND":
		bit, _ := strconv.ParseInt(logic[0], 10, 32)
		_, nope := strconv.ParseInt(logic[2], 10, 32)
		if nope == nil {
			return zero & int(bit)
		} else {
			return zero & circuits.calculate(logic[2], allWires)
		}
	case "OR":
		return zero | circuits.calculate(logic[2], allWires)
	case "LSHIFT":
		return zero << circuits.calculate(logic[2], allWires)
	case "RSHIFT":
		return zero >> circuits.calculate(logic[2], allWires)
	}
	return -1
}

func (d Day07) wireCircuits(input []string) CircuitMap {
	regex := regexp.MustCompile(`(.+)\s->\s([a-z]+)`)
	circuits := make(CircuitMap, len(input))
	for _, str := range input {
		m := regex.FindStringSubmatch(str)
		circuits[m[2]] = strings.Split(m[1], " ")
	}

	return circuits
}
