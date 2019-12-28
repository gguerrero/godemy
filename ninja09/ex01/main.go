package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	const gs = 10000
	wg.Add(gs)

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	for i := 0; i < gs; i++ {
		go runGoRun(i)
	}

	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Done.")
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
}

func runGoRun(i int) {
	fmt.Println("GoRoutine #", i)
	wg.Done()
}
