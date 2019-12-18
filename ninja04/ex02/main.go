package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30, 40, 50}

	for i, v := range a {
		fmt.Printf("[%d] %d\n", i, v)
	}
	fmt.Printf("Type of 'a' is %T\n", a)
}
