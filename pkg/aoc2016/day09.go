package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/9">Day 9</a>
 */

import (
	"strconv"
	"strings"
	"sync"

	"github.com/kevansimpson/util"
)

type Day09 struct{}

func (d Day09) decompressFile(input string) (int, int) {
	var wg sync.WaitGroup
	wg.Add(2)
	channel1 := make(chan int)
	channel2 := make(chan int)

	go d.decompressV1(input, channel1, &wg)
	go d.decompressV2(input, channel2, &wg)
	v1, v2 := <-channel1, <-channel2

	close(channel1)
	close(channel2)
	wg.Wait()
	return v1, v2
}

func (d Day09) decompressV1(input string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	length, start, index := 0, 0, strings.Index(input, "(")
	for index >= 0 {
		length += index - start
		close := util.IndexAfter(input, ")", index)
		marker := strings.Split(input[index+1:close], "x")
		m0, _ := strconv.Atoi(marker[0])
		m1, _ := strconv.Atoi(marker[1])
		length += m0 * m1
		start = close + m0 + 1
		index = util.IndexAfter(input, "(", start)
	}

	ch <- length + len(input[start:])
}

func (d Day09) decompressV2(input string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- d.runDecompressionV2(input)
}

func (d Day09) runDecompressionV2(input string) int {
	length, start, index := 0, 0, strings.Index(input, "(")
	for index >= 0 {
		length += index - start
		close := util.IndexAfter(input, ")", index)
		marker := strings.Split(input[index+1:close], "x")
		m0, _ := strconv.Atoi(marker[0])
		m1, _ := strconv.Atoi(marker[1])
		start = close + 1
		repeated := input[start : start+m0]
		if strings.Contains(repeated, "(") {
			length += m1 * d.runDecompressionV2(repeated)
		} else {
			length += m0 * m1
		}
		start += m0
		index = util.IndexAfter(input, "(", start)
	}

	return length + len(input[start:])
}
