package main

import "math"

type Shape interface {
	area() float64
	perimeter() float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) perimeter() float64 {
	return 2*math.Abs(r.x2-r.x1) + 2*math.Abs(r.y2-r.y1)
}

func (c *Circle) perimeter() float64 {
	return 2 * c.r * math.Pi
}
