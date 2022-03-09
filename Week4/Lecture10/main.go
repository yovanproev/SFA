package main

import (
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	condition int
}

var cond *sync.Cond = sync.NewCond(new(sync.Mutex))

func main() {
	times := 5
	cp := &ConcurrentPrinter{}

	cp.printFoo(times)
	cp.printBar(times)

	time.Sleep(10 * time.Millisecond)

}

func (cp *ConcurrentPrinter) printFoo(times int) {
	go func() {
		for i := 0; i < times; i++ {
			cond.L.Lock()
			for cp.condition == 0 {
				cond.Wait()
			}
			cp.condition = cp.condition - 1
			print("bar")
			cond.Signal()
			cond.L.Unlock()
		}
	}()
}

func (cp *ConcurrentPrinter) printBar(times int) {
	for i := 0; i < times; i++ {
		time.Sleep(time.Second)
		cond.L.Lock()
		for cp.condition == 3 {
			cond.Wait()
		}
		cp.condition = cp.condition + 1
		print("foo")
		cond.Signal()
		cond.L.Unlock()
	}
}

// Output: foobarfoobarfoobarfoobarfoobar
