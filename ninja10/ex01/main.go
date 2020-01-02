package main

import (
	"fmt"
)

func main() {
	anonymousFunc()
	bufferedChannel()
}

func anonymousFunc(){
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println("Inside `anonymousFunc`:", <-c)
}

func bufferedChannel(){
	c := make(chan int, 1)

	c <- 42
	fmt.Println("Inside `bufferedChannel`:", <-c)
}
