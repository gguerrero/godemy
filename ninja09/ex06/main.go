package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Architecture:\t\t", runtime.GOARCH)
	fmt.Println("OS:\t\t\t", runtime.GOOS)
}
