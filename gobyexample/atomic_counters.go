package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func incrementCounter(workerID int, counter *uint64, times int, waitGroup *sync.WaitGroup) {
    for i := 0; i < times; i++ {
        fmt.Println("Worker:", workerID, "| Iteration:", i)
        atomic.AddUint64(counter, 1)
    }
    waitGroup.Done()
}

func main() {

    var ops uint64

    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)
        go incrementCounter(i, &ops, 10, &wg)
    }

    wg.Wait()

    fmt.Println("ops:", ops)
}
