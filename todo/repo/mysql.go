package mysql

import (
	"time"
	"todos/domain"

	"github.com/jinzhu/gorm"
)

type mysqlTodoRepository struct {
	db *gorm.DB
}

// NewMysqlTodoRepository : create instance of repository
func NewMysqlTodoRepository(con *gorm.DB) domain.TodoRepository {
	return &mysqlTodoRepository{con}
}
func (m *mysqlTodoRepository) FetchTodos() (todos []domain.Todo, err error) {
	m.db.Find(&todos)
	return
}
func (m *mysqlTodoRepository) GetTodo(id int) (todo domain.Todo, err error) {
	err = m.db.First(&todo, id).Error
	return
}

func (m *mysqlTodoRepository) CreateTodo(todo domain.Todo) (err error) {
	todo.CreatedAt = time.Now()
	err = m.db.Create(&todo).Error
	return
}

func (m *mysqlTodoRepository) UpdateTodo(todo domain.Todo) (err error) {
	result := m.db.Model(&domain.Todo{}).Update(todo).RowsAffected
	if result == 0 {
		err = domain.ErrTodoNotFound
	}
	return
}

func (m *mysqlTodoRepository) DeleteTodo(id int) (err error) {
	result := m.db.Delete(&domain.Todo{}, id).RowsAffected
	if result == 0 {
		err = domain.ErrTodoNotFound
	}
	return
}
