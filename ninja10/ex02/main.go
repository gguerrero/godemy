package main

import (
	"fmt"
)

func main() {
	bidirectionalChannel := make(chan int)
	senderChannel        := (chan<- int)(bidirectionalChannel)
	receiverChannel      := (<-chan int)(bidirectionalChannel)

	go func(){
		senderChannel <- 42

		fmt.Printf("------\n")
		fmt.Printf("channel\t%T\n", senderChannel)
	}()

	fmt.Println(<-receiverChannel)

	fmt.Printf("------\n")
	fmt.Printf("channel\t%T\n", receiverChannel)
}
