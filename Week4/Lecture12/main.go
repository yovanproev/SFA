package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	out := generateThrottled("foo", 3, time.Second)
	for f := range out {
		log.Println(f)
	}
}

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	channel := make(chan string, bufferLimit)
	go func() {
		ticker := time.NewTicker(clearInterval)

		for {

			select {
			case <-ticker.C:
				for i := 0; i < bufferLimit; i++ {
					log.Println(data)
				}
			case <-channel:
				fmt.Println("done")
				return
			}

		}
	}()

	return channel
}

//Output buffer size 2:
// 2022/03/15 02:08:40 foo
// 2022/03/15 02:08:40 foo
// 2022/03/15 02:08:41 foo
// 2022/03/15 02:08:41 foo
// 2022/03/15 02:08:42 foo
// 2022/03/15 02:08:42 foo
// 2022/03/15 02:08:43 foo
// 2022/03/15 02:08:43 foo
// 2022/03/15 02:08:44 foo
// 2022/03/15 02:08:44 foo
// 2022/03/15 02:08:45 foo
// 2022/03/15 02:08:45 foo
// 2022/03/15 02:08:46 foo
// 2022/03/15 02:08:46 foo
// 2022/03/15 02:08:47 foo
// 2022/03/15 02:08:47 foo
// 2022/03/15 02:08:48 foo
// 2022/03/15 02:08:48 foo
// 2022/03/15 02:08:49 foo
// 2022/03/15 02:08:49 foo
// 2022/03/15 02:08:50 foo
// 2022/03/15 02:08:50 foo
// 2022/03/15 02:08:51 foo
// 2022/03/15 02:08:51 foo
// 2022/03/15 02:08:52 foo
// 2022/03/15 02:08:52 foo
// 2022/03/15 02:08:53 foo
// 2022/03/15 02:08:53 foo
// exit status 0xc000013a

// buffer size 3 Output

// 2022/03/15 02:09:49 foo
// 2022/03/15 02:09:49 foo
// 2022/03/15 02:09:49 foo
// 2022/03/15 02:09:50 foo
// 2022/03/15 02:09:50 foo
// 2022/03/15 02:09:50 foo
// 2022/03/15 02:09:51 foo
// 2022/03/15 02:09:51 foo
// 2022/03/15 02:09:51 foo
// 2022/03/15 02:09:52 foo
// 2022/03/15 02:09:52 foo
// 2022/03/15 02:09:52 foo
// exit status 0xc000013a
