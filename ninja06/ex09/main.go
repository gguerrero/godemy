package main

import (
	"fmt"
	"strings"
)

func main() {
	f := createFunc()
	runFunctionWithParams(f, "Guillermo", "G", "B")
}

func createFunc() func(...string) {
	return func(words ...string) {
		fmt.Printf("Some words: %v\n", strings.Join(words, ", "))
	}
}

func runFunctionWithParams(f func(...string), words ...string) {
	f(words...)
}
