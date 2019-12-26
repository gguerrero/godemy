package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// MyWriter is my own specific wrter
type MyWriter os.File

func (MyWriter) Write(p []byte) (int, error) {
	writes := 0
	for i, b := range p {
		fmt.Printf("[%d] %s\n", i, string(b))
		writes++
	}

	return writes, nil
}

// Person struct
type Person struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Sayings   []string `json:"sayings"`
}

func main() {
	p1 := Person{
		FirstName: "Guillermo",
		LastName:  "GB",
		Age:       34,
		Sayings:   []string{"foo", "bar"},
	}

	p2 := Person{
		FirstName: "Irem",
		LastName:  "C",
		Age:       34,
		Sayings:   []string{"remove", "ion"},
	}

	var writer MyWriter
	encoder := json.NewEncoder(writer)
	encoder.Encode([]Person{p1, p2})
}
