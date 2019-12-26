package main

import (
	"fmt"
)

// Person struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Speak function
func (person Person) Speak() {
	fullName := person.FirstName + " " + person.LastName
	fmt.Printf("Hey there, I'm %s and I'm %d years old.\n", fullName, person.Age)
}

func main() {
	laura := Person{"Laura", "G", 23}
	laura.Speak()
}
