package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// Структура юзера в БД
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {
	// Открываем подключение к базе данных.
	// Данные из нашего doker-compose.yml
	// MYSQL_USER: "vladimir"
	// MYSQL_PASSWORD: "QWERTY"
	// Используем TCP соединение и аргументами передаем
	//     1) - Название сервиса БД прописанное в doker-compose (db)
	//     2) - Порт (3306)
	// MYSQL_DATABASE: "my_db"
	db, err := sql.Open("mysql", "vladimir:QWERTY@tcp(db:3306)/my_db")

	// Если при открытии соединения произошла ошибка, выведем ошибку
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Выполняем запрос
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // Обработка ошибок вместо паники 
	}

	var users []*User
	for results.Next() {
		var u User
		// Перебираем юзеров и добавляем их из БД в массив
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error()) // Обработка ошибок вместо паники 
		}

		users = append(users, &u)
	}

	return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Добро пожаловать на главную страницу!")
	fmt.Println("Endpoint Hit: homePage")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}