package main

import (
	"fmt"
)

type custInt int

var x custInt

func main() {
	fmt.Println(x)
	fmt.Printf(`Type of "x": %T%s`, x, "\n")

	x = 42

	fmt.Println(x)
}
