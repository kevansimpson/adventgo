package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/8">Day 8</a>
 */

import (
	"fmt"
	"strings"

	"github.com/kevansimpson/util"
)

type Day08 struct{}

const WIDTH_08_2016 = 50
const HEIGHT_08_2016 = 6

func (d Day08) swipeMagneticCard(input []string) int {
	grid := make(util.Set[util.Point])
	for _, instruction := range input {
		nums := util.ExtractInts(instruction)
		if instruction[0:4] == "rect" {
			d.rect(grid, nums)

		} else if strings.Contains(instruction, "row") {
			d.row(grid, nums)

		} else if strings.Contains(instruction, "column") {
			d.column(grid, nums)
		}
	}

	d.display(grid)
	return len(grid)
}

func (d Day08) rect(grid util.Set[util.Point], nums []int) {
	for y := 0; y < nums[1]; y++ {
		for x := 0; x < nums[0]; x++ {
			util.Add(grid, util.Point{X: x, Y: y})
		}
	}
}

func (d Day08) row(grid util.Set[util.Point], nums []int) {
	row := make(util.Set[util.Point])
	for pt := range grid {
		if pt.Y == nums[0] {
			util.Add(row, pt)
		}
	}
	for pt := range row {
		delete(grid, pt)
	}
	for pt := range row {
		util.Add(grid, util.Point{X: ((nums[1] + pt.X) % WIDTH_08_2016), Y: pt.Y})
	}
}

func (d Day08) column(grid util.Set[util.Point], nums []int) {
	col := make(util.Set[util.Point])
	for pt := range grid {
		if pt.X == nums[0] {
			util.Add(col, pt)
		}
	}
	for pt := range col {
		delete(grid, pt)
	}
	for pt := range col {
		util.Add(grid, util.Point{X: pt.X, Y: ((nums[1] + pt.Y) % HEIGHT_08_2016)})
	}
}

func (d Day08) display(grid util.Set[util.Point]) {
	fmt.Println()
	for y := 0; y < HEIGHT_08_2016; y++ {
		fmt.Println()
		for x := 0; x < WIDTH_08_2016; x++ {
			_, hasPt := grid[util.Point{X: x, Y: y}]
			if hasPt {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println("\n--- day08,2016")
	fmt.Println()
}
