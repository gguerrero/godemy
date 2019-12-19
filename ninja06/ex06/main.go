package main

import (
	"fmt"
)

func main() {
	func(s string) {
		fmt.Printf("Hello %v, from anounymous func!\n", s)
	}("Guillermo Guerrero")
}
