package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	point1 := NewPoint(432, 31)
	point2 := NewPoint(243, 434)

	diff := point1.Distance(point2)

	fmt.Printf("Distance between is %.2f", diff)
}
