package main

import (
	"fmt"
)

type Person struct {
	Name         string
	Lastname     string
	FavIceCreams []string
}

func main() {
	me := Person{
		Name:         "Guillermo",
		Lastname:     "GB",
		FavIceCreams: []string{"Chocolate", "Vanilla"},
	}

	fmt.Printf("%+v\n", me)

	for _, v := range me.FavIceCreams {
		fmt.Printf("An Ice Cream Flavor: %v\n", v)
	}
}
