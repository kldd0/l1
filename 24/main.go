package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func New(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func distance(p1, p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := New(1.43, 5.89)
	p2 := New(2.14, -9.23)

	dist := distance(p1, p2)
	fmt.Printf("расстояние между точками: %.2f\n", dist)
}
