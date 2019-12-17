package main

import (
	"fmt"
)

func main() {
	switch {
		case 1 == 1:
			fmt.Println("case 1 == 1")
			fallthrough
		case false:
			fmt.Println("case false")
			fallthrough
		case true:
			fmt.Println("case !true")
			fallthrough
		case !true:
			fmt.Println("case !true")
			fallthrough
		default:
			fmt.Println("default case")
	}
}
