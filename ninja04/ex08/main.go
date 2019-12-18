package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`James`: []string{`red`, `blue`, `orange`},
		`Will`:  []string{`red`, `green`, `yellow`},
	}

	fmt.Println(m)
	for key, colors := range(m) {
		fmt.Println(key)
		for i, value := range colors {
			fmt.Printf("\t[%d] %v\n", i, value)
		}
	}
}
