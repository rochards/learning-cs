package todo

import "fmt"

type Todo struct {
	Description string
	Completed   bool
}

var todos []Todo

func AddTodo(description string) {
	todo := Todo{Description: description, Completed: false}
	todos = append(todos, todo)
	fmt.Println("Added: ", todo.Description)
}

func ListTodos() {

	if len(todos) == 0 {
		fmt.Println("No to-dos yet!")
	}

	for i, todo := range todos {
		status := "Incomplete"
		if todo.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s - [%s]\n", i+1, todo.Description, status)
	}
}

func CompleteTodo(index int) {
	if index < 1 || index > len(todos) {
		fmt.Println("Invalid index")
		return
	}

	todos[index-1].Completed = true
	fmt.Println("Marked as completed: ", todos[index-1].Description)
}

func DeleteTodo(index int) {

	if index < 1 || index > len(todos) {
		fmt.Println("Invalid index")
		return
	}

	deleted := todos[index-1]
	todos = append(todos[:index-1], todos[index:]...)
	fmt.Println("Deleted: ", deleted.Description)
}
