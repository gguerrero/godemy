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
		Lastname:     "G",
		FavIceCreams: []string{"Chocolate", "Vanilla"},
	}

	nayri := Person{
		Name:         "Nayri",
		Lastname:     "B",
		FavIceCreams: []string{"Vanilla", "Strawberry"},
	}

	persons := map[string]Person{
		me.Lastname:    me,
		nayri.Lastname: nayri,
	}

	for k, v := range persons {
		fmt.Printf("Person %v:\t%+v\n", k, v)
	}
}
