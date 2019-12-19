package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

type Circle struct {
	Radius float64
}

func (s Square) Area() float64 {
	return s.Length * 4
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * c.Radius
}

type Geometry interface {
	Area() float64
}

func main() {
	geometries := []Geometry{
		Square{20.5},
		Square{2.0},
		Circle{5.0},
		Circle{3.48},
		Circle{5.01},
	}

	for _, geometry := range geometries {
		area := geometry.Area()
		fmt.Printf("The area of the geometry is %v.\n", area)
	}
}
