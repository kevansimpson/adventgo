package aoc2015

import (
	"testing"

	"github.com/kevansimpson/util"
	"github.com/stretchr/testify/assert"
)

func TestDay01Solutions(t *testing.T) {
	input := util.ReadSingleLine("input01.txt")
	assert.Equal(t, 74, whatFloorSanta(input), "whatFloorSanta")
	assert.Equal(t, 1795, santaEntersBasement(input), "santaEntersBasement")
}

func TestDay02Solutions(t *testing.T) {
	input := util.ReadLines("input02.txt")
	paper, ribbon := howMuchPaperAndRibbon(input)
	assert.Equal(t, 1588178, paper, "howMuchWrappingPaper")
	assert.Equal(t, 3783758, ribbon, "howMuchRibbon")
}

func TestDay03Solutions(t *testing.T) {
	input := util.ReadSingleLine("input03.txt")
	assert.Equal(t, 2081, santaRoute(input), "santaRoute")
	assert.Equal(t, 2341, roboSantaRoute(input), "roboSantaRoute")
}
