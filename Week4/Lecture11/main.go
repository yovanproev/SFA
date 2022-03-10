package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var inputs = []int{1, 17, 34, 56, 2, 8, 106, 111, 112, 1050}
var wg sync.WaitGroup

func main() {
	evenChan := processEven(inputs)
	for number := range evenChan {
		fmt.Println("Even num: ", number)
	}

	oddCnan := processOdd(inputs)
	for number := range oddCnan {
		fmt.Println("Odd num: ", number)
	}

	time.Sleep(time.Second * 1)
}

func processEven(inputs []int) chan int {
	evenNumChan := make(chan int)

	go func() {
		log.Println("start")
		for _, input := range inputs {
			wg.Add(1)

			go func(number int) {
				defer wg.Done()

				if number%2 == 0 {
					evenNumChan <- number
				}
			}(input)
		}

		wg.Wait()
		log.Println("finish")
		close(evenNumChan)
	}()
	return evenNumChan
}

func processOdd(inputs []int) chan int {
	oddNumChan := make(chan int)

	go func() {
		log.Println("start")
		for _, input := range inputs {
			wg.Add(1)

			go func(number int) {
				defer wg.Done()

				if number%2 != 0 {
					oddNumChan <- number
				}
			}(input)
		}

		wg.Wait()
		log.Println("finish")
		close(oddNumChan)
	}()
	return oddNumChan
}

// Output:
// 2022/03/10 14:41:33 start
// Even num:  34
// Even num:  1050
// Even num:  56
// Even num:  2
// Even num:  8
// Even num:  106
// Even num:  112
// 2022/03/10 14:41:33 finish
// 2022/03/10 14:41:33 start
// Odd num:  1
// Odd num:  17
// Odd num:  111
// 2022/03/10 14:41:33 finish
