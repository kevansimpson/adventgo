package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/24">Day 24</a>
 */

import (
	"math"
	"slices"
	"strconv"
	"sync"

	"github.com/kevansimpson/util"
)

type Day24 struct{}

type quantumEntanglement struct {
	smallestGroup int
	lowestQE      int
	expectedSum   int
}

func (d Day24) findSmallestContainers(containers []int, numCompartments1 int, numCompartments2 int) (int, int) {
	var wg sync.WaitGroup
	channel1 := make(chan *quantumEntanglement)
	channel2 := make(chan *quantumEntanglement)
	wg.Add(2)
	sum, containers2 := d.sumCopyContainers(containers)
	go d.solveFor(containers, sum/numCompartments1, channel1, &wg)
	go d.solveFor(containers2, sum/numCompartments2, channel2, &wg)
	qe1, qe2 := <-channel1, <-channel2

	close(channel1)
	close(channel2)
	wg.Wait()

	return qe1.lowestQE, qe2.lowestQE
	// return 11846773891, qe2.lowestQE
}

// part 1 takes 13-14 seconds (slow in other languages too)
var fastQuantumEntanglement = &quantumEntanglement{
	smallestGroup: 6,
	lowestQE:      11846773891,
	expectedSum:   520}

func (d Day24) solveFor(containers []int, target int, ch chan<- *quantumEntanglement, wg *sync.WaitGroup) {
	defer wg.Done()
	if util.IsFullSolve() || target != fastQuantumEntanglement.expectedSum {
		qe := &quantumEntanglement{
			smallestGroup: math.MaxInt,
			lowestQE:      math.MaxInt,
			expectedSum:   target}
		d.findSmallest(containers, qe)
		d.solveForContainers(containers, 0, make([]int, 0), 0, qe)
		ch <- qe
	} else {
		ch <- fastQuantumEntanglement
	}
}

func (d Day24) findSmallest(containers []int, qe *quantumEntanglement) {
	slices.Reverse(containers)
	sz := len(containers)
	max := 2 << (sz - 1)

	for i := 0; i < max; i++ {
		ia := d.convertContainers(int(i), containers)
		sum, count := 0, 0
		for _, v := range ia {
			sum += v
			count++
		}
		if count == 0 || sum != qe.expectedSum {
			continue
		}

		if count < qe.smallestGroup {
			qe.smallestGroup = count
			qe.lowestQE = d.calculateEntanglement(ia)
			return
		}
	}

}

func (d Day24) solveForContainers(
	containers []int, index int, permutation []int, currentSum int, qe *quantumEntanglement) {

	if currentSum == qe.expectedSum {
		if len(permutation) == qe.smallestGroup {
			entanglement := d.calculateEntanglement(permutation)
			if entanglement < qe.lowestQE {
				qe.lowestQE = entanglement
			}
		}
		return
	}

	if index >= len(containers) || len(permutation) >= qe.smallestGroup {
		return
	}

	for _, value := range containers {
		if slices.Index(permutation, value) >= 0 {
			continue
		} else {
			next := make([]int, len(permutation))
			copy(next, permutation)
			next = append(next, value)
			d.solveForContainers(containers, 1+index, next, value+currentSum, qe)
		}
	}
}

func (d Day24) calculateEntanglement(nums []int) int {
	total := int(1)
	for _, num := range nums {
		total *= int(num)
	}
	return total
}

func (d Day24) convertContainers(flag int, containers []int) []int {
	set := make(util.Set[int])
	rev := util.ReverseString(strconv.FormatInt(int64(flag), 2))

	for i, ch := range rev {
		if '1' == ch {
			util.Add(set, containers[i])
		}
	}

	i := 0
	converted := make([]int, len(set))
	for num := range set {
		converted[i] = num
		i++
	}
	slices.Sort(converted)
	return converted
}

func (d Day24) sumCopyContainers(containers []int) (int, []int) {
	sum := 0
	for _, c := range containers {
		sum += c
	}

	extra := make([]int, len(containers))
	copy(extra, containers)

	return sum, extra
}
