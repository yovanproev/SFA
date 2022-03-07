package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	LargestArea() float64
}

type Shapes struct {
	Shape []Shape
}

func (s Shapes) LargestArea() float64 {
	var largestNumber []float64
	x := s.Shape
	for _, s := range x {
		largestNumber = append(largestNumber, s.Area())
	}

	fmt.Println("All shapes", largestNumber)
	var maxNumber float64
	for _, e := range largestNumber {
		if e > maxNumber {
			maxNumber = e
		}
	}

	return maxNumber
}

func (s Shapes) Area() float64 {
	return 3.14
}

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

func (s Square) LargestArea() float64 {
	return s.width * s.height
}

func (c Circle) LargestArea() float64 {
	return math.Pi * c.radius * c.radius
}

func measureLA(s Shape) {
	fmt.Println("LargestArea", s.LargestArea())
}

func measureA(s Shape) {
	fmt.Println("Area", s.Area())
}

func main() {
	s := Square{width: 5, height: 4}
	c := Circle{radius: 2}

	measureA(s)
	measureA(c)

	shape1 := []Shape{&Square{width: 5, height: 4}, &Circle{radius: 2}}
	measureLA(Shapes{shape1})
}

// Output:
// Area 20
// Area 12.566370614359172
// All shapes [20 12.566370614359172]
// LargestArea 20
