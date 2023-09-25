package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/7">Day 7</a>
 */

import (
	"sync"

	"github.com/kevansimpson/util"
)

type Day07 struct{}

func (d Day07) chainAmplifiers(input []int) (int, int) {
	var wg sync.WaitGroup
	channel1, channel2 := make(chan int), make(chan int)
	wg.Add(2)

	go d.maxThrusterSignal(input, false, []int{0, 1, 2, 3, 4}, channel1, &wg)
	go d.maxThrusterSignal(input, true, []int{5, 6, 7, 8, 9}, channel2, &wg)
	sansFeedback, withFeedback := <-channel1, <-channel2
	close(channel1)
	close(channel2)
	wg.Wait()

	return sansFeedback, withFeedback
}

func (d Day07) maxThrusterSignal(codes []int, feedback bool, boosts []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	var calcWG sync.WaitGroup
	calcWG.Add(120) // 5!
	shared := make(chan int, 120)
	for _, perm := range util.Permutations(boosts) {
		go d.calcThrust(codes, feedback, perm, shared, &calcWG)
	}

	maxThrust := 0
	for i := 0; i < 120; i++ {
		t := <-shared
		if t > maxThrust {
			maxThrust = t
		}
	}

	close(shared)
	calcWG.Wait()
	ch <- maxThrust
}

func (d Day07) calcThrust(codes []int, feedback bool, boosts []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	amps := d.stackAmps(codes, feedback, boosts)
	var abcdeWG sync.WaitGroup
	abcdeWG.Add(5)
	for _, p := range amps {
		go p.runProgramAsync(&abcdeWG)
	}
	amps["A"].Input <- 0
	abcdeWG.Wait()

	ch <- <-amps["E"].Output
}

func (d Day07) stackAmps(codes []int, feedback bool, boosts []int) map[string]Program {
	amps := make(map[string]Program, 5)
	ab, bc, cd, de, outE := make(chan int, 2), make(chan int, 2), make(chan int, 2), make(chan int, 2), make(chan int, 2)
	ab <- boosts[1]
	bc <- boosts[2]
	cd <- boosts[3]
	de <- boosts[4]
	var inA chan int
	if feedback {
		inA = outE
	} else {
		inA = make(chan int, 2)
	}

	inA <- boosts[0]
	amps["A"] = CreateProgram(codes, inA, ab)
	amps["B"] = CreateProgram(codes, ab, bc)
	amps["C"] = CreateProgram(codes, bc, cd)
	amps["D"] = CreateProgram(codes, cd, de)
	amps["E"] = CreateProgram(codes, de, outE)
	return amps
}
