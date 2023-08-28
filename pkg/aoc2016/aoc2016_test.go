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
