package domain

import (
	"time"
)

// Todo struct for 'todos' table
type Todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TodoRepository represents a usecase interface for Todo
type TodoRepository interface {
	FetchTodos() ([]Todo, error)
	CreateTodo(todo Todo) (err error)
	UpdateTodo(todo Todo) (err error)
	DeleteTodo(id int) (err error)
	GetTodo(id int) (res Todo, err error)

	// CheckDBConnection checks the health state of the database instance
	CheckDBConnection() (err error)
}

// TodoUsecase represents a usecase interface for Todo
type TodoUsecase interface {
	FetchTodos() ([]Todo, error)
	CreateTodo(todo Todo) (err error)
	UpdateTodo(todo Todo) (err error)
	DeleteTodo(id int) (err error)
	GetTodo(id int) (res Todo, err error)

	// CheckDBConnection checks the health state of the database instance
	CheckDBConnection() (err error)
}

// Check struct is used for returning state of particular service
type Check struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type HealthCheck struct {
	Status  string   `json:"status"`
	Version string   `json:"version"`
	Checks  []*Check `json:"checks"`
}
