package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var counter int64

	const gs = 250
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go readAndIncrement(i, &counter)
	}

	wg.Wait()
	fmt.Println("Final value for counter is", counter)
}

func readAndIncrement(i int, counter *int64) {
	atomic.AddInt64(counter, 1)

	runtime.Gosched()

	fmt.Println("Running GoRoutine #", i, "with counter value", atomic.LoadInt64(counter))

	wg.Done()
}
