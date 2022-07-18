package Point

type Point struct {
	x, y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func CalcDistance(p1 *Point, p2 *Point) int {
	return (p2.x - p1.x) ^ 2 + (p2.y - p1.y) ^ 2
}
