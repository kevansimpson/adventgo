package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/02">Day 02</a>
 */

import (
	"sync"

	"github.com/kevansimpson/util"
)

type Day02 struct{}

func (d Day02) enterBathroomCode(input []string) (string, string) {
	var wg sync.WaitGroup
	wg.Add(2)
	channel1 := make(chan string)
	channel2 := make(chan string)

	go d.readFrontDeskDoc(input, squarePad, channel1, &wg)
	go d.readFrontDeskDoc(input, diamondPad, channel2, &wg)
	code1, code2 := <-channel1, <-channel2

	close(channel1)
	close(channel2)
	wg.Wait()
	return code1, code2
}

func (d Day02) readFrontDeskDoc(input []string, pad numberPad, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	code := make([]rune, 5)
	five := pad.starting5()

	for i, instruction := range input {
		start := five
		for _, dir := range instruction {
			next := start.RuneStep(dir)
			_, has := pad[next]
			if has {
				start = next
			}
		}

		code[i] = pad[start]
	}

	ch <- string(code)
}

func (pad numberPad) starting5() util.Point {
	for pt, ch := range pad {
		if ch == '5' {
			return pt
		}
	}
	panic("definitely not 5x5")
}

type numberPad map[util.Point]rune

// 1 2 3
// 4 5 6
// 7 8 9
var squarePad = numberPad{
	{X: -1, Y: 1}: '1', {X: 0, Y: 1}: '2', {X: 1, Y: 1}: '3',
	{X: -1, Y: 0}: '4', util.ORIGIN: '5', {X: 1, Y: 0}: '6',
	{X: -1, Y: -1}: '7', {X: 0, Y: -1}: '8', {X: 1, Y: -1}: '9'}

/**
 *	   1
 *   2 3 4
 * 5 6 7 8 9
 *   A B C
 *     D
 */
var diamondPad = numberPad{
	{X: 0, Y: 2}:  '1',
	{X: -1, Y: 1}: '2', {X: 0, Y: 1}: '3', {X: 1, Y: 1}: '4',
	{X: -2, Y: 0}: '5', {X: -1, Y: 0}: '6', util.ORIGIN: '7', {X: 1, Y: 0}: '8', {X: 2, Y: 0}: '9',
	{X: -1, Y: -1}: 'A', {X: 0, Y: -1}: 'B', {X: 1, Y: -1}: 'C',
	{X: 0, Y: -2}: 'D'}
