package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

// type Shapes []Shape
// func (s Shapes) LargestArea() float64 {}

// Not sure how to use this type here....?????
// shape is an interface

type Square struct {
	width, height float64
}
type Circle struct {
	radius float64
}

//The implementation for Area().
func (s Square) Area() float64 {
	return s.width * s.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(s Shape) {
	fmt.Println("Area", s.Area())
}

func main() {
	s := Square{width: 3, height: 4}
	c := Circle{radius: 5}

	measure(s)
	measure(c)
}

// Output:
// 1 12
// 1 78.53981633974483
