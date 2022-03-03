// fizzbuzz
// A program that prints the numbers from 1 to n. But for multiples of three
// print “Fizz” instead of the number and for the multiples of five print “Buzz”.
// For numbers which are multiples of both three and five print “FizzBuzz”.
package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter integer: ")
	var input int
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(i int) {
	fizz := "fizz"
	buzz := "buzz"

	if i%3 == 0 && i%5 == 0 {
		fmt.Println(i, fizz+buzz)
	} else if i%3 == 0 {
		fmt.Println(i, fizz)
	} else if i%5 == 0 {
		fmt.Println(i, buzz)
	} else {
		fmt.Println(i)
	}
}
