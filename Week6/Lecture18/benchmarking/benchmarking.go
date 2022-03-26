package benchmarking

import (
	"sync"
	"time"
)

func PrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			res = append(res, i)
			time.Sleep(sleep)
		}
	}
	return res
}

func GoPrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	ch := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for i := 2; i < n; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				if n%idx == 0 {
					ch <- idx
					time.Sleep(sleep)
				}
			}(i)
		}
		wg.Wait()
	}()

	for i := 0; i < n; i++ {
		go func() {
			res = append(res, <-ch)
		}()
	}

	return res
}
