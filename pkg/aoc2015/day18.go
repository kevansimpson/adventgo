package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/18">Day 18</a>
 */

type Day18 struct{}
type LightGrid struct {
	area [][]bool
}

func (d Day18) totalLights(input []string) (int, int) {
	grid := d.loadGrid(input)
	gridBrokenCorners := d.copyGrid(grid)
	gridBrokenCorners.breakCornerLights()

	for i := 0; i < 100; i++ {
		grid, gridBrokenCorners = d.clickBoth(grid, gridBrokenCorners)
	}

	return grid.countLights(), gridBrokenCorners.countLights()
}

func (d Day18) clickBoth(grid LightGrid, gridBrokenCorners LightGrid) (LightGrid, LightGrid) {
	next, nextBC := d.click(grid), d.click(gridBrokenCorners)
	nextBC.breakCornerLights()

	return next, nextBC
}

func (d Day18) click(grid LightGrid) LightGrid {
	next := d.copyGrid(grid)
	size := len(grid.area[0])
	for i := 1; i <= size-2; i++ {
		for j := 1; j <= size-2; j++ {
			neighbors := grid.countNeighbors(i, j)
			next.area[i][j] = neighbors == 3 || (neighbors == 2 && grid.area[i][j])
		}
	}
	return next
}

func (grid *LightGrid) breakCornerLights() {
	size := len(grid.area[0])
	grid.area[1][1] = true
	grid.area[1][size-2] = true
	grid.area[size-2][1] = true
	grid.area[size-2][size-2] = true
}

func (grid LightGrid) countLights() int {
	on := 0
	for _, row := range grid.area {
		for _, lightIsOn := range row {
			if lightIsOn {
				on += 1
			}
		}
	}

	return on
}

func (grid LightGrid) countNeighbors(i int, j int) int {
	on := 0
	if grid.area[i-1][j-1] {
		on += 1
	}
	if grid.area[i][j-1] {
		on += 1
	}
	if grid.area[i+1][j-1] {
		on += 1
	}
	if grid.area[i-1][j] {
		on += 1
	}
	if grid.area[i+1][j] {
		on += 1
	}
	if grid.area[i-1][j+1] {
		on += 1
	}
	if grid.area[i][j+1] {
		on += 1
	}
	if grid.area[i+1][j+1] {
		on += 1
	}
	return on
}

func (d Day18) copyGrid(grid LightGrid) LightGrid {
	size := len(grid.area)
	var area [][]bool
	for _, row := range grid.area {
		dupe := make([]bool, size)
		copy(dupe, row)
		area = append(area, dupe)
	}
	return LightGrid{area}
}

func (d Day18) loadGrid(input []string) LightGrid {
	size := len(input)
	var area [][]bool
	area = append(area, make([]bool, size+2)) // top row
	for _, str := range input {
		row := make([]bool, size+2)
		for j, bit := range str {
			row[j+1] = bit == '#'
		}
		area = append(area, row)
	}
	area = append(area, make([]bool, size+2)) // bottom row

	return LightGrid{area}
}
