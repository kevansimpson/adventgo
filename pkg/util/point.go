package util

import (
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Point3D struct {
	x, y, z int
}

type BigPoint struct {
	x, y int64
}

var ORIGIN = Point{0, 0}

func (pt Point) Move(dx int, dy int) Point {
	return Point{pt.X + dx, pt.Y + dy}
}

func (pt Point) RuneStep(dir rune) Point {
	return pt.Advance(string(dir), 1)
}

func (pt Point) Step(dir string) Point {
	return pt.Advance(dir, 1)
}

func (pt Point) Advance(dir string, distance int) Point {
	switch dir {
	case ">", "R", "E":
		return pt.Move(distance, 0)
	case "<", "L", "W":
		return pt.Move(-distance, 0)
	case "^", "U", "N":
		return pt.Move(0, distance)
	case "v", "D", "S":
		return pt.Move(0, -distance)
	default:
		return pt
	}
}

func (pt Point) ManhattanDistance(dest Point) int {
	return AbsInt(dest.X-pt.X) + AbsInt(dest.Y-pt.Y)
}

func (pt Point) ManhattanDistanceOrigin() int {
	return AbsInt(pt.X) + AbsInt(pt.Y)
}

// factory functions

func MakeBigPoint(str string) BigPoint {
	xy := strings.Split(str, ",")
	x, _ := strconv.ParseInt(xy[0], 10, 64)
	y, _ := strconv.ParseInt(xy[1], 10, 64)
	return BigPoint{x, y}
}

func MakePoint(str string) Point {
	xy := strings.Split(str, ",")
	x, _ := strconv.ParseInt(xy[0], 10, 32)
	y, _ := strconv.ParseInt(xy[1], 10, 32)
	return Point{int(x), int(y)}
}
