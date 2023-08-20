package aoc2015

import (
	"testing"

	"github.com/kevansimpson/util"
	"github.com/stretchr/testify/assert"
)

func TestDay01Solutions(t *testing.T) {
	input := util.ReadSingleLine("data/input01.txt")
	assert.Equal(t, 74, whatFloorSanta(input), "whatFloorSanta")
	assert.Equal(t, 1795, santaEntersBasement(input), "santaEntersBasement")
}

func TestDay02Solutions(t *testing.T) {
	input := util.ReadLines("data/input02.txt")
	paper, ribbon := howMuchPaperAndRibbon(input)
	assert.Equal(t, 1588178, paper, "howMuchWrappingPaper")
	assert.Equal(t, 3783758, ribbon, "howMuchRibbon")
}

func TestDay03Solutions(t *testing.T) {
	input := util.ReadSingleLine("data/input03.txt")
	assert.Equal(t, 2081, santaRoute(input), "santaRoute")
	assert.Equal(t, 2341, roboSantaRoute(input), "roboSantaRoute")
}

func TestDay04Solutions(t *testing.T) {
	assert.Equal(t, 254575, fiveZeroHash(SecretKey_04_2015), "fiveZeroHash")
	assert.Equal(t, 1038736, sixZeroHash(SecretKey_04_2015), "sixZeroHash")
}

func TestDay05Solutions(t *testing.T) {
	input := util.ReadLines("data/input05.txt")
	assert.Equal(t, 258, oldNiceStrings(input), "oldNiceStrings")
	assert.Equal(t, 53, newNiceStrings(input), "newNiceStrings")
}

func TestDay06Solutions(t *testing.T) {
	input := util.ReadLines("data/input06.txt")
	lightsLit, totalBrightness := followLightCommands(input)
	assert.Equal(t, 543903, lightsLit, "lightsLit")
	assert.Equal(t, 14687245, totalBrightness, "totalBrightness")
}

func TestDay07Solutions(t *testing.T) {
	input := util.ReadLines("data/input07.txt")
	signalWireA, overrideSignal := assembleCircuits(input)
	assert.Equal(t, 46065, signalWireA, "signalWireA")
	assert.Equal(t, 14134, overrideSignal, "overrideSignal")
}

func TestDay08Solutions(t *testing.T) {
	input := util.ReadLines("data/input08.txt")
	oldEncoding, newEncoding := encodeSantasList(input)
	assert.Equal(t, 1333, oldEncoding, "oldEncoding")
	assert.Equal(t, 2046, newEncoding, "newEncoding")
}

func TestDay09Solutions(t *testing.T) {
	input := util.ReadLines("data/input09.txt")
	shortestRoute, longestRoute := calculateSantaRoutes(input)
	assert.Equal(t, 207, shortestRoute, "shortestRoute")
	assert.Equal(t, 804, longestRoute, "longestRoute")
}

func TestDay10Solutions(t *testing.T) {
	fortyTimes, fiftyTimes := applyLookAndSay(LOOK_AND_SAY_10_2015, 40, 50)
	assert.Equal(t, 492982, fortyTimes, "fortyTimes")
	assert.Equal(t, 6989950, fiftyTimes, "fiftyTimes")
}

func TestDay11Solutions(t *testing.T) {
	firstPswd, secondPswd := nextTwoPasswords(SANTA_CURRENT_PASSWORD_11_2015)
	assert.Equal(t, "vzbxxyzz", firstPswd, "firstPswd")
	assert.Equal(t, "vzcaabcc", secondPswd, "secondPswd")
}

func TestDay12Solutions(t *testing.T) {
	input := util.ReadSingleLine("data/input12.txt")
	assert.Equal(t, 111754, sumDocumentNumbers(input), "sumDocumentNumbers")
	assert.Equal(t, 65402, sumWithoutDoubleCountingRed(input), "sumWithoutDoubleCountingRed")
}

func TestDay13Solutions(t *testing.T) {
	input := util.ReadLines("data/input13.txt")
	happyWithoutMe, lessHappyWithMe := optimalHappinessWithMeOptional(input)
	assert.Equal(t, 733, happyWithoutMe, "happyWithoutMe")
	assert.Equal(t, 725, lessHappyWithMe, "lessHappyWithMe") // lolz
}
