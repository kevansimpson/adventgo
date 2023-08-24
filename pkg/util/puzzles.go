package util

import (
	"bufio"
	"flag"
	"os"
	"strconv"
)

var fullSolve = flag.String("full", "false", "bool flag to run full solutions")

func IsFullSolve() bool {
	fs, _ := strconv.ParseBool(*fullSolve)
	return fs
}

func FastOrFull[T any](fast T, full func() T) T {
	if IsFullSolve() {
		return full()
	} else {
		return fast
	}
}

func FullOrFast[T any](full T, fast T) T {
	if IsFullSolve() {
		return full
	} else {
		return fast
	}
}

func ReadSingleLine(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func ReadLines(filename string) []string {
	data, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	data.Close()

	return lines
}

func ReadNumbers(filename string) []int {
	lines := ReadLines(filename)
	var nums []int
	for _, str := range lines {
		n, _ := strconv.Atoi(str)
		nums = append(nums, n)
	}

	return nums
}
