package composite

import (
	"fmt"
	"math"
)

type Anything interface{}


type HasDistance interface{
	DistanceFromZero() float64
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (point Point)String() string {
	return fmt.Sprintf("(%d| %d)", point.X, point.Y)
}

func (point Point) DistanceFromZero() float64 {
	return math.Sqrt(
		math.Pow(float64(point.X), 2) + math.Pow(float64(point.Y), 2))
}
