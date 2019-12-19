package main

import (
	"fmt"
)

func main() {
	todos := TodoList()
	anotherTodos := TodoList()

	fmt.Println(todos("Task #1"))
	fmt.Println(todos("Task #2"))
	fmt.Println(todos("Task #3"))
	fmt.Println(todos("Task #4"))
	fmt.Println(todos("Task #5"))
	fmt.Println(todos("Task #6"))

	fmt.Println(anotherTodos("Another Task #1"))
	fmt.Println(anotherTodos("Another Task #2"))
	fmt.Println(anotherTodos("Another Task #3"))
	fmt.Println(anotherTodos("Another Task #4"))
}

func TodoList() func(string) []string {
	todos := make([]string, 0, 5)

	return func(t string) []string {
		todos = append(todos, t)
		return todos
	}
}
