package aoc2015

import (
	"testing"

	"github.com/kevansimpson/util"
	"github.com/stretchr/testify/assert"
)

func TestDay01Solutions(t *testing.T) {
	input := util.ReadSingleLine("data/input01.txt")
	assert.Equal(t, 74, Day01{}.whatFloorSanta(input), "whatFloorSanta")
	assert.Equal(t, 1795, Day01{}.santaEntersBasement(input), "santaEntersBasement")
}

func TestDay02Solutions(t *testing.T) {
	input := util.ReadLines("data/input02.txt")
	paper, ribbon := Day02{}.howMuchPaperAndRibbon(input)
	assert.Equal(t, 1588178, paper, "howMuchWrappingPaper")
	assert.Equal(t, 3783758, ribbon, "howMuchRibbon")
}

func TestDay03Solutions(t *testing.T) {
	input := util.ReadSingleLine("data/input03.txt")
	assert.Equal(t, 2081, Day03{}.santaRoute(input), "santaRoute")
	assert.Equal(t, 2341, Day03{}.roboSantaRoute(input), "roboSantaRoute")
}

func TestDay04Solutions(t *testing.T) {
	assert.Equal(t, 254575, Day04{}.fiveZeroHash(SecretKey_04_2015), "fiveZeroHash")
	assert.Equal(t, 1038736, Day04{}.sixZeroHash(SecretKey_04_2015), "sixZeroHash")
}

func TestDay05Solutions(t *testing.T) {
	input := util.ReadLines("data/input05.txt")
	assert.Equal(t, 258, Day05{}.oldNiceStrings(input), "oldNiceStrings")
	assert.Equal(t, 53, Day05{}.newNiceStrings(input), "newNiceStrings")
}

func TestDay06Solutions(t *testing.T) {
	input := util.ReadLines("data/input06.txt")
	lightsLit, totalBrightness := Day06{}.followLightCommands(input)
	assert.Equal(t, 543903, lightsLit, "lightsLit")
	assert.Equal(t, 14687245, totalBrightness, "totalBrightness")
}

func TestDay07Solutions(t *testing.T) {
	input := util.ReadLines("data/input07.txt")
	signalWireA, overrideSignal := Day07{}.assembleCircuits(input)
	assert.Equal(t, 46065, signalWireA, "signalWireA")
	assert.Equal(t, 14134, overrideSignal, "overrideSignal")
}

func TestDay08Solutions(t *testing.T) {
	input := util.ReadLines("data/input08.txt")
	oldEncoding, newEncoding := Day08{}.encodeSantasList(input)
	assert.Equal(t, 1333, oldEncoding, "oldEncoding")
	assert.Equal(t, 2046, newEncoding, "newEncoding")
}

func TestDay09Solutions(t *testing.T) {
	input := util.ReadLines("data/input09.txt")
	shortestRoute, longestRoute := Day09{}.calculateSantaRoutes(input)
	assert.Equal(t, 207, shortestRoute, "shortestRoute")
	assert.Equal(t, 804, longestRoute, "longestRoute")
}

func TestDay10Solutions(t *testing.T) {
	fortyTimes, fiftyTimes := Day10{}.applyLookAndSay(LOOK_AND_SAY_10_2015, 40, 50)
	assert.Equal(t, 492982, fortyTimes, "fortyTimes")
	assert.Equal(t, 6989950, fiftyTimes, "fiftyTimes")
}

func TestDay11Solutions(t *testing.T) {
	firstPswd, secondPswd := Day11{}.nextTwoPasswords(SANTA_CURRENT_PASSWORD_11_2015)
	assert.Equal(t, "vzbxxyzz", firstPswd, "firstPswd")
	assert.Equal(t, "vzcaabcc", secondPswd, "secondPswd")
}

func TestDay12Solutions(t *testing.T) {
	input := util.ReadSingleLine("data/input12.txt")
	assert.Equal(t, 111754, Day12{}.sumDocumentNumbers(input), "sumDocumentNumbers")
	assert.Equal(t, 65402, Day12{}.sumWithoutDoubleCountingRed(input), "sumWithoutDoubleCountingRed")
}

func TestDay13Solutions(t *testing.T) {
	input := util.ReadLines("data/input13.txt")
	happyWithoutMe, lessHappyWithMe := Day13{}.optimalHappinessWithMeOptional(input)
	assert.Equal(t, 733, happyWithoutMe, "happyWithoutMe")
	assert.Equal(t, 725, lessHappyWithMe, "lessHappyWithMe") // lolz
}

func TestDay14Solutions(t *testing.T) {
	input := util.ReadLines("data/input14.txt")
	distanceTraveled, pointsAwarded := Day14{}.hostReindeerOlympics(input)
	assert.Equal(t, 2696, distanceTraveled, "distanceTraveled")
	assert.Equal(t, 1084, pointsAwarded, "pointsAwarded")
}

func TestDay15Solutions(t *testing.T) {
	input := util.ReadLines("data/input15.txt")
	highestScore, highestScoreWithCaloricReq := Day15{}.findHighestScore(input, 500)
	assert.Equal(t, 18965440, highestScore, "findHighestScore(...)")
	assert.Equal(t, 15862900, highestScoreWithCaloricReq, "findHighestScore(500)")
}

func TestDay16Solutions(t *testing.T) {
	input := util.ReadLines("data/input16.txt")
	giftSue, realSue := Day16{}.identifyAuntSue(input)
	assert.Equal(t, 40, giftSue, "giftSue")
	assert.Equal(t, 241, realSue, "realSue")
}

func TestDay17Solutions(t *testing.T) {
	input := util.ReadNumbers("data/input17.txt")
	totalPermutations, totalWith150Litres := Day17{}.transferEggnog(input)
	assert.Equal(t, 1304, totalPermutations, "totalPermutations")
	assert.Equal(t, 18, totalWith150Litres, "totalWith150Litres")
}

func TestDay18Solutions(t *testing.T) {
	input := util.ReadLines("data/input18.txt")
	totalLights, totalWithBrokenCorners := Day18{}.totalLights(input)
	assert.Equal(t, 821, totalLights, "totalLights")
	assert.Equal(t, 886, totalWithBrokenCorners, "totalWithBrokenCorners")
}

func TestDay19Solutions(t *testing.T) {
	input := util.ReadLines("data/input19.txt")
	distinctMolecules, fewestSteps := Day19{}.createMolecules(input)
	assert.Equal(t, 509, distinctMolecules, "distinctMolecules")
	assert.Equal(t, 195, fewestSteps, "fewestSteps")
}

func TestDay20Solutions(t *testing.T) {
	house1, house2 := Day20{}.lowestHouseNumbers(34000000)
	assert.Equal(t, 786240, house1, "house1")
	assert.Equal(t, 831600, house2, "house2")
}

func TestDay21Solutions(t *testing.T) {
	bestOutfitCost, worstOutfitCost := Day21{}.fightBosses(Boss{109, 8, 2})
	assert.Equal(t, 111, bestOutfitCost, "bestOutfitCost")
	assert.Equal(t, 188, worstOutfitCost, "worstOutfitCost")
}

func TestDay22Solutions(t *testing.T) {
	leastMana, leastManaHard := Day22{}.wizardsFightBoss(Boss{71, 10, 0})
	assert.Equal(t, 1824, leastMana, "leastMana")
	assert.Equal(t, 1937, leastManaHard, "leastManaHard")
}
func TestDay23Solutions(t *testing.T) {
	input := util.ReadLines("data/input23.txt")
	startsWith0, startsWith1 := Day23{}.followComputerInstructions(input)
	assert.Equal(t, 307, startsWith0, "startsWith0")
	assert.Equal(t, 160, startsWith1, "startsWith1")
}

func TestDay24Solutions(t *testing.T) {
	input := util.ReadNumbers("data/input24.txt")
	threeCompartments, fourCompartments := Day24{}.findSmallestContainers(input, 3, 4)
	assert.Equal(t, 11846773891, threeCompartments, "threeCompartments")
	assert.Equal(t, 80393059, fourCompartments, "fourCompartments")
}
