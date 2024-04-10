package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Структура юзера в БД
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {
	db, err := sql.Open("mysql", "vladimir:QWERTY@tcp(db:3306)/my_db")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	var users []*User
	for results.Next() {
		var u User
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, &u)
	}
	return users
}

func userPage(c *gin.Context) {
	users := getUsers()
	c.JSON(http.StatusOK, users)
}

func addUser(c *gin.Context) {
	u := new(User)
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sql.Open("mysql", "vladimir:QWERTY@tcp(db:3306)/my_db")
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)
}

func main() {
	router := gin.Default()

	// Если эндпоинта нет перенаправим на главную страницу
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "public/")
	})

	// Эндпоинты
	router.GET("/users", userPage)
	router.POST("/users", addUser)
	// Статические файлы
	router.StaticFS("/public", http.Dir("./public"))

	// Запускаем сервер
	router.Run(":8080")
}
