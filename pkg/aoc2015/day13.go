package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/13">Day 13</a>
 */

import (
	"regexp"
	"slices"
	"strconv"

	"github.com/kevansimpson/util"
)

const HAPPY_ME_13_2015 = "ME"

func optimalHappinessWithMeOptional(input []string) (int, int) {
	happyMap := buildHappinessMap(input)
	people := gatherHappyPeople(happyMap)
	withoutMe := happinessBetweenAllPeople(happyMap, people)
	withMe := happinessBetweenAllPeople(happyMap, append(people, HAPPY_ME_13_2015))

	return withoutMe, withMe
}

// Calculates optimal happiness of all seating arrangements
func happinessBetweenAllPeople(happyMap map[TwoPeople]int, people []string) int {
	optimal := -1
	for _, arrangement := range util.Permutations(people) {
		happiness := happinessBetweenSomePeople(happyMap, arrangement)
		if happiness > optimal {
			optimal = happiness
		}
	}

	return optimal
}

// Calculates the happiness of a single seating arrangement
func happinessBetweenSomePeople(happyMap map[TwoPeople]int, arrangement []string) int {
	last := len(arrangement) - 1
	dist := happinessBetweenTwoPeople(happyMap, arrangement[last], arrangement[0])
	dist += happinessBetweenTwoPeople(happyMap, arrangement[0], arrangement[last])
	for ix := 0; ix < last; ix++ {
		p1, p2 := arrangement[ix], arrangement[ix+1]
		dist += happinessBetweenTwoPeople(happyMap, p1, p2)
		dist += happinessBetweenTwoPeople(happyMap, p2, p1)
	}

	return dist
}

func happinessBetweenTwoPeople(happyMap map[TwoPeople]int, one string, two string) int {
	if one == HAPPY_ME_13_2015 || two == HAPPY_ME_13_2015 {
		return 0
	}

	return happyMap[TwoPeople{one, two}]
}

type TwoPeople struct {
	one string
	two string
}

func buildHappinessMap(input []string) map[TwoPeople]int {
	regex := regexp.MustCompile(`(.+) would (.+) (\d+).*to (.+)\.`)
	happinessMap := make(map[TwoPeople]int)
	for _, str := range input {
		m := regex.FindStringSubmatch(str)
		p1, p2, gainOrLose := m[1], m[4], m[2]
		happiness, _ := strconv.Atoi(m[3])
		key := TwoPeople{p1, p2}
		if "gain" == gainOrLose {
			happinessMap[key] = happiness
		} else {
			happinessMap[key] = happiness * -1
		}
	}

	return happinessMap
}

func gatherHappyPeople(happyMap map[TwoPeople]int) []string {
	var people []string
	for key := range happyMap {
		if !slices.Contains(people, key.one) {
			people = append(people, key.one)
		}
		if !slices.Contains(people, key.two) {
			people = append(people, key.two)
		}
	}

	return people
}
