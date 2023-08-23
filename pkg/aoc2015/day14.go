package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/14">Day 14</a>
 */

import (
	"regexp"
	"strconv"
)

func hostReindeerOlympics(input []string) (int, int) {
	reindeerMap := buildReindeerMap(input)
	pointMap := make(map[string]int)
	winningPoints, maxDistance := 0, 0

	for sec := 1; sec < 2504; sec++ {
		// goldMedal, winners := identifyGoldMedalReindeers(reindeerMap, sec)
		dist, winners := identifyGoldMedalReindeers(reindeerMap, sec)
		for _, w := range winners {
			pts, ok := pointMap[w]
			if ok {
				pointMap[w] = pts + 1
			} else {
				pointMap[w] = 1
			}
		}
		if dist > maxDistance {
			maxDistance = dist
		}
	}

	for _, points := range pointMap {
		if points > winningPoints {
			winningPoints = points
		}
	}
	return maxDistance, winningPoints
}

func identifyGoldMedalReindeers(reindeerMap map[string]Reindeer, seconds int) (int, []string) {
	snapshot := make(map[int][]string)

	for name, r := range reindeerMap {
		dist := r.distanceTraveled(seconds)
		prev, ok := snapshot[dist]
		if ok {
			snapshot[dist] = append(prev, name)
		} else {
			snapshot[dist] = []string{name}
		}
	}

	var winners []string
	goldMedal := 0
	for pts, w := range snapshot {
		if pts > goldMedal {
			goldMedal = pts
			winners = w
		}
	}

	return goldMedal, winners
}

type Reindeer struct {
	kmPerSec, goTime, restTime int
}

func (r Reindeer) distanceTraveled(seconds int) int {
	total := r.totalTime()
	a, b := (seconds/total)*r.goTime*r.kmPerSec, seconds%total
	if b < r.goTime {
		return a + b*r.kmPerSec
	} else {
		return a + r.goTime*r.kmPerSec
	}
}

func (r Reindeer) totalTime() int {
	return r.goTime + r.restTime
}

func buildReindeerMap(input []string) map[string]Reindeer {
	regex := regexp.MustCompile(`(.+) can.+ (\d+) km.+ (\d+) seconds.* (\d+) .+`)
	reindeerMap := make(map[string]Reindeer)
	for _, str := range input {
		m := regex.FindStringSubmatch(str)
		kmPerSec, _ := strconv.Atoi(m[2])
		goTime, _ := strconv.Atoi(m[3])
		restTime, _ := strconv.Atoi(m[4])
		reindeerMap[m[1]] = Reindeer{kmPerSec, goTime, restTime}
	}

	return reindeerMap
}
