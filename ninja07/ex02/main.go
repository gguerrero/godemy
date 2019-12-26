package main

import (
	"fmt"
	"strings"
)

// Human interface for Rename Method sets
type Human interface {
	Rename(string)
}

// Person struct my friend
type Person struct {
	FirstName string
	LastName  string
}

// Rename function helps rename and modify a Person struct
func (person *Person) Rename(fullName string) {
	fmt.Printf("Person type inside 'rename' function is %T\n", person)
	firstAndLastName := strings.Split(fullName, " ")

	// Same thing to access with (*variable).field selector than variable.field
	(*person).FirstName = firstAndLastName[0]
	person.LastName = firstAndLastName[1]
}

func main() {
	p := Person{
		FirstName: "Guillermo",
		LastName:  "GB",
	}

	fmt.Println(p)

	p.Rename("Jaime G")
	fmt.Println(p)

	// Need to use the address to person when received argument is it's interface
	renameHuman(&p, "Maite B")
	fmt.Println(p)
}

func renameHuman(human Human, fullName string) {
	fmt.Printf("Human type inside 'renameHuman' function is %T\n", human)
	human.Rename(fullName)
}
