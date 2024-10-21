package model

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}
