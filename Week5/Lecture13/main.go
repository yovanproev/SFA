package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type BufferedContext struct {
	context.Context
	channel    chan string
	bufferSize int
}

func main() {

	ctx := NewBufferedContext(time.Second, 10)
	modifiedDone := ctx.Done()

	ctx.Run(func(ctx context.Context, buffer chan string) {

		for {
			select {
			case <-modifiedDone:
				fmt.Println(ctx.Err())
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 200) //try different values here
				fmt.Println("bar")
			}
		}
	})
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	out := make(chan string, bufferSize)

	return &BufferedContext{
		Context:    ctx,
		channel:    out,
		bufferSize: bufferSize,
	}
}

func (bc *BufferedContext) Done() <-chan struct{} {
	var wg sync.WaitGroup
	channel := bc.channel
	var cummulateChannel []string

	wg.Add(1)
	go func() {
		for i := 0; i < bc.bufferSize; i++ {
			defer wg.Done()
			cummulateChannel = append(cummulateChannel, <-channel)
		}
		if len(cummulateChannel) == bc.bufferSize {
			close(channel)
			fmt.Println("Channel closed because of overloading")
		}
		wg.Wait()
	}()
	return bc.Context.Done()
}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	fn(bc.Context, bc.channel)
}

// Output:
// bar
// bar
// bar
// bar
// bar
// bar
// bar
// bar
// bar
// Channel closed because of overloading
// bar
// context canceled
