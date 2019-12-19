package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 15, 16, 42, 82}

	fmt.Println(foo(nums...))
	fmt.Println(bar(nums))
}

func foo(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum
}

func bar(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum
}
