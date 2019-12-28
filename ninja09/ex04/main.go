package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	counter := 0

	const gs = 150
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go readAndIncrement(i, &counter)
	}

	wg.Wait()
	fmt.Println("Final value for counter is", counter)
}

func readAndIncrement(i int, counter *int) {
	mutex.Lock()

	currentCounter := *counter

	runtime.Gosched()

	currentCounter++
	*counter = currentCounter
	fmt.Println("Running GoRoutine #", i, "with counter value", *counter)

	mutex.Unlock()
	wg.Done()
}
