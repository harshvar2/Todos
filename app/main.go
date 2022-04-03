package main

import (
	"log"
	"todos/config"

	"todos/todo/delivery/http"
	mysql "todos/todo/repo"
	"todos/todo/usecase"

	docs "todos/docs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var (
	e *echo.Echo
)

func main() {
	// Initialise echo context for routes
	e = echo.New()

	// Load database config from config.yml
	err := config.GetDatabaseConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db, err := gorm.Open(config.DatabaseConfig.DbType, config.DatabaseConfig.DatabaseURL)
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.LogMode(true)
	todoRepo := mysql.NewMysqlUserRepository(db)
	todoUsecase := usecase.NewUserUsecase(todoRepo)
	http.NewUserHandler(e, todoUsecase)

	docs.NewDocumentation(e, e.Group(""))
	port := viper.GetString("APPLICATION_PORT")
	log.Fatal(e.Start(":" + port))
}
