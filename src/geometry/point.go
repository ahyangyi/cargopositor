package geometry

type Point struct {
	X, Y, Z int
}

type PointWithColour struct {
	Point  Point
	Colour byte
}

type Bounds struct {
	Min, Max Point
}
