package main

import (
	"fmt"
)

const (
	y1 int = 2015
	y2 int = y1 + iota
	y3 int = y1 + iota
	y4 int = y2 + iota
)

func main() {
	fmt.Println(y1, y2, y3 ,y4)
}
