package main

import (
	"fmt"
)

func main() {
	if 1 == 2 {
		fmt.Println("1 == 2")
	} else if 1 == 1 {
		fmt.Println("1 == 1")
	} else {
		fmt.Println("1 != 1")
	}
}
