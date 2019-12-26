package main

import (
	"fmt"
	"sort"
)

// Person struct
type Person struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Sayings   []string `json:"sayings"`
}

func (p *Person) sortSayings() {
	sort.Strings(p.Sayings)
}

type byFirstName []Person

func (people byFirstName) Len() int {
	return len(people)
}

func (people byFirstName) Less(i, j int) bool {
	return people[i].FirstName < people[j].FirstName
}

func (people byFirstName) Swap(i, j int) {
	people[i], people[j] = people[j], people[i]
}

type byAge []Person

func (people byAge) Len() int {
	return len(people)
}

func (people byAge) Less(i, j int) bool {
	return people[i].Age < people[j].Age
}

func (people byAge) Swap(i, j int) {
	people[i], people[j] = people[j], people[i]
}

func main() {
	p1 := Person{
		FirstName: "Nayri",
		LastName:  "B",
		Age:       33,
		Sayings:   []string{"so am I", "Who knows that"},
	}

	p2 := Person{
		FirstName: "Guillermo",
		LastName:  "GB",
		Age:       34,
		Sayings:   []string{"foo", "bar"},
	}

	p3 := Person{
		FirstName: "Irem",
		LastName:  "C",
		Age:       21,
		Sayings:   []string{"remove", "ion"},
	}

	fmt.Println("-------> Sorting people byFirstName")
	people := []Person{p1, p2, p3}
	fmt.Println(people)
	sort.Sort(byFirstName(people))
	fmt.Println(people)

	fmt.Println("-------> Sorting people byAge")
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)

	fmt.Println("-------> Sorting people's Sayings")
	p1.sortSayings()
	p2.sortSayings()
	p3.sortSayings()

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
