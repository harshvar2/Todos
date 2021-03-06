// Code generated by mockery (devel). DO NOT EDIT.

package mocks

import (
	domain "todos/domain"

	mock "github.com/stretchr/testify/mock"
)

// TodoUsecase is an autogenerated mock type for the TodoUsecase type
type TodoUsecase struct {
	mock.Mock
}

// CheckDBConnection provides a mock function with given fields:
func (_m *TodoUsecase) CheckDBConnection() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTodo provides a mock function with given fields: todo
func (_m *TodoUsecase) CreateTodo(todo domain.Todo) error {
	ret := _m.Called(todo)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Todo) error); ok {
		r0 = rf(todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTodo provides a mock function with given fields: id
func (_m *TodoUsecase) DeleteTodo(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchTodos provides a mock function with given fields:
func (_m *TodoUsecase) FetchTodos() ([]domain.Todo, error) {
	ret := _m.Called()

	var r0 []domain.Todo
	if rf, ok := ret.Get(0).(func() []domain.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodo provides a mock function with given fields: id
func (_m *TodoUsecase) GetTodo(id int) (domain.Todo, error) {
	ret := _m.Called(id)

	var r0 domain.Todo
	if rf, ok := ret.Get(0).(func(int) domain.Todo); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Todo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTodo provides a mock function with given fields: todo
func (_m *TodoUsecase) UpdateTodo(todo domain.Todo) error {
	ret := _m.Called(todo)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Todo) error); ok {
		r0 = rf(todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
