package util

import (
	"bufio"
	"os"
)

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
