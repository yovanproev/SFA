// package main

// import "fmt"

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} // Initialization of the slice a and 0 < n < len(a) - 1.
// 	difs := make(chan int, 3)

// 	go routine(a[0:1], difs)
// 	go routine(a[2:3], difs)
// 	go routine(a[8:12], difs)

// 	result := []int{<-difs, <-difs, <-difs}

// 	fmt.Println(result) // Display the first result returned by one of the routine.
// }

// func routine(a []int, out chan<- int) {
// 	result := 0 // Long computation.
// 	for _, v := range a {
// 		result += v
// 	}
// 	out <- result
// }

package main

import (
	"fmt"
	"sync"
)

func ToSlice(c chan interface{}) []interface{} {
	s := make([]interface{}, 0)
	for i := range c {
		s = append(s, i)
	}
	return s
}

func main() {
	c := make(chan interface{})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		r := ToSlice(c)
		fmt.Printf("Result: %#v\n", r)
		for _, e := range r {
			if ei := e.(int); ei > 11 {
				fmt.Println("e is", ei)
			}
		}
		wg.Done()
	}()

	c <- 10
	c <- 12
	c <- 20
	close(c)

	wg.Wait()

}
