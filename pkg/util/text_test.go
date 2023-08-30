package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumns(t *testing.T) {
	rows := []string{"adg", "beh", "cfi"}
	cols := Columns(rows)
	assert.Equal(t, "abc", cols[0], "cols[0]")
	assert.Equal(t, "def", cols[1], "cols[0]")
	assert.Equal(t, "ghi", cols[2], "cols[0]")
}

func TestShiftText(t *testing.T) {
	assert.Equal(t, "def", ShiftText("abc", 3), "abc->3")
	assert.Equal(t, "abc", ShiftText("xyz", 3), "xyz->3")
}
