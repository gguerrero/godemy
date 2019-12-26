package main

import (
	"fmt"
)

func main() {
	f := func(s string) {
		fmt.Printf("Hello %v, from anounymous func!\n", s)
	}

	f("Guillermo GB")
}
