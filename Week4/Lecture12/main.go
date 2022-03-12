package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		log.Println(f)
	}
}

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	channel := make(chan string, bufferLimit)
	var wg sync.WaitGroup

	go func() {
		for {
			wg.Add(1)
			go func() {
				defer wg.Done()
				channel <- data
				channel <- data
				time.Sleep(clearInterval)
			}()
			wg.Wait()
		}
	}()

	return channel
}

//Output:
// 2022/03/12 12:11:12 foo
// 2022/03/12 12:11:12 foo
// 2022/03/12 12:11:13 foo
// 2022/03/12 12:11:13 foo
// 2022/03/12 12:11:14 foo
// 2022/03/12 12:11:14 foo
// 2022/03/12 12:11:15 foo
// 2022/03/12 12:11:15 foo
// 2022/03/12 12:11:16 foo
// 2022/03/12 12:11:16 foo
// exit status 0xc000013a
