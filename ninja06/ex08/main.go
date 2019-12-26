package main

import (
	"fmt"
)

func main() {
	f := createFunc()

	f("Guillermo GB")
}

func createFunc() func(string) {
	return func(s string) {
		fmt.Printf("Hello %v, from anounymous func!\n", s)
	}
}
