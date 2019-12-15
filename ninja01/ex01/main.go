package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "Bond, James"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("Type of x is %T\n", x)
	fmt.Printf("Type of y is %T\n", y)
	fmt.Printf("Type of z is %T\n", z)
}
