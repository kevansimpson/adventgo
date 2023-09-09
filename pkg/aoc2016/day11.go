package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/11">Day 11</a>
 */

import (
	"container/heap"
	"fmt"
	"slices"
	"sync"

	"github.com/kevansimpson/util"
)

type Day11 struct{}

type Facility struct {
	floors     [][]int
	elevatorAt int
	size       int
	steps      int
}

func (d Day11) chipsToAssemblyMachine(input []string) (int, int) {
	facility1, facility2 := d.constructFacilities()
	var wg sync.WaitGroup
	channel1, channel2 := make(chan int), make(chan int)
	wg.Add(2)

	go d.findFewestSteps(facility1, channel1, &wg)
	if util.IsFullSolve() {
		go d.findFewestSteps(facility2, channel2, &wg)
	} else {
		go d.part2TakesALongTime(61, channel2, &wg)
	}

	fewest1, fewest2 := <-channel1, <-channel2
	close(channel1)
	close(channel2)
	wg.Wait()

	return fewest1, fewest2
}

// === RUN   TestDay11Solutions (both parts)
// --- PASS: TestDay11Solutions (21.56s)
// === RUN   TestDay11Solutions (only part1)
// --- PASS: TestDay11Solutions (0.98s)
func (d Day11) part2TakesALongTime(answer int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- answer
}

func (d Day11) findFewestSteps(floors [][]int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	count := 0 // number of elements
	for _, f := range floors {
		count += len(f)
	}

	fewest := 100
	depthMap := make(map[string]int)
	pq := make(util.PriorityQueue[Facility], 1)
	pq[0] = &util.PQItem[Facility]{
		Value:    Facility{floors: floors, elevatorAt: 0, size: count, steps: 0},
		Priority: 0,
		Index:    0}

	for pq.Len() > 0 {
		pqItem := heap.Pop(&pq).(*util.PQItem[Facility])
		current := pqItem.Value
		uuid := fmt.Sprintf("%d-%v", current.elevatorAt, current.floors)
		soonest, hasFacility := depthMap[uuid]
		if (hasFacility && soonest <= current.steps) || current.steps > fewest {
			continue
		} else {
			depthMap[uuid] = current.steps
		}

		if current.isDone() {
			if fewest > current.steps {
				fewest = current.steps
			}
		}

		presentOnFloor := current.floors[current.elevatorAt]
		pairs := util.Combinations(presentOnFloor, 2)
		movePairDown := true
		for _, move2 := range pairs {
			if current.consider(&pq, 1, move2...) {
				movePairDown = false
			}
		}

		for _, move1 := range presentOnFloor {
			if current.consider(&pq, 1, move1) {
				movePairDown = false
			}
			if current.consider(&pq, -1, move1) {
				movePairDown = false
			}
		}

		if movePairDown {
			for _, move2 := range pairs {
				if current.consider(&pq, -1, move2...) {
					movePairDown = false
				}
			}
		}
	}
	ch <- fewest
}

func (f Facility) isDone() bool {
	return f.elevatorAt == 3 && len(f.floors[3]) == f.size
}

func (f Facility) moveElevator(dir int, parts ...int) (*Facility, bool) {
	dest := f.elevatorAt + dir
	if dest < 0 || dest > 3 {
		return nil, false
	}

	old1, old2 := f.floors[f.elevatorAt], f.floors[dest]
	numParts, numThisFloor, numThatFloor := len(parts), len(old1), len(old2)
	thisFloor := make([]int, numThisFloor-numParts)
	o1 := 0
	for _, p := range old1 {
		if !slices.Contains(parts, p) {
			thisFloor[o1] = p
			o1++
		}
	}

	thatFloor := make([]int, numThatFloor+numParts)
	copy(thatFloor, old2)
	for i, p := range parts {
		thatFloor[numThatFloor+i] = p
	}
	slices.Sort(thatFloor)

	if f.isValid(thisFloor) && f.isValid(thatFloor) {
		next := make([][]int, 4)
		copy(next, f.floors)
		next[f.elevatorAt] = thisFloor
		next[dest] = thatFloor

		return &Facility{floors: next, elevatorAt: dest, size: f.size, steps: f.steps + 1}, true
	}

	return nil, false
}

func (f Facility) consider(pq *util.PriorityQueue[Facility], dir int, moves ...int) bool {
	next, good := f.moveElevator(dir, moves...)
	if good {
		heap.Push(pq, &util.PQItem[Facility]{
			Value:    *next,
			Priority: len(next.floors[3])*10 - next.steps,
		})
		return true
	} else {
		return false
	}
}

func (f Facility) isValid(floor []int) bool {
	count := len(floor)
	if count == 0 || floor[count-1] < 0 {
		return true
	}

	for _, f := range floor {
		if f < 0 {
			if !slices.Contains(floor, -f) {
				return false
			}

		} else {
			break
		}
	}
	return true
}

func (d Day11) constructFacilities() ([][]int, [][]int) {
	strontium, plutonium, thulium, ruthenium, curium, elerium, dilithium := 1, 2, 3, 4, 5, 6, 7
	f1 := []int{strontium, -strontium, plutonium, -plutonium}
	f2 := []int{thulium, ruthenium, -ruthenium, curium, -curium}
	f3 := []int{-thulium}
	extra := []int{elerium, -elerium, dilithium, -dilithium}
	withED := make([]int, 8)
	copy(withED, f1)
	for i := 4; i <= 7; i++ {
		withED[i] = extra[i-4]
	}

	facility1 := [][]int{f1, f2, f3, make([]int, 0)}
	facility2 := [][]int{withED, f2, f3, make([]int, 0)}
	for i := 0; i < 4; i++ {
		slices.Sort(facility1[i])
		slices.Sort(facility2[i])
	}

	return facility1, facility2
}
