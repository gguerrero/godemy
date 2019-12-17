package main

import (
	"fmt"
	"time"
)

func main() {
	bd := 1985

	for bd <= time.Now().Year() {
		fmt.Println(bd)
		bd++
	}
}
