package main

import (
	"benchmarking/benchmarking"
	"fmt"
	"time"
)

func main() {
	fmt.Println("1", benchmarking.PrimesAndSleep(100, 1*time.Millisecond))
	fmt.Println("2", benchmarking.GoPrimesAndSleep(100, 1*time.Millisecond))
}

// Output
// 1 [2 4 5 10 20 25 50]
// 2 [2 4 5 10 20 25]
