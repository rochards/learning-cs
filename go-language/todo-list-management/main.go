package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Description string
	Completed   bool
}

var todos []Todo

func addTodo(description string) {
	todo := Todo{Description: description, Completed: false}
	todos = append(todos, todo)
	fmt.Println("Added: ", description)
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No to-dos yet!")
		return
	}

	for i, todo := range todos {
		status := "Incomplete"
		if todo.Completed {
			status = "Complete"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, todo.Description, status)
	}
}

func completeTodo(index int) {
	if index < 1 || index > len(todos) {
		fmt.Println("Invalid index. Try again")
		return
	}

	todos[index-1].Completed = true
	fmt.Println("Marked as completed: ", todos[index-1].Description)
}

func deleteTodo(index int) {
	if index < 1 || index > len(todos) {
		fmt.Println("Invalid index. Try again")
		return
	}

	deleted := todos[index-1]
	todos = append(todos[:index-1], todos[index:]...)
	fmt.Println("Deleted:", deleted.Description)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("Select an option: add, list, complete, delete, quit")
		var option string
		fmt.Scan(&option)

		switch option {
		case "add":
			fmt.Println("Enter a description:")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			addTodo(description)

		case "list":
			listTodos()

		case "complete":
			fmt.Println("Enter the number of the to-do to complete")
			var index int
			fmt.Scan(&index)
			completeTodo(index)

		case "delete":
			fmt.Println("Enter the number of the to-do to delete")
			var index int
			fmt.Scan(&index)
			deleteTodo(index)

		case "quit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Unknown option")
		}
	}
}
