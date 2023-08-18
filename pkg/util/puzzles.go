package util

import (
	"bufio"
	"os"
)

type PuzzleSolver interface {
	solve1(input any) any
	solve2(input any) any
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
