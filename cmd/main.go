package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"svi_danie/internal/handlers"
	"svi_danie/internal/repositories"
	"svi_danie/internal/services"
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

	// Регистрируем маршрут
	http.HandleFunc("/add_user", userHandler.AddUser)

	http.HandleFunc("/add_proj", projHandler.CreateProj)
	http.HandleFunc("/delete_proj", projHandler.DeleteProj)
	http.HandleFunc("/get_proj", projHandler.GetProj)
	http.HandleFunc("/get_all_proj", projHandler.GetAllProj)

	http.HandleFunc("/add_page", pageHandler.CreatePage)
	http.HandleFunc("/delete_page", pageHandler.DeletePage)
	http.HandleFunc("/edit_page", pageHandler.EditPage)
	http.HandleFunc("/get_page", pageHandler.GetPage)
	http.HandleFunc("/get_all_pages", pageHandler.GetAllProjectPages)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Добро пожаловать на сервер!") // Отправляет сообщение на корневом маршруте
	})

	// Запускаем сервер
	port := "5003"
	fmt.Printf("Сервер запущен на http://localhost:%s", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Println(err)
	}

	//pageRepo := &repositories.PageRepository{
	//	Db: db,
	//}

	//userRepo := &repositories.UserRepository{
	//	Db: db,
	//}

	// Пример создания нового пользователя
	//newUser := models.User{
	//	Id:       uuid.New(),
	//	Login:    "new_user",
	//	Password: "password123",
	//}
	//err = userRepo.Create(newUser)
	//if err != nil {
	//	fmt.Printf("create: %s\n", err)
	//}
	//
	//// Пример чтения пользователя
	//userID := newUser.Id
	//user, err := userRepo.Read(userID)
	//if err != nil {
	//	fmt.Printf("read: %s\n", err)
	//} else {
	//	fmt.Printf("id: %s\nlogin: %s\n", user.Id, user.Login)
	//}
	//
	//// Пример обновления пользователя
	//updatedUser := models.User{
	//	Id:       userID,
	//	Login:    "updated_user",
	//	Password: "new_password123",
	//}
	//err = userRepo.Update(updatedUser)
	//if err != nil {
	//	fmt.Printf("update: %s\n", err)
	//}

	// Пример удаления пользователя
	//err = userRepo.Delete(uuid.MustParse("99af3993-2402-402e-bad8-c277c2e0485e"))
	//if err != nil {
	//	fmt.Printf("delete: %s\n", err)
	//}

	//// Пример создания новой страницы
	//newPage := models.Page{
	//	Id:      uuid.New(),
	//	OwnerId: userID,
	//	Title:   "New Page Title",
	//	Data:    json.RawMessage(`{"key": "value"}`),
	//}
	//err = pageRepo.Create(newPage)
	//if err != nil {
	//	fmt.Printf("create: %s\n", err)
	//}
	//
	//// Пример чтения страницы
	//pageID := newPage.Id
	//page, err := pageRepo.Read(pageID)
	//if err != nil {
	//	fmt.Printf("read: %s\n", err)
	//} else {
	//	fmt.Printf("id: %s\ntitle: %s\n", page.Id, page.Title)
	//}
	//
	//// Пример обновления страницы
	//updatedPage := models.Page{
	//	Id:      newPage.Id,
	//	OwnerId: newPage.OwnerId,
	//	Title:   "Updated Page Title",
	//	Data:    json.RawMessage(`{"key": "updated_value"}`),
	//}
	//err = pageRepo.Update(updatedPage)
	//if err != nil {
	//	fmt.Printf("update: %s", err)
	//}
	//
	//// Пример удаления страницы
	//err = pageRepo.Delete(newPage.Id)
	//if err != nil {
	//	fmt.Printf("update: %s", err)
	//}

	//// Пример выполнения запроса
	//rows, err := db.Query("SELECT version();")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var version string
	//	if err := rows.Scan(&version); err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println("PostgreSQL version:", version)
	//}
	//
	//if err := rows.Err(); err != nil {
	//	log.Fatal(err)
	//}
}
