package main

import (
	"database/sql"
	"log"
	"svi_danie/internal/controllers"
	"svi_danie/internal/decorators"
	"svi_danie/internal/repositories"
	"svi_danie/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := initDbConnection()
	defer db.Close()

	userRepo := &repositories.UserRepository{Db: db}
	imageRepo := &repositories.ImgRepository{Db: db}
	projectRepo := &repositories.ProjectRepository{Db: db}
	pageRepo := &repositories.PageRepository{Db: db}

	userService := &services.UserService{UserRepo: userRepo}
	imageService := &services.ImgService{ImgRepo: imageRepo}
	pageService := &services.PageService{PageRepo: pageRepo}
	projectService := &services.ProjService{ProjRepo: projectRepo, PageRepo: pageRepo}

	authDecorator := &decorators.AuthDecorator{UserService: userService}

	userController := &controllers.UserController{UserService: userService}
	pageController := &controllers.PageController{
		AuthDecorator: authDecorator,
		PageService:   pageService,
		ProjService:   projectService,
	}
	projectController := &controllers.ProjectController{
		AuthDecorator:  authDecorator,
		ProjectService: projectService,
	}
	imageController := &controllers.ImageController{
		AuthDecorator: authDecorator,
		ImgService:    imageService,
	}

	router := gin.Default()
	router.Use(cors.New(initCorsConfig()))
	initRouter(router, userController, imageController, pageController, projectController)
	if err := router.Run(":5003"); err != nil {
		log.Println(err)
	}
}

func initRouter(router gin.IRouter, ctrlList ...controllers.Controller) {
	for _, ctrl := range ctrlList {
		ctrl.InitRouter(router)
	}
}

func initDbConnection() *sql.DB {
	// Строка подключения к базе данных
	connStr := "host=postgres-db port=5432 user=postgres password=postgres dbname=svi_db sslmode=disable"

	// Установка соединения с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	return db
}

func initCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"http://localhost:3001", "https://xv0z-qhtn-bbzg.gw-1a.dockhost.net/"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
	}
}
