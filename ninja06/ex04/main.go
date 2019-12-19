package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}

func (person Person) Speak() {
	fullName := person.FirstName + " " + person.LastName
	fmt.Printf("Hey there, I'm %s and I'm %d years old.\n", fullName, person.Age)
}

func main() {
	laura := Person{"Laura", "Guerero", 23}
	laura.Speak()
}
