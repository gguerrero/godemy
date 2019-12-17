package main

import (
	"fmt"
)

func main() {
	x := 42

	fmt.Printf("Base2:  %b\n",  x)
	fmt.Printf("Base10: %d\n",  x)
	fmt.Printf("Base16: %#X\n", x)
}
