package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.RWMutex
	counter int
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.Add(2)
	cp.counter = 1

	cp.printFoo(times)
	cp.printBar(times)

	time.Sleep(10 * time.Millisecond)
	cp.Wait()
}

func (cp *ConcurrentPrinter) printFoo(times int) {
	go func() {
		defer cp.Done()
		for times > cp.counter {
			cp.RLock()
			fmt.Print("foo")
			time.Sleep(time.Second)
			cp.RUnlock()
			cp.counter++
		}
	}()
}

func (cp *ConcurrentPrinter) printBar(times int) {
	defer cp.Done()
	for times > cp.counter {
		time.Sleep(time.Second * 1 / 2)
		cp.Lock()
		fmt.Print("bar")
		time.Sleep(time.Second)
		cp.Unlock()
		cp.counter++
	}
}

// Output: foobarfoobarfoobarfoobarfoobar
