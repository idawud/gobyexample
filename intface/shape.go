package intface

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width, height float64
}

func NewRect(width, height float64) *Rect {
	r := Rect{width, height}
	return &r
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	c := Circle{radius}
	return &c
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func Measure(g Geometry) {
	fmt.Println("Method: ", g)
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimeter: ", g.Perimeter())
}
