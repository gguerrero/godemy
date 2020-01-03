package main

import (
	"fmt"
	"github.com/gguerrero/godemy/ninja13/ex01/godog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.ToDogYears(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.ToDogYears2(20))
}
