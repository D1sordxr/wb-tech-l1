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
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	pointA := NewPoint(1.11, 2.22)
	pointB := NewPoint(4.44, 6.67)

	distance := pointA.Distance(pointB)
	fmt.Printf("Distance between A(%.2f, %.2f) and B(%.2f, %.2f) is %.2f\n",
		pointA.x, pointA.y,
		pointB.x, pointB.y,
		distance,
	)
}
