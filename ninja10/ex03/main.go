package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	go generate(c, 150)
	receive(c)

	fmt.Println("End of main...")
}

func generate(c chan<- int, times int) {
	for i :=0; i < times; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for value := range c {
		fmt.Printf("Value received from channel %v: %v\n", c, value)
	}
}
