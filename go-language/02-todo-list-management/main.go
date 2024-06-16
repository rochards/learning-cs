package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-list-management/todo"
)

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
			todo.AddTodo(description)

		case "list":
			todo.ListTodos()

		case "complete":
			fmt.Println("Enter the number of the to-do to complete")
			var index int
			fmt.Scan(&index)
			todo.CompleteTodo(index)

		case "delete":
			fmt.Println("Enter the number of the to-do to delete")
			var index int
			fmt.Scan(&index)
			todo.DeleteTodo(index)

		case "quit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Unknown option")
		}
	}
}
