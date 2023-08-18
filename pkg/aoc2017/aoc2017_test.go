package aoc2017

import (
	"testing"

	"github.com/kevansimpson/util"
	"github.com/stretchr/testify/assert"
)

func TestDay01Solutions(t *testing.T) {
	input := util.ReadSingleLine("input01.txt")
	assert.Equal(t, 1251, solve1(input), "solve1")
	assert.Equal(t, 1244, solve2(input), "solve2")
}
