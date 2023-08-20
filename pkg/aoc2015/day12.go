package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/12">Day 12</a>
 */

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kevansimpson/util"
)

func sumDocumentNumbers(input string) int {
	sum := 0
	for _, num := range util.ExtractInts(input) {
		sum += num
	}
	return sum
}

func sumWithoutDoubleCountingRed(input string) int {
	var doc map[string]interface{}
	err := json.Unmarshal([]byte(input), &doc)
	if err != nil {
		return -1
	}

	return sumJsonObj(doc)
}

func sumJsonObj(doc map[string]interface{}) int {
	sum := 0
	for _, element := range doc {
		switch json := element.(type) {
		case map[string]interface{}:
			sum += sumJsonObj(json)
		case []interface{}:
			sum += sumJsonArray(json)
		case string:
			if json == "red" {
				return 0
			}
		default:
			num, _ := strconv.Atoi(fmt.Sprintf("%v", json))
			sum += num
		}
	}

	return sum
}

func sumJsonArray(array []interface{}) int {
	sum := 0
	for _, element := range array {
		switch json := element.(type) {
		case map[string]interface{}:
			sum += sumJsonObj(json)
		case []interface{}:
			sum += sumJsonArray(json)
		case string:
			// do nothing
		default:
			num, _ := strconv.Atoi(fmt.Sprintf("%v", json))
			sum += num
		}
	}

	return sum
}
