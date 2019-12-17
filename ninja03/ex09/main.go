package main

import (
	"fmt"
)

func main() {
	favSport := "Squash"
	switch favSport {
		case "Cycling":
			fmt.Println("Go to the mountains")
		case "Squash":
			fmt.Println("Go to my place")
		default:
			fmt.Println("mmmm... nothing to say!")
	}
}
