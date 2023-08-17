package util

import (
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
