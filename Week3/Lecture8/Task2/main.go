package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Shapes []Shape

func (s Shapes) LargestArea() float64 {
	var largestNumber []float64

	for _, s := range s {
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

func measureA(s Shape) {
	fmt.Println("Area", s.Area())
}

func measureLA(s Shapes) {
	fmt.Println("Lar.Area", s.LargestArea())
}

func main() {
	s := Square{
		width:  5,
		height: 3,
	}
	c := Circle{
		radius: 5,
	}

	measureA(s)
	measureA(c)

	var shapes = Shapes{s, c}

	measureLA(shapes)
}

// Output:
// Area 15
// Area 78.53981633974483
// All shapes [15 78.53981633974483]
// Lar.Area 78.53981633974483
