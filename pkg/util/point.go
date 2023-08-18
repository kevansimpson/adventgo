package util

type Point struct {
	x, y int
}

type Point3D struct {
	x, y, z int
}

type BigPoint struct {
	x, y int64
}

var ORIGIN = Point{0, 0}

func (pt Point) Move(dx int, dy int) Point {
	return Point{pt.x + dx, pt.y + dy}
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
