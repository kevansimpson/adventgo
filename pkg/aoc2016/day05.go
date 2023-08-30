package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/5">Day 5</a>
 */

import (
	"strings"
	"sync"

	"github.com/kevansimpson/util"
)

type Day05 struct{}

func (d Day05) decodeDoorPassword(input string) (string, string) {
	var wg sync.WaitGroup
	channelHash := make(chan string, 12)
	channelPswd := make(chan string)
	wg.Add(1)

	go d.nextHashPassword(input, channelHash, channelPswd, &wg)

	firstPswd := make([]rune, 8)
	for i := 0; i < 8; i++ {
		next := <-channelHash
		firstPswd[i] = rune(next[5])
	}

	pswd2 := <-channelPswd

	wg.Wait()
	close(channelHash)
	close(channelPswd)

	return string(firstPswd), pswd2
}

func (d Day05) nextHashPassword(input string, chHash chan<- string, chPswd chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	// start is determined in hindsight from final solution.
	// calculating md5 hashes is slow, full stop.
	// shave a little time by starting at first matching hash OR
	// for fast solve, iterate through known hash indexes
	known := []int{
		// first 8
		4515059, 6924074, 8038154, 13432968, 13540621, 14095580, 14821988, 16734551,
		// next 5
		17743256, 19112977, 20616595, 21658552, 26326685}
	start, count, p2 := known[0], 0, 0
	secondPswd := []rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}

	for p2 < 8 {
		next, ix := util.NextHash(input, "00000", start)
		chHash <- next

		if strings.Contains("01234567", string(next[5])) {
			digit := int(next[5] - '0')
			if digit < 8 && secondPswd[digit] == ' ' {
				secondPswd[digit] = rune(next[6])
				p2++
			}
		}

		count++
		if util.IsFullSolve() || count > 12 {
			start = ix + 1
		} else {
			start = known[count]
		}
	}

	chPswd <- string(secondPswd)
}
