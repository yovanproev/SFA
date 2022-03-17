package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	channel    chan string
	bufferSize int
	context.CancelFunc
}

func main() {

	ctx := NewBufferedContext(time.Second, 10)
	modifiedDone := ctx.Done()

	ctx.Run(func(ctx context.Context, buffer chan string) {
		defer func() {
			// recover from panic caused by writing to a closed channel
			if r := recover(); r != nil {
				return
			}
		}()

		for {
			select {
			case <-modifiedDone:
				fmt.Println(ctx.Err())
				return
			case buffer <- "bar":
				fmt.Println("bar")
				time.Sleep(time.Millisecond * 200) //try different values here
			}
		}
	})
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	out := make(chan string, bufferSize)

	return &BufferedContext{
		Context:    ctx,
		channel:    out,
		bufferSize: bufferSize,
		CancelFunc: cancel,
	}
}

func (bc *BufferedContext) Done() <-chan struct{} {
	channel := bc.channel
	var cummulateChannel []string

	go func() {
		for i := 0; i < bc.bufferSize; i++ {
			cummulateChannel = append(cummulateChannel, <-channel)
		}
		if len(cummulateChannel) == bc.bufferSize {
			close(channel)
			bc.CancelFunc()
			fmt.Printf("channel closed, buffer size of %v reached", bc.bufferSize)
		}
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
// bar
// channel closed, buffer size of 10 reached

// Output:
// bar
// bar
// context deadline exceeded

// Output:
// bar
// bar
// bar
// bar
// bar
// bar
// context deadline exceeded

// etc....
