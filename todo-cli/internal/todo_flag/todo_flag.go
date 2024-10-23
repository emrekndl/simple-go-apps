package todo_flag

import (
	"flag"
	"fmt"
	"log"
	"os"
	"todo-cli/internal/db"
)

type TodoRepository interface {
	GetTodoById(id int) (*db.Todo, error)
	GetTodos() ([]*db.Todo, error)
	CreateTodo(todo *db.Todo) error
	UpdateTodo(todo *db.Todo) error
	DeleteTodo(id int) error
}

type TodoFlagService struct {
	t    *db.Todo
	repo TodoRepository
}

func NewTodoFlagService(todoRepository TodoRepository) *TodoFlagService {
	t := &db.Todo{}
	return &TodoFlagService{
		t:    t,
		repo: todoRepository,
	}
}

func (todo *TodoFlagService) TodoAdd(f *flag.FlagSet) {
	t := todo.t
	f.StringVar(&t.Name, "name", "", "Todo name")
	f.StringVar(&t.Description, "description", "", "Todo description")
	f.BoolVar(&t.Completed, "completed", false, "Todo completed")
	f.Parse(os.Args[2:])

	if err := todo.repo.CreateTodo(t); err != nil {
		log.Fatal(err)
	}
}

func (todo *TodoFlagService) TodoUpdate(f *flag.FlagSet) {
	t := todo.t
	f.IntVar(&t.Id, "id", 0, "Todo id")
	f.StringVar(&t.Name, "name", "", "Todo name")
	f.StringVar(&t.Description, "description", "", "Todo description")
	f.BoolVar(&t.Completed, "completed", false, "Todo completed")
	f.Parse(os.Args[2:])

	if err := todo.repo.UpdateTodo(t); err != nil {
		log.Fatal(err)
	}
}

func (todo *TodoFlagService) TodoDelete(f *flag.FlagSet) {
	t := todo.t
	f.IntVar(&t.Id, "id", 0, "Todo id")
	f.Parse(os.Args[2:])

	if err := todo.repo.DeleteTodo(t.Id); err != nil {
		log.Fatal(err)
	}
}

func (todo *TodoFlagService) TodoList() {
	listTodos, err := todo.repo.GetTodos()
	if err != nil {
		log.Fatal(err)
	}

	if len(listTodos) == 0 {
		log.Println("No todos found")
		return
	}

	fmt.Println("---------Todos---------")
	for _, todo := range listTodos {
		fmt.Printf("ID: %d, Name: %s, Description: %s, Completed: %t, Created At: %s, Updated At: %s\n",
			todo.Id, todo.Name, todo.Description, todo.Completed, todo.CreatedAt, todo.UpdatedAt)
	}
}

func (todo *TodoFlagService) TodoGet(flag *flag.FlagSet) {
	t := todo.t
	flag.IntVar(&t.Id, "id", 0, "Todo id")
	flag.Parse(os.Args[2:])

	if todo, err := todo.repo.GetTodoById(t.Id); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("ID: %d, Name: %s, Description: %s, Completed: %t, Created At: %s, Updated At: %s\n",
			todo.Id, todo.Name, todo.Description, todo.Completed, todo.CreatedAt, todo.UpdatedAt)
	}
}
