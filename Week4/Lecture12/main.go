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
