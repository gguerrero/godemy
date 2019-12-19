package main

import (
	"fmt"
)

func main() {
	fileName := "/tmp/foo.bar"

	if open(fileName) {
		defer close(fileName)
		doStuff(fileName)
	} else {
		fmt.Println("Cannot open file!")
	}
}

func open(f string) bool {
	fmt.Println("Openning file:", f)

	return true
}

func close(f string) {
	fmt.Println("File closed:", f)
}

func doStuff(f string) {
	fmt.Println("Doing stuff on file:", f)
}
