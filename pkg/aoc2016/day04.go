package aoc2016

/**
 * <a href="https://adventofcode.com/2016/day/4">Day 4</a>
 */

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/kevansimpson/util"
)

type Day04 struct{}

func (d Day04) identifyRealRooms(input []string) (int, int) {
	roomList := d.scanRooms(input)
	real, northPoleRoomId := 0, -1
	for _, room := range roomList {
		if room.isReal() {
			real += room.sectorId
		}

		shift := util.ShiftText(room.name, room.sectorId)
		if strings.Contains(shift, "northpole") {
			northPoleRoomId = room.sectorId
		}
	}
	return real, northPoleRoomId
}

func (d Day04) scanRooms(input []string) []bunnyRoom {
	regex := regexp.MustCompile(`([\-a-z]+)-(\d+)\[([a-z]+)]`)
	rooms := make([]bunnyRoom, len(input))
	for i, scan := range input {
		m := regex.FindStringSubmatch(scan)
		sid, _ := strconv.Atoi(m[2])
		rooms[i] = bunnyRoom{m[1], sid, m[3]}
	}
	return rooms
}

type bunnyRoom struct {
	name     string
	sectorId int
	checksum string
}

func (r bunnyRoom) isReal() bool {
	counts := make(map[rune]int)
	for _, ch := range strings.ReplaceAll(r.name, "-", "") {
		counts[ch] = strings.Count(r.name, string(ch))
	}

	pairs := make([][2]interface{}, 0, len(counts))
	for k, v := range counts {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		comp := pairs[i][1].(int) - pairs[j][1].(int)
		if comp == 0 {
			return pairs[i][0].(rune) < pairs[j][0].(rune)
		} else {
			return comp > 0
		}
	})

	var code []rune
	for i := 0; i < 5; i++ {
		code = append(code, pairs[i][0].(rune))
	}

	return string(code) == r.checksum
}
