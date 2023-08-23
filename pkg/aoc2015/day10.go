package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/10">Day 10</a>
 * Credit goes to
 * <a href="https://www.reddit.com/r/adventofcode/comments/3w6h3m/day_10_solutions/cxtuu0n/">
 * 	segfaultvicta
 * </a>
 * because concatenation solution using either bytes.Buffer or strings.Builder
 * was brutally slow no matter
 * <a href="https://golang-examples.tumblr.com/post/86169510884/fastest-string-contatenation">
 *	the benchmarks
 * </a>
 */

const LOOK_AND_SAY_10_2015 = "1321131112"

type Day10 struct{}

func (d Day10) applyLookAndSay(input string, firstRun int, secondRun int) (int, int) {
	said := []byte(input)
	for ix := 0; ix < firstRun; ix += 1 {
		said = d.lookSay(said)
	}

	firstLength := len(said)
	for ix := firstRun; ix < secondRun; ix++ {
		said = d.lookSay(said)
	}

	return firstLength, len(said)
}

func (d Day10) lookSay(said []byte) []byte {
	next := []byte{}
	for jx := 0; jx < len(said); {
		runEndsAt := d.indexToEndRun(said, jx)
		runSize := byte((runEndsAt + 1) - jx)
		next = append(next, runSize+byte(48)) //shens ??
		next = append(next, said[jx])
		jx = runEndsAt + 1
	}

	return next
}

func (d Day10) indexToEndRun(str []byte, from int) int {
	for ix := from + 1; ix < len(str); ix++ {
		if str[ix] != str[from] {
			return ix - 1
		}
	}
	return len(str) - 1
}
