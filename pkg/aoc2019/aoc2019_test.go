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

func TestDay06Solutions(t *testing.T) {
	input := util.ReadLines("data/input06.txt")
	totalOrbits, minimumTransfers := Day06{}.mapOrbitalTransfers(input)
	assert.Equal(t, 249308, totalOrbits, "totalOrbits")
	assert.Equal(t, 349, minimumTransfers, "minimumTransfers")
}

func TestDay07Solutions(t *testing.T) {
	input := util.ReadCsvNumbers("data/input07.txt")
	sansFeedback, withFeedback := Day07{}.chainAmplifiers(input)
	assert.Equal(t, 75228, sansFeedback, "sansFeedback")
	assert.Equal(t, 79846026, withFeedback, "withFeedback")
}

func TestDay08Solutions(t *testing.T) {
	input := util.ReadSingleLine("data/input08.txt")
	zeroLayer, drawImage := Day08{}.layerImages(input)
	assert.Equal(t, 1935, zeroLayer, "zeroLayer")
	expectedImage := "011001111010000100101000010010100001000010010100001000011100100001001010000100001000010000100101000010010100001000010010100000110010000111100110011110"
	assert.Equal(t, expectedImage, drawImage, "drawImage")
}

func TestDay09Solutions(t *testing.T) {
	input := util.ReadCsvNumbers("data/input09.txt")
	code1, code2 := Day09{}.runProgram(input)
	assert.Equal(t, 2453265701, code1, "code1")
	assert.Equal(t, 80805, code2, "code2")
}
