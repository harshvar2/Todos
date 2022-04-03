package usecase

import (
	"testing"
	"todos/domain"
	"todos/domain/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchTodos(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	t.Run("Test1", func(t *testing.T) {
		mockTodoRepo.On("FetchTodos").Return([]domain.Todo{}, nil).Once()
		u := NewTodoUsecase(mockTodoRepo)
		result, err := u.FetchTodos()
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.IsType(t, []domain.Todo{}, result)
	})
}

func TestGetTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	t.Run("Test1", func(t *testing.T) {
		todoID := 1
		mockTodoRepo.On("GetTodo", mock.AnythingOfType("int")).Return(domain.Todo{}, nil).Once()
		u := NewTodoUsecase(mockTodoRepo)
		result, err := u.GetTodo(todoID)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.IsType(t, domain.Todo{}, result)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestCreateTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	t.Run("Test1", func(t *testing.T) {
		todo := domain.Todo{Name: "Todo", Address: "localhost", DateOfBirth: "29-09-2014"}
		mockTodoRepo.On("CreateTodo", mock.AnythingOfType("domain.Todo")).Return(nil).Once()
		u := NewTodoUsecase(mockTodoRepo)
		err := u.CreateTodo(todo)
		assert.Nil(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestUpdateTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	t.Run("Test1", func(t *testing.T) {
		todo := domain.Todo{Name: "Todo", Address: "localhost", DateOfBirth: "29-09-2014"}
		mockTodoRepo.On("UpdateTodo", mock.AnythingOfType("domain.Todo")).Return(nil).Once()
		u := NewTodoUsecase(mockTodoRepo)
		err := u.UpdateTodo(todo)
		assert.Nil(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestDeleteTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	t.Run("Test1", func(t *testing.T) {
		todoID := 1
		mockTodoRepo.On("DeleteTodo", mock.AnythingOfType("int")).Return(nil).Once()
		u := NewTodoUsecase(mockTodoRepo)
		err := u.DeleteTodo(todoID)
		assert.Nil(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
}
