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
