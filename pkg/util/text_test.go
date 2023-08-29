package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftText(t *testing.T) {
	assert.Equal(t, "def", ShiftText("abc", 3), "abc->3")
	assert.Equal(t, "abc", ShiftText("xyz", 3), "xyz->3")
}
