package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("%T: %v\n", s, s.area())
}

func main() {
	s1 := square{
		length: 10,
		width:  20,
	}

	info(s1)

	c1 := circle{
		radius: 15,
	}

	info(c1)
}
