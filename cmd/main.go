package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"svi_danie/internal/handlers"
	"svi_danie/internal/repositories"
	"svi_danie/internal/services"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	// Строка подключения к базе данных
	connStr := "host=local-postgres port=5432 user=postgres password=postgres dbname=svi_db sslmode=disable"

	// Установка соединения с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	userRepo := &repositories.UserRepository{
		Db: db,
	}
	userService := &services.UserService{
		UserRepo: userRepo,
	}
	userHandler := &handlers.UserHandler{
		UserService: userService,
	}

	projRepo := &repositories.ProjectRepository{
		Db: db,
	}
	projService := &services.ProjService{
		ProRepo: projRepo,
	}
	projHandler := &handlers.ProjectHandler{
		ProjService: projService,
	}

	pageRepo := &repositories.PageRepository{
		Db: db,
	}
	pageService := &services.PageService{
		PageRepo: pageRepo,
	}
	pageHandler := &handlers.PageHandler{
		PageService: pageService,
	}

	projService.PageService = pageService

	imgRepo := &repositories.ImgRepository{
		Db: db,
	}
	imgService := &services.ImgService{
		ImgRepo: imgRepo,
	}
	imgHandler := &handlers.ImgHandler{
		ImgService: imgService,
	}

	// Создаем маршрутизатор
	mux := http.NewServeMux()

	// Регистрируем маршруты
	mux.HandleFunc("/create_img", imgHandler.CreateImage)
	mux.HandleFunc("/get_img", imgHandler.GetImage)
	mux.HandleFunc("/add_user", userHandler.AddUser)
	mux.HandleFunc("/add_proj", projHandler.CreateProj)
	mux.HandleFunc("/delete_proj", projHandler.DeleteProj)
	mux.HandleFunc("/get_proj", projHandler.GetProj)
	mux.HandleFunc("/get_all_proj", projHandler.GetAllProj)
	mux.HandleFunc("/add_page", pageHandler.CreatePage)
	mux.HandleFunc("/delete_page", pageHandler.DeletePage)
	mux.HandleFunc("/edit_page", pageHandler.EditPage)
	mux.HandleFunc("/get_page", pageHandler.GetPage)
	mux.HandleFunc("/get_all_pages", pageHandler.GetAllProjectPages)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Добро пожаловать на сервер!") // Отправляет сообщение на корневом маршруте
	})

	// Настройка CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	// Оборачиваем маршрутизатор в обработчик CORS
	handler := c.Handler(mux)

	// Запускаем сервер
	port := "5003"
	fmt.Printf("Сервер запущен на http://localhost:%s\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	if err != nil {
		log.Println(err)
	}
}
