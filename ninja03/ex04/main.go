package main

import (
	"fmt"
	"time"
)

func main() {
	bd := 1985

	for {
		if bd > time.Now().Year() {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
