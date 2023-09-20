package aoc2019

import (
	"testing"

	"github.com/kevansimpson/util"
	"github.com/stretchr/testify/assert"
)

func TestDay01Solutions(t *testing.T) {
	input := util.ReadNumbers("data/input01.txt")
	moduleFuel, withAddedFuel := Day01{}.sumFuelRequirements(input)
	assert.Equal(t, 3266288, moduleFuel, "moduleFuel")
	assert.Equal(t, 4896582, withAddedFuel, "withAddedFuel")
}

func TestDay02Solutions(t *testing.T) {
	input := util.ReadCsvNumbers("data/input02.txt")
	moduleFuel, withAddedFuel := Day02{}.runProgram(input)
	assert.Equal(t, 4570637, moduleFuel, "moduleFuel")
	assert.Equal(t, 5485, withAddedFuel, "withAddedFuel")
}

func TestDay03Solutions(t *testing.T) {
	input := util.ReadLines("data/input03.txt")
	centralPort, fewestToIntersection := Day03{}.installCircuits(input)
	assert.Equal(t, 352, centralPort, "centralPort")
	assert.Equal(t, 43848, fewestToIntersection, "fewestToIntersection")
}

func TestDay04Solutions(t *testing.T) {
	pswdCount1, pswdCount2 := Day04{}.countValidPasswords(235741, 706948)
	assert.Equal(t, 1178, pswdCount1, "pswdCount1")
	assert.Equal(t, 763, pswdCount2, "pswdCount2")
}

func TestDay05Solutions(t *testing.T) {
	input := util.ReadCsvNumbers("data/input05.txt")
	diagCode1, diagCode5 := Day05{}.runProgram(input)
	assert.Equal(t, 13285749, diagCode1, "diagCode1")
	assert.Equal(t, 5000972, diagCode5, "diagCode5")
}
