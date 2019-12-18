package main

import (
	"fmt"
)

func main() {
	a := []string{"a1", "a2"}
	b := []string{"b1", "b2", "b3"}

	table := [][]string{a, b}
	fmt.Println(table)
}
