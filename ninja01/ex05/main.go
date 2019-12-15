package main

import (
	"fmt"
)

type custInt int

var x custInt
var y int

func main() {
	fmt.Println(x)
	fmt.Printf(`Type of "x": %T%s`, x, "\n")

	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf(`Type of "y": %T%s`, y, "\n")
	fmt.Println(y)
}
