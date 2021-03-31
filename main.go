package main

import (
	"github.com/Lairon/db-go/pgdb"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/laironacosta/ms-gin-go/controllers"
	"github.com/laironacosta/ms-gin-go/migrations"
	repo "github.com/laironacosta/ms-gin-go/repository"
	"github.com/laironacosta/ms-gin-go/router"
	"github.com/laironacosta/ms-gin-go/services"
	"os"
)

func main() {
	gin := gin.Default()

	db := pgdb.NewPgDB(&pg.Options{
		Addr:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})
	migrations.Init(db)

	userRepo := repo.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := router.NewRouter(gin, userController)
	r.Init()

	gin.Run() // listen and serve on 0.0.0.0:8080
}
