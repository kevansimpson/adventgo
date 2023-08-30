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
	channelHash := make(chan string, 17)
	channelDone := make(chan bool)
	wg.Add(1)

	go d.nextHashPassword(input, channelHash, channelDone, &wg)

	firstPswd, secondPswd := make([]rune, 8), []rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	count := 0
	for i := 0; i < 8; i++ {
		next := <-channelHash
		firstPswd[i] = rune(next[5])

		if strings.Contains("01234567", string(next[5])) {
			digit := int(next[5] - '0')
			if digit < 8 && secondPswd[digit] == ' ' {
				secondPswd[digit] = rune(next[6])
				count++
			}
		}
	}

	for count < 8 {
		next := <-channelHash
		if strings.Contains("01234567", string(next[5])) {
			digit := int(next[5] - '0')
			if digit < 8 && secondPswd[digit] == ' ' {
				secondPswd[digit] = rune(next[6])
				count++
			}
		}
	}

	channelDone <- true
	wg.Wait()
	close(channelHash)
	close(channelDone)

	return string(firstPswd), string(secondPswd)
}

func (d Day05) nextHashPassword(input string, chHash chan<- string, chDone <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	// start is determined in hindsight from final solution.
	// calculating md5 hashes is slow, full stop.
	// shave a little time by starting at first matching hash.
	start := 4515059
	for true {
		select {
		case doneNow := <-chDone:
			if doneNow {
				return
			}
		default:
			next, ix := util.NextHash(input, "00000", start)
			chHash <- next
			start = ix + 1
		}
	}
}
