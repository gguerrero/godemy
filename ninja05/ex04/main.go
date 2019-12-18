package main

import (
	"fmt"
	"strings"
)

func main() {
	paella := struct{
		Ingredients []string
		Steps map[string]string
	}{
		Ingredients: []string{"Onion", "Rice", "Chicken"},
		Steps: map[string]string{
			"Step 1": "Chop the onion and put it in the hot pan with olive oil.",
			"Step 2": "Mix the rice and chicken, cover with boiling water and wait 20 mins.",
		},
	}

	fmt.Printf("Let's cook a paella with: %v\n", strings.Join(paella.Ingredients, ", "))
	fmt.Println("--------------------------------------------->")
	for step, instruction := range paella.Steps {
		fmt.Printf("%v:\t%v\n", step, instruction)
	}
}
