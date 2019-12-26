package main

import (
	"encoding/json"
	"fmt"
)

// Person struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := Person{
		Name: "Guillermo",
		Age:  34,
	}

	p2 := Person{
		Name: "Nayri",
		Age:  33,
	}

	bs, err := json.Marshal([]Person{p1, p2})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}
}
