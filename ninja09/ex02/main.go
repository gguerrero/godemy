package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area method
func (c *Circle) Area() float64 {
	return c.Radius * 2 * math.Pi
}

func infoArea(s Shape) { fmt.Println("Area of shape is", s.Area()) }

func main() {
	c := Circle{10.5}
	infoArea(&c)
	fmt.Println("Area of this circle is", c.Area())
}
