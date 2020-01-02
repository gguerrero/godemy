package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	q := make(chan bool)

	go generate(c, q, 150)
	receive(c, q)

	fmt.Println("End of main...")
}

func generate(c chan<- int, quit chan<- bool, times int) {
	for i :=0; i < times; i++ {
		c <- i
	}
	quit <- true
	close(quit)
	close(c)
}

func receive(c <-chan int, q <-chan bool) {
	for {
		select {
		case value := <- c:
			fmt.Println("Pulling out value from channel:", value)
		case quitValue := <- q:
			if quitValue {
				fmt.Println("Quiting the program...")
			}
			return
		}
	}
}
