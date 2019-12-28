package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

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
	// For some reason, using the pointer value does create any race condition but ensure the final value
	// even though all the GoRoutines suppose to share the same address to increment the value :|
	// runtime.Gosched()
	// *counter++

	currentCounter := *counter

	runtime.Gosched()

	currentCounter++
	*counter = currentCounter

	fmt.Println("Running GoRoutine #", i, "with counter value", *counter)

	wg.Done()
}
