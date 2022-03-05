package main

type Shapes []Shape

type Shape interface {
	Area()
}

type ReceiverType struct {
}

type Square struct {
	NewSquare int
}

type Circle interface {
	Area() float64
}

func (s Shapes) LargestArea() float64 {
	return 3.14
}

func (r *ReceiverType) Area() float64 {
	return 3.14
}

func main() {

}
