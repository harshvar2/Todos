package usecase

import (
	"todos/domain"
)

type todoUsecase struct {
	todoRepo domain.TodoRepository
}

// NewTodoUsecase represents a new instance for Todo Interface
func NewTodoUsecase(todoRepo domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		todoRepo: todoRepo,
	}
}
func (u *todoUsecase) FetchTodos() (todos []domain.Todo, err error) {
	todos, err = u.todoRepo.FetchTodos()
	return
}
func (u *todoUsecase) GetTodo(id int) (todoResponse domain.Todo, err error) {
	todoResponse, err = u.todoRepo.GetTodo(id)
	return
}
func (u *todoUsecase) CreateTodo(todo domain.Todo) (err error) {
	err = u.todoRepo.CreateTodo(todo)
	return
}

func (u *todoUsecase) UpdateTodo(todo domain.Todo) (err error) {
	err = u.todoRepo.UpdateTodo(todo)
	return
}
func (u *todoUsecase) DeleteTodo(id int) (err error) {
	err = u.todoRepo.DeleteTodo(id)
	return
}
