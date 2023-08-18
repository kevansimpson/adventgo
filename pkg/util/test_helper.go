package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertSolutions(t *testing.T, s PuzzleSolver, input any, expected1 any, expected2 any) {
	assert.Equal(t, expected1, s.solve1(input), "solve1")
	assert.Equal(t, expected2, s.solve2(input), "solve1")
}
