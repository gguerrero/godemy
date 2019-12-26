package main

import "fmt"

func main() {
	x := "Hello World!"

	fmt.Printf("For the variable x, with value %v, it's address is %v\n", x, &x)
}
