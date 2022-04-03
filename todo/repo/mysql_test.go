package mysql_test

import (
	"testing"
	"time"
	"todos/domain"
	mysql "todos/todo/repo"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var userColumn = []string{"id", "name", "date_of_birth", "address", "created_at", "updated_at"}

func TestFetchtodos(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	currentTime := time.Now()

	rows := sqlmock.NewRows(userColumn).
		AddRow(1, "Todo1", "19-09-1999", "INDIA", currentTime, currentTime).
		AddRow(2, "Todo2", "29-09-1999", "INDIA", currentTime, currentTime)

	mock.ExpectQuery("^SELECT (.+) FROM `todos`").WillReturnRows(rows)

	mysqlRepo := mysql.NewMysqlTodoRepository(db)

	todos, err := mysqlRepo.FetchTodos()

	assert.Nil(t, err)
	assert.NotNil(t, todos)
	assert.Equal(t, len(todos), 2)
	assert.Nil(t, err)
}

func TestGetTodo(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	currentTime := time.Now()

	row := sqlmock.NewRows(userColumn).
		AddRow(1, "Todo1", "19-09-1999", "INDIA", currentTime, currentTime)

	mock.ExpectQuery("^SELECT (.+) FROM `todos` WHERE .*").WillReturnRows(row)

	userID := 1
	mysqlRepo := mysql.NewMysqlTodoRepository(db)
	user, err := mysqlRepo.GetTodo(userID)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, 1)
	assert.NotNil(t, user)
}

func TestCreateTodo(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `todos` .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	currentTime := time.Now()

	newTodo := domain.Todo{ID: 1, Name: "Todo", Description: "INDIA", CreatedAt: currentTime}

	mysqlRepo := mysql.NewMysqlTodoRepository(db)
	err = mysqlRepo.CreateTodo(newTodo)
	assert.NoError(t, err)
}

func TestUpdateTodo(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `todos` SET .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	currentTime := time.Now()

	newTodo := domain.Todo{ID: 1, Name: "Todo", Description: "INDIA", CreatedAt: currentTime}

	mysqlRepo := mysql.NewMysqlTodoRepository(db)
	err = mysqlRepo.UpdateTodo(newTodo)
	assert.NoError(t, err)
}

func TestDeleteTodo(t *testing.T) {
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM `todos` WHERE .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userID := 1

	mysqlRepo := mysql.NewMysqlTodoRepository(db)
	err = mysqlRepo.DeleteTodo(userID)
	assert.NoError(t, err)
}
