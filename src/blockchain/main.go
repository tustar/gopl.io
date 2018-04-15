package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {

	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/blockchain?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()
	v1 := router.Group("v1")

	//
	v1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	// user
	v1.POST("/user", register(db))
	router.Run(":4000")
}

func register(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Request.FormValue("name")
		password := c.Request.FormValue("password")
		rs, err := db.Exec("INSERT INTO user(name, password) VALUES (?, ?)", name, password)
		if err != nil {
			log.Fatalln(err)
		}

		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("insert user Id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{"msg": msg})
	}
}
