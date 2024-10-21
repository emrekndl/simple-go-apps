package repository

import (
	"database/sql"
	"fmt"
	"log"

	"todo-list-api/model"

	_ "github.com/mattn/go-sqlite3"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		"id" TEXT NOT NULL PRIMARY KEY,
		"title" TEXT,
		"body" TEXT,
		"completed" INTEGER
	);`

	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}

	return db
}

func (r *TodoRepository) GetTodos() ([]model.Todo, error) {
	rows, err := r.DB.Query("SELECT id, title, body, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *TodoRepository) GetTodoByID(id string) (*model.Todo, error) {
	row := r.DB.QueryRow("SELECT id, title, body, completed FROM todos WHERE id = ?", id)

	var todo model.Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.Completed); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) CreateTodo(todo *model.Todo) error {
	stmt, err := r.DB.Prepare("INSERT INTO todos (id, title, body, completed) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.ID, todo.Title, todo.Body, todo.Completed)
	return err
}

func (r *TodoRepository) UpdateTodoByID(todo *model.Todo, id string) error {
	stmt, err := r.DB.Prepare("UPDATE todos SET title = ?, body = ?, completed = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Body, todo.Completed, id)
	return err
}

func (r *TodoRepository) DeleteTodoByID(id string) error {
	stmt, err := r.DB.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		log.Printf("Todo with ID %s not found", id)
		return fmt.Errorf("Todo not found")
	}

	return nil
}
