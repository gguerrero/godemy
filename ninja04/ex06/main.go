package main

import (
	"fmt"
)

func main() {
	emirates := make([]string, 3, 7)
	emirates[0] = "Dubai"
	emirates[1] = "Sharjah"
	emirates[2] = "Fujairah"

	fmt.Println(emirates)
	fmt.Println(len(emirates))
	fmt.Println(cap(emirates))

	for i := 0; i < len(emirates); i++ {
		fmt.Printf("[%d] %v\n", i, emirates[i])
	}
}
