package aoc2016

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

/**
 * <a href="https://adventofcode.com/2016/day/10">Day 10</a>
 */

type Day10 struct{}
type BotInstruction struct {
	BotId, Low, High string
}

func (d Day10) inspectMicrochips(input []string) (int, int) {
	bots, instructions := d.chipFactory(input)
	botNumber := 0

	for d.hasComplexInstruction(bots) {
		for k, v := range bots {
			if len(v) == 2 {
				if slices.Contains(v, 17) && slices.Contains(v, 61) {
					target, _ := strconv.Atoi(k[3:])
					botNumber = target
				}

				ab := make([]int, len(v))
				copy(ab, v)
				slices.Sort(ab)
				bots[k] = make([]int, 0)
				instr := instructions[k]
				low, hasLow := bots[instr.Low]
				if hasLow {
					bots[instr.Low] = append(low, ab[0])
				} else {
					bots[instr.Low] = []int{ab[0]}
				}
				high, hasHigh := bots[instr.High]
				if hasHigh {
					bots[instr.High] = append(high, ab[1])
				} else {
					bots[instr.High] = []int{ab[1]}
				}
			}
		}
	}

	return botNumber, bots["output0"][0] * bots["output1"][0] * bots["output2"][0]
}

func (d Day10) hasComplexInstruction(bots map[string][]int) bool {
	for _, v := range bots {
		if len(v) > 1 {
			return true
		}
	}
	return false
}

func (d Day10) chipFactory(input []string) (map[string][]int, map[string]BotInstruction) {
	initialBots := make(map[string][]int)
	instructions := make(map[string]BotInstruction)

	for _, line := range input {
		bits := strings.Split(line, " ")
		if "value" == bits[0] {
			key := d.chipKey(bits, 4, 5)
			chip, _ := strconv.Atoi(bits[1])
			list, hasList := initialBots[key]
			if hasList {
				initialBots[key] = append(list, chip)
			} else {
				initialBots[key] = []int{chip}
			}
		} else {
			instr := BotInstruction{
				BotId: d.chipKey(bits, 0, 1),
				Low:   d.chipKey(bits, 5, 6),
				High:  d.chipKey(bits, 10, 11)}
			instructions[instr.BotId] = instr
		}
	}

	return initialBots, instructions
}

func (d Day10) chipKey(bits []string, ix1 int, ix2 int) string {
	return fmt.Sprintf("%s%s", bits[ix1], bits[ix2])
}
