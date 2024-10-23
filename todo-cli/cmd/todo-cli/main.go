package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"todo-cli/internal/db"
	"todo-cli/internal/todo_flag"
)

func main() {

	d := db.InitDB("todo.db")
	defer d.Close()

	todoRepository := db.NewTodoRepository(d)
	todoFlagService := todo_flag.NewTodoFlagService(todoRepository)

	if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Println("Expected one of: add, update, delete, list, get")
		fmt.Println("Usage: todo-cli [command] [flags]")
		fmt.Println("Flags:")
		fmt.Println("  -id, --id <int>          Todo id")
		fmt.Println("  -name, --name <string>   Todo name")
		fmt.Println("  -description, --description <string>   Todo description")
		fmt.Println("  -completed, --completed=<bool>   Todo completed")
		fmt.Println("  --help               Show this help message and exit")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide a command, expected one of: add, update, delete, list, get")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		todoFlagService.TodoAdd(addCmd)
	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		todoFlagService.TodoUpdate(updateCmd)
	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		todoFlagService.TodoDelete(deleteCmd)
	case "list":
		todoFlagService.TodoList()
	case "get":
		getCmd := flag.NewFlagSet("get", flag.ExitOnError)
		todoFlagService.TodoGet(getCmd)
	default:
		log.Fatal("Unknown command")
		os.Exit(1)
	}

}
