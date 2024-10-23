package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type TodoRepository struct {
	database *sql.DB
}

func NewTodoRepository(database *sql.DB) *TodoRepository {
	return &TodoRepository{database}
}

func createSqliteUri(filepath string) string {
	return fmt.Sprintf("file:%s?cache=shared", filepath)
}

func InitDB(filepath string) *sql.DB {
	filepath = createSqliteUri(filepath)
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTableSql := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        completed BOOLEAN NOT NULL,
        created_at TEXT NOT NULL,
        updated_at TEXT NOT NULL
    );
    `
	_, err = database.Exec(createTableSql)
	if err != nil {
		log.Fatal(err)
	}

	return database
}

func (repo *TodoRepository) GetTodoById(id int) (*Todo, error) {
	query := "SELECT id, name, description, completed, created_at, updated_at FROM todos WHERE id = ?"
	row := repo.database.QueryRow(query, id)

	var todo Todo
	err := row.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (repo *TodoRepository) GetTodos() ([]*Todo, error) {
	query := "SELECT id, name, description, completed, created_at, updated_at FROM todos"
	rows, err := repo.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (repo *TodoRepository) CreateTodo(todo *Todo) error {
	todo.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	todo.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	query := "INSERT INTO todos (name, description, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := repo.database.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Name, todo.Description, todo.Completed, todo.CreatedAt, todo.UpdatedAt)
	return err
}

func (repo *TodoRepository) UpdateTodo(todo *Todo) error {
	todo.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	query := "UPDATE todos SET name = ?, description = ?, completed = ?, updated_at = ? WHERE id = ?"
	stmt, err := repo.database.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(todo.Name, todo.Description, todo.Completed, todo.UpdatedAt, todo.Id)
	return err
}

func (repo *TodoRepository) DeleteTodo(id int) error {
	query := "DELETE FROM todos WHERE id = ?"
	stmt, err := repo.database.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
