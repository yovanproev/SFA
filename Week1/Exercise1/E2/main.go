package main

import (
	"fmt"
)

func main() {
	var inputMax, inputMin int

	fmt.Printf("place your Min range: ")
	fmt.Scanf("%d", &inputMin)

	fmt.Printf("place your Max range: ")
	fmt.Scanf("%d", &inputMax)

	if inputMax != 0 {
		makeRange(inputMin, inputMax)
	}

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	fmt.Println(a)
	return a
}
