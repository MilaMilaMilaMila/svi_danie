package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"svi_danie/internal/repositories"
)

func main() {
	// Строка подключения к базе данных
	connStr := "host=localhost port=5433 user=postgres password=postgres dbname=svi_db sslmode=disable"

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

	//pageRepo := &repositories.PageRepository{
	//	Db: db,
	//}

	userRepo := &repositories.UserRepository{
		Db: db,
	}

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
	err = userRepo.Delete(uuid.MustParse("99af3993-2402-402e-bad8-c277c2e0485e"))
	if err != nil {
		fmt.Printf("delete: %s\n", err)
	}

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
