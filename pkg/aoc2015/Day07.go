package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/7">Day 7</a>
 */

import (
	"regexp"
	"strconv"
	"strings"
)

func assembleCircuits(input []string) (int, int) {
	circuits := wireCircuits(input)
	a := calculate("a", make(map[string]int), circuits)

	return a, calculate("a", map[string]int{"b": a}, circuits)
}

type CircuitMap map[string][]string

func calculate(wire string, allWires map[string]int, circuits CircuitMap) int {
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
					return calculate(logic[0], allWires, circuits)
				} else {
					return int(sig)
				}
			case 2:
				if "NOT" == logic[0] {
					allWires[wire] = 65535 - calculate(logic[1], allWires, circuits)
					// fmt.Printf("NOT logic @ %s = %d\n", wire, allWires[wire])
				}
			case 3:
				allWires[wire] = circuitOperation(logic, allWires, circuits)
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

func circuitOperation(logic []string, allWires map[string]int, circuits CircuitMap) int {
	zero := calculate(logic[0], allWires, circuits)
	switch logic[1] {
	case "AND":
		bit, _ := strconv.ParseInt(logic[0], 10, 32)
		_, nope := strconv.ParseInt(logic[2], 10, 32)
		if nope == nil {
			return zero & int(bit)
		} else {
			return zero & calculate(logic[2], allWires, circuits)
		}
	case "OR":
		return zero | calculate(logic[2], allWires, circuits)
	case "LSHIFT":
		return zero << calculate(logic[2], allWires, circuits)
	case "RSHIFT":
		return zero >> calculate(logic[2], allWires, circuits)
	}
	return -1
}

func wireCircuits(input []string) CircuitMap {
	regex := regexp.MustCompile(`(.+)\s->\s([a-z]+)`)
	circuits := make(CircuitMap, len(input))
	for _, str := range input {
		m := regex.FindStringSubmatch(str)
		circuits[m[2]] = strings.Split(m[1], " ")
	}

	return circuits
}
