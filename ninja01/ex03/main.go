package main

import (
	"fmt"
)

var x = 42
var y = "Bond, James"
var z = true

func main() {
	s1 := fmt.Sprintf("x: %v - y: %v - z: %v", x, y, z)
	s2 := fmt.Sprintf("x (type): %T - y (type): %T - z (type): %T", x, y, z)
	s3 := fmt.Sprintf("%s is %d and that is %t", y, x, z)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
