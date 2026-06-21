package main

import "fmt"

type Todo struct {
	name        string
	date        string
	description string
}

func (todo *Todo) setName(name string) {
	todo.name = name
}

func (todo *Todo) setDate(date string) {
	todo.date = date
}

func (todo *Todo) setDescription(description string) {
	todo.description = description
}

func (todo Todo) String() string {
	return fmt.Sprintf("Name: %s \nDate: %s \nDescription: %s\n", 
		todo.name, todo.date, todo.description)
}

type List struct {
	owner string
	name  string
	items []Todo
}

func (lst *List) removeTodo(idx int) {
	fmt.Println("removing todo #", idx)

	lst.items = append(lst.items[:idx], lst.items[idx+1:]...)
}

func (lst *List) addTodo(name string, date string, description string) {
	newIdx := len(lst.items)
	fmt.Println("adding todo #", newIdx)

	lst.items = append(lst.items, Todo{name, date, description})
}

func (lst List) displayList() {
	for i, item := range lst.items {
		fmt.Println("Item #", i)
		fmt.Println(item)
	}
}

func main() {
	list := List{"James","new todo", nil}
	
	list.addTodo("Text Sunny", "Sunday", "Text Sunny Good Morning")
	list.addTodo("Father's Day", "Sunday", "Tell Dad Happy Father's Day")
	list.addTodo("Breakfast", "Monday", "Eat Breakfast")

	list.displayList()

	list.removeTodo(1)

	list.displayList()
}