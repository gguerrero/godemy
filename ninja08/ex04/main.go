package main

import (
	"fmt"
	"sort"
)

func main() {
	integerList := []int{1, 3, 2, 51, 12, 1245, 14, 41, 12, 214, 9, 1, 0, 7}
	stringList := []string{"Foo", "Bar", "bar", "she", "we", "beer"}

	sort.Ints(integerList)
	sort.Strings(stringList)

	fmt.Println(integerList)
	fmt.Println(stringList)
}
