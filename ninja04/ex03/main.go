package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(a[:5])
	fmt.Println(a[5:])
	fmt.Println(a[3:8])
	fmt.Println(a[4:9])
}
