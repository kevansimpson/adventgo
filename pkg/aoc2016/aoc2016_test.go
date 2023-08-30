package aoc2016

import (
	"testing"

	"github.com/kevansimpson/util"
	"github.com/stretchr/testify/assert"
)

func TestDay01Solutions(t *testing.T) {
	input := util.ReadSingleLine("data/input01.txt")
	hqDist, alreadyVisited := Day01{}.visitEasterBunnyHQ(input)
	assert.Equal(t, 288, hqDist, "hqDist")
	assert.Equal(t, 111, alreadyVisited, "alreadyVisited")
}

func TestDay02Solutions(t *testing.T) {
	input := util.ReadLines("data/input02.txt")
	squareCode, diamondCode := Day02{}.enterBathroomCode(input)
	assert.Equal(t, "76792", squareCode, "squareCode")
	assert.Equal(t, "A7AC3", diamondCode, "diamondCode")
}

func TestDay03Solutions(t *testing.T) {
	input := util.ReadLines("data/input03.txt")
	pythagorean, byColumn := Day03{}.countValidTriangles(input)
	assert.Equal(t, 982, pythagorean, "pythagorean")
	assert.Equal(t, 1826, byColumn, "byColumn")
}

func TestDay04Solutions(t *testing.T) {
	input := util.ReadLines("data/input04.txt")
	sumRealSectors, northPoleSector := Day04{}.identifyRealRooms(input)
	assert.Equal(t, 137896, sumRealSectors, "sumRealSectors")
	assert.Equal(t, 501, northPoleSector, "northPoleSector")
}

func TestDay05Solutions(t *testing.T) {
	firstPswd, secondPswd := Day05{}.decodeDoorPassword("uqwqemis")
	assert.Equal(t, "1a3099aa", firstPswd, "firstPswd")
	assert.Equal(t, "694190cd", secondPswd, "secondPswd")
}
