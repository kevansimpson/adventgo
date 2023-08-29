package aoc2016

import (
	"sort"

	"github.com/kevansimpson/util"
)

/**
 * <a href="https://adventofcode.com/2016/day/3">Day 3</a>
 */

type Day03 struct{}

func (d Day03) countValidTriangles(input []string) (int, int) {
	triangles := d.readTriangles(input)
	// count by column first b/c pythagorean sorts triangle arrays
	byColumn := d.countTrianglesByColumn(triangles)
	pythagorean := 0
	for _, tri := range triangles {
		if d.isValidTriangle(tri...) {
			pythagorean += 1
		}
	}
	return pythagorean, byColumn
}

func (d Day03) isValidTriangle(triangle ...int) bool {
	sort.Ints(triangle)
	return (triangle[0] + triangle[1]) > triangle[2]
}

func (d Day03) countTrianglesByColumn(triangles [][]int) int {
	count := 0
	for i := 0; i < len(triangles); i += 3 {
		row1, row2, row3 := triangles[i], triangles[i+1], triangles[i+2]
		for col := 0; col < 3; col++ {
			if d.isValidTriangle(row1[col], row2[col], row3[col]) {
				count += 1
			}
		}
	}

	return count
}

func (d Day03) readTriangles(input []string) [][]int {
	triangles := make([][]int, len(input))
	for i, row := range input {
		triangles[i] = util.ExtractInts(row)
	}

	return triangles
}
